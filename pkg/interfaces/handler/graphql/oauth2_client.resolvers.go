package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"errors"
	"net/http"

	"github.com/n-creativesystem/short-url/pkg/interfaces/handler/graphql/models"
	"github.com/n-creativesystem/short-url/pkg/interfaces/request"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
	"github.com/n-creativesystem/short-url/pkg/service"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateOAuthApplication is the resolver for the createOAuthApplication field.
func (r *mutationResolver) CreateOAuthApplication(ctx context.Context, input models.OAuthApplicationInput) (*models.OAuthApplication, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	user, err := authorize(ctx)
	if err != nil {
		return nil, err
	}
	req := request.RegisterApplication{
		AppName: input.Name,
	}
	if err := req.Valid(); err != nil {
		return nil, err.GraphQLError(ctx)
	}
	result, err := r.oauth2clientSvc.RegisterClient(ctx, user, req.AppName)
	if err != nil {
		return nil, err
	}
	return r.Query().OauthApplication(ctx, result.ClientId)
}

// UpdateOAuthApplication is the resolver for the updateOAuthApplication field.
func (r *mutationResolver) UpdateOAuthApplication(ctx context.Context, id string, input models.OAuthApplicationInput) (*models.OAuthApplication, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	user, err := authorize(ctx)
	if err != nil {
		return nil, err
	}
	req := request.RegisterApplication{
		AppName: input.Name,
	}
	if err := r.oauth2clientSvc.UpdateClient(ctx, id, user, req.AppName); err != nil {
		return nil, err
	}
	return r.Query().OauthApplication(ctx, id)
}

// DeleteOAuthApplication is the resolver for the deleteOAuthApplication field.
func (r *mutationResolver) DeleteOAuthApplication(ctx context.Context, id string) (bool, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	user, err := authorize(ctx)
	if err != nil {
		return false, err
	}
	if err := r.oauth2clientSvc.DeleteClient(ctx, user, id); err != nil {
		return false, err
	}
	return true, nil
}

// OauthApplications is the resolver for the oauthApplications field.
func (r *queryResolver) OauthApplications(ctx context.Context, token *string) (*models.OAuthApplicationType, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	user, err := authorize(ctx)
	if err != nil {
		return nil, err
	}
	values, err := r.oauth2clientSvc.FindAll(ctx, user)
	if err != nil {
		return nil, err
	}
	result := models.OAuthApplicationType{
		Result:   make([]*models.OAuthApplication, len(values)),
		Metadata: &models.MetadataType{},
	}
	for idx, value := range values {
		app := response.OAuth2ApplicationResponseModel(value)
		result.Result[idx] = &app
	}
	return &result, nil
}

// OauthApplication is the resolver for the oauthApplication field.
func (r *queryResolver) OauthApplication(ctx context.Context, id string) (*models.OAuthApplication, error) {
	ctx, span := tracer.Start(ctx, "")
	defer span.End()
	user, err := authorize(ctx)
	if err != nil {
		return nil, err
	}
	if id == "" {
		return nil, errors.New("id is required field.")
	}
	value, err := r.oauth2clientSvc.FindByID(ctx, id, user)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return nil, &gqlerror.Error{
				Message:    err.Error(),
				Extensions: response.Extensions{}.SetCode(http.StatusNotFound),
			}
		}
		return nil, err
	}
	app := response.OAuth2ApplicationResponseModel(*value)
	return &app, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
