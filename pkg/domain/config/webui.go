package config

import (
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gin-contrib/cors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/n-creativesystem/short-url/pkg/domain/social"
	"github.com/n-creativesystem/short-url/pkg/utils"
)

type Cors struct {
	AllowMethods []string
	AllowHeaders []string
	AllowOrigins []string
	MaxAge       time.Duration
}

func (cfg *Cors) ToCorsConfig() cors.Config {
	config := cors.Config{
		AllowMethods:     cfg.AllowMethods,
		AllowHeaders:     cfg.AllowHeaders,
		AllowOrigins:     cfg.AllowOrigins,
		AllowCredentials: true,
		MaxAge:           cfg.MaxAge,
	}
	if utils.IsDev() {
		config.AllowAllOrigins = true
		config.AllowOrigins = nil
	}
	return config
}

type Csrf struct {
	TokenBase             bool
	CorsAndOriginalHeader bool
	HeaderName            string
}

func (c Csrf) IsValid() error {
	var (
		tokenBaseField = validation.Field(&c.TokenBase,
			validation.When(
				!c.CorsAndOriginalHeader,
				validation.Required,
			),
		)
		corsAndOriginalHeaderField = validation.Field(&c.CorsAndOriginalHeader,
			validation.When(
				!c.TokenBase,
				validation.Required,
			),
		)
		headerNameField = validation.Field(&c.HeaderName,
			validation.When(
				c.CorsAndOriginalHeader,
				validation.Required,
				validation.By(checkCustomHeaderName),
			),
		)
	)

	err := validation.ValidateStruct(&c,
		tokenBaseField,
		corsAndOriginalHeaderField,
		headerNameField,
	)
	if err == nil {
		return nil
	}
	if errs, ok := err.(validation.Errors); ok {
		tokenBase, _ := errs["TokenBase"].(validation.Error)
		corsAndOriginalHeader, _ := errs["CorsAndOriginalHeader"].(validation.Error)
		if tokenBase != nil && tokenBase.Code() == "validation_required" &&
			corsAndOriginalHeader != nil && corsAndOriginalHeader.Code() == "validation_required" {
			return ErrCSRFSetting
		}
		return errs
	} else {
		return err
	}
}

type WebUI struct {
	LoginSuccessURL  string
	LogoutSuccessURL string
	Prefix           string
	IsUI             bool
	Cors             Cors
	CSRF             Csrf
	Store            scs.Store
	Providers        map[string]*social.Config
	Domain           string
	RedirectURI      string
}
