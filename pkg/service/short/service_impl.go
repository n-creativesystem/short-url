package short

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/domain/short"
	"github.com/n-creativesystem/short-url/pkg/domain/tx"
	"github.com/n-creativesystem/short-url/pkg/service"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"github.com/skip2/go-qrcode"
)

type serviceImpl struct {
	repo      short.Repository
	appConfig *config.Application
	tx        tx.ContextBeginner
}

func newServiceImpl(repo short.Repository, appConfig *config.Application, tx tx.ContextBeginner) *serviceImpl {
	return &serviceImpl{repo: repo, appConfig: appConfig, tx: tx}
}

func NewService(repo short.Repository, appConfig *config.Application, tx tx.ContextBeginner) Service {
	return newServiceImpl(repo, appConfig, tx)
}

func (impl *serviceImpl) GetURL(ctx context.Context, key string) (string, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	result, err := impl.repo.Get(ctx, key)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return "", service.ErrNotFound
		}
		return "", service.Wrap(err, "Service shortURL: An error occurred while retrieving the URL.")
	}
	return result.GetURL(), nil
}

func (impl *serviceImpl) GenerateShortURL(ctx context.Context, url, key, author string) (*short.ShortWithTimeStamp, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	value := short.NewShort(url, key, author)
	value.New()
	var result *short.ShortWithTimeStamp
	err := impl.tx.BeginTx(ctx, func(ctx context.Context) error {
		if err := value.Valid(); err != nil {
			return service.NewClientError(err)
		}

		var existsCheckErr = func(err error) error {
			return service.Wrap(err, "Service shortURL: An error occurred while checking for duplicates.")
		}

		// キーの重複チェック
		isExists, err := impl.repo.Exists(ctx, value.GetKey())
		if err != nil {
			return existsCheckErr(err)
		}
		// 自動生成でない場合はエラーとする
		if !value.KeyGenerated() {
			return service.ErrKeyDuplicate
		}
		loopCounter := 0
		for isExists && impl.appConfig.RetryGenerateCount > loopCounter {
			value.ReGenerate()
			isExists, err = impl.repo.Exists(ctx, value.GetKey())
			if err != nil {
				return existsCheckErr(err)
			}
			loopCounter++
		}

		if impl.appConfig.RetryGenerateCount > 0 && loopCounter == impl.appConfig.RetryGenerateCount {
			return errors.New("The number of URL generation attempts reached, but URL could not be generated.")
		}

		if v, err := impl.repo.Put(ctx, *value); err != nil {
			return service.Wrap(err, "Service shortURL: An error occurred during URL generation.")
		} else {
			result = v
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (impl *serviceImpl) GenerateQRCode(ctx context.Context, key string) (io.Reader, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	exists, err := impl.repo.Exists(ctx, key)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, service.Wrap(err, "Service shortURL: An error occurred while retrieving the URL.")
	}
	if !exists {
		return nil, service.ErrNotFound
	}
	png, err := qrcode.Encode(utils.MustURL(impl.appConfig.BaseURL, key), qrcode.Medium, 256)
	if err != nil {
		msg := "Service shortURL: QR Code generation failed."
		logging.FromContext(ctx).With(logging.WithErr(err)).ErrorContext(ctx, msg)
		return nil, errors.New(msg)
	}
	return bytes.NewReader(png), nil
}

func (impl *serviceImpl) Remove(ctx context.Context, key, author string) error {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	return impl.tx.BeginTx(ctx, func(ctx context.Context) error {
		if _, err := impl.repo.Del(ctx, key, author); err != nil {
			if errors.Is(err, repository.ErrRecordNotFound) {
				return nil
			}
			return service.Wrap(err, "Service shortURL: An error occurred during deletion.")
		}
		return nil
	})
}

func (impl *serviceImpl) Update(ctx context.Context, key, author, url string) (*short.ShortWithTimeStamp, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	v, err := impl.FindByKeyAndAuthor(ctx, key, author)
	if err != nil {
		return nil, service.Wrap(err, "Service shortURL: Update failed")
	}
	v.Short.SetURL(url)
	if err := v.Valid(); err != nil {
		return nil, err
	}
	if err = impl.tx.BeginTx(ctx, func(ctx context.Context) error {
		var err error
		v, err = impl.repo.Put(ctx, *v.Short)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return v, nil
}

func (impl *serviceImpl) FindAll(ctx context.Context, author string) ([]short.ShortWithTimeStamp, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	if v, err := impl.repo.FindAll(ctx, author); err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, service.Wrap(err, "Service shortURL: An error occurred while retrieving data.")
	} else {
		short.ShortWithTimeStamps(v).Update()
		return v, nil
	}
}

func (impl *serviceImpl) FindByKeyAndAuthor(ctx context.Context, key, author string) (*short.ShortWithTimeStamp, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	result, err := impl.repo.FindByKeyAndAuthor(ctx, key, author)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, service.Wrap(err, "Service shortURL: An error occurred while retrieving the URL.")
	}
	result.Update()
	return result, nil
}
