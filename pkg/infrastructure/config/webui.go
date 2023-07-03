package config

import (
	"context"
	"time"

	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/sethvargo/go-envconfig"
)

type cors struct {
	AllowMethods []string      `env:"ALLOW_METHODS,default=GET,POST,OPTIONS"`
	AllowHeaders []string      `env:"ALLOW_HEADERS,default=*"`
	AllowOrigins []string      `env:"ALLOW_ORIGINS,default="`
	MaxAge       time.Duration `env:"MAX_AGE,default=14h"`
}

type csrf struct {
	TokenBase             bool   `env:"TOKEN,default=true"`
	CorsAndOriginalHeader bool   `env:"CORS_AND_ORIGINAL_HEADER,default=false"`
	HeaderName            string `env:"ORIGINAL_HEADER_NAME,default=X-Requested-With"`
}

type webUI struct {
	LoginSuccessURL  string `env:"LOGIN_REDIRECT,default=/"`
	LogoutSuccessURL string `env:"LOGOUT_REDIRECT,default=/"`
	Prefix           string `env:"API_PREFIX,default=/api"`
	IsUI             bool   `env:"IS_UI,default=false"`
	Cors             cors   `env:",prefix=CORS_"`
	Csrf             csrf   `env:",prefix=CSRF_"`
	Domain           string `env:"DOMAIN,default=localhost"`
	RedirectURI      string `env:"REDIRECT,default=localhost:3000"`
}

func NewWebUI(ctx context.Context) *config.WebUI {
	cfg := &webUI{}
	_ = envconfig.ProcessWith(ctx, cfg, envconfig.PrefixLookuper("WEB_UI_", envconfig.OsLookuper()))
	return &config.WebUI{
		LoginSuccessURL:  cfg.LoginSuccessURL,
		LogoutSuccessURL: cfg.LogoutSuccessURL,
		Prefix:           cfg.Prefix,
		IsUI:             cfg.IsUI,
		Cors: config.Cors{
			AllowMethods: cfg.Cors.AllowMethods,
			AllowHeaders: cfg.Cors.AllowHeaders,
			AllowOrigins: cfg.Cors.AllowOrigins,
			MaxAge:       cfg.Cors.MaxAge,
		},
		CSRF: config.Csrf{
			TokenBase:             cfg.Csrf.TokenBase,
			CorsAndOriginalHeader: cfg.Csrf.CorsAndOriginalHeader,
			HeaderName:            cfg.Csrf.HeaderName,
		},
		RedirectURI: cfg.RedirectURI,
	}
}
