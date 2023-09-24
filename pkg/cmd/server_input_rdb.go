package cmd

import (
	"context"

	config_infra "github.com/n-creativesystem/short-url/pkg/infrastructure/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/oauth2"
	short_infra "github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/short"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/user"
	"github.com/n-creativesystem/short-url/pkg/interfaces/router"
	oauth2client "github.com/n-creativesystem/short-url/pkg/service/oauth2_client"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
)

func getRDBInput(ctx context.Context) (*router.RouterInput, func(), error) {
	cfg := config_infra.NewDBConfig()
	db, err := rdb.NewDB(cfg)
	if err != nil {
		logging.FromContext(ctx).With(logging.WithErr(err)).ErrorContext(ctx, "Setup rdb")
		return nil, nil, err
	}
	rdb.SetClient(db.Client)
	beginner, err := rdb.NewBeginner()
	if err != nil {
		return nil, nil, err
	}
	socialRepository := user.NewRepository()
	shortRepository := short_infra.NewRepository()
	oauth2ClientRepository := oauth2.NewOAuthClient()
	oauth2ClientService := oauth2client.NewService(oauth2ClientRepository, beginner)
	oauth2Store := oauth2.NewOAuth2Token(ctx, 60, oauth2ClientRepository)

	input := router.NewRouterInput(
		shortRepository,
		oauth2Store,
		oauth2ClientService,
		beginner,
		socialRepository,
	)
	return input, func() { _ = db.Close() }, nil
}
