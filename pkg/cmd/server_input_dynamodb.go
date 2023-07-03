package cmd

import (
	"context"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/aws"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb/oauth2"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/dynamodb/short"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/noop"
	"github.com/n-creativesystem/short-url/pkg/interfaces/router"
	oauth2client "github.com/n-creativesystem/short-url/pkg/service/oauth2_client"
)

func getDynamoDBInput(ctx context.Context) (*router.RouterInput, func(), error) {
	endpoint := aws.NewEndpoint()
	cfg, err := aws.NewConfig(ctx, endpoint.EndpointResolver())
	if err != nil {
		return nil, nil, err
	}
	client, err := dynamodb.NewDynamoDB(ctx, dynamodb.WithAwsConfig(cfg))
	if err != nil {
		return nil, nil, err
	}
	dynamodb.SetDB(client)
	beginner, _ := noop.NewBeginner()
	shortRepository := short.NewRepository()

	oauth2ClientRepository := oauth2.NewOAuth2Client()
	oauth2ClientService := oauth2client.NewService(oauth2ClientRepository, beginner)
	oauth2Store := oauth2.NewOAuth2Token(60, oauth2ClientRepository)

	input := router.NewRouterInput(
		shortRepository,
		oauth2Store,
		oauth2ClientService,
		beginner,
		nil,
	)
	return input, func() {}, nil
}
