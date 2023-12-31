package short

import (
	"context"

	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/domain/short"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/predicate"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent/shorts"
)

type shortImpl struct{}

func (d *shortImpl) Get(ctx context.Context, key string) (*short.Short, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	if v, err := d.findOne(ctx, shorts.KeyEQ(key)); err != nil {
		return nil, err
	} else {
		return v.Short, nil
	}
}

func (d *shortImpl) Put(ctx context.Context, value short.Short) (*short.ShortWithTimeStamp, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	db := rdb.GetExecutor(ctx)
	if value.IsNew() {
		model := db.Shorts.Create()
		model.SetKey(value.GetKey())
		model.SetURL(value.GetEncryptURL())
		model.SetAuthor(value.GetAuthor())
		if entity, err := model.Save(ctx); err != nil {
			return nil, err
		} else {
			return toModel(entity), nil
		}
	} else {
		where := shorts.KeyEQ(value.GetKey())
		model := db.Shorts.Update()
		model.SetURL(value.GetEncryptURL())
		model.Where(where)
		if count, err := model.Save(ctx); err != nil {
			return nil, err
		} else if count == 0 {
			return nil, repository.ErrRecordNotFound
		} else {
			return d.findOne(ctx, where)
		}
	}
}

func (d *shortImpl) Del(ctx context.Context, key, author string) (bool, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	db := rdb.GetExecutor(ctx)
	result, err := db.Shorts.Delete().Where(shorts.And(shorts.KeyEQ(key), shorts.AuthorEQ(author))).Exec(ctx)
	return result > 0, err
}

func (d *shortImpl) Exists(ctx context.Context, key string) (bool, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	db := rdb.GetExecutor(ctx)
	return db.Shorts.Query().Where(shorts.KeyEQ(key)).Exist(ctx)
}

func (d *shortImpl) FindAll(ctx context.Context, author string) ([]short.ShortWithTimeStamp, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	db := rdb.GetExecutor(ctx)
	values, err := db.Shorts.Query().Where(shorts.AuthorEQ(author)).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	results := make([]short.ShortWithTimeStamp, len(values))
	for idx, value := range values {
		v := toModel(value)
		results[idx] = *v
	}
	return results, nil
}

func (d *shortImpl) FindByKeyAndAuthor(ctx context.Context, key, author string) (*short.ShortWithTimeStamp, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	db := rdb.GetExecutor(ctx)
	ps := []predicate.Shorts{
		shorts.KeyEQ(key),
		shorts.AuthorEQ(author),
	}
	value, err := db.Shorts.Query().Where(ps...).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	}
	return toModel(value), nil
}

func (d *shortImpl) findOne(ctx context.Context, ps ...predicate.Shorts) (*short.ShortWithTimeStamp, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	db := rdb.GetExecutor(ctx)
	v, err := db.Shorts.Query().Where(ps...).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	}
	return toModel(v), nil
}

func toModel(v *ent.Shorts) *short.ShortWithTimeStamp {
	value := short.NewShort(v.URL.UnmaskedString(), v.Key, v.Author)
	return &short.ShortWithTimeStamp{
		Short:     value,
		CreatedAt: v.CreateTime,
		UpdatedAt: v.UpdateTime,
	}
}
