package short

import (
	"net/url"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	_ "github.com/n-creativesystem/short-url/pkg/domain/validation"
	"github.com/n-creativesystem/short-url/pkg/utils"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
)

type Short struct {
	key    string
	url    string
	author string

	keyGenerated bool
}

func NewShort(url string, key, author string) *Short {
	keyEmpty := key == ""
	if keyEmpty {
		key = GenerateKey()
	}
	if author == "" {
		author = "anonymous"
	}
	return &Short{
		key:          key,
		url:          url,
		author:       author,
		keyGenerated: keyEmpty,
	}
}

func (s *Short) GetKey() string {
	return s.key
}

func (s *Short) GetURL() string {
	return s.url
}

func (s *Short) GetUrl() url.URL {
	v, _ := url.Parse(s.url)
	return *v
}

func (s *Short) GetEncryptURL() credentials.EncryptString {
	return credentials.NewEncryptString(s.url)
}

func (s *Short) GetAuthor() string {
	return s.author
}

func (s *Short) ReGenerate() {
	s.key = GenerateKey()
}

func (s *Short) KeyGenerated() bool {
	return s.keyGenerated
}

func (s *Short) Valid() error {
	err := validation.ValidateStruct(s,
		validation.Field(
			&s.url,
			validation.Required,
			is.URL,
		),
		validation.Field(
			&s.key,
			validation.Required,
			validation.RuneLength(1, 255),
		),
	)
	if err == nil {
		return nil
	}
	errRes := ValidationErrors{
		Errors: make(validation.Errors),
	}
	if errs, ok := err.(validation.Errors); ok {
		for key, err := range errs {
			e := err.(validation.Error)
			params := e.Params()
			if params == nil {
				params = make(map[string]interface{})
			}
			params["key"] = key
			_ = e.SetParams(params)
			errRes.Errors[key] = ValidationError(err.(validation.Error))
		}
	}
	return errRes
}

func (s *Short) ServiceURL(baseURL string) string {
	return utils.MustURL(baseURL, s.GetKey())
}

type ShortWithTimeStamp struct {
	*Short
	CreatedAt time.Time
	UpdatedAt time.Time
}
