package short

import (
	"context"
	"testing"
	"time"

	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/n-creativesystem/short-url/pkg/domain/short"
	"github.com/stretchr/testify/require"
)

func TestShort(t *testing.T, ctx context.Context, repoImpl short.Repository) {
	require := require.New(t)

	// Create
	model := short.NewShort("http://localhost:8080/example", "", "")
	domainModel, err := repoImpl.Put(ctx, *model)
	require.NoError(err)
	require.NotNil(domainModel)
	require.NotEmpty(domainModel.CreatedAt)
	require.NotEmpty(domainModel.UpdatedAt)
	require.Equal(domainModel.GetURL(), "http://localhost:8080/example")

	// Read by key
	s, err := repoImpl.Get(ctx, model.GetKey())
	require.NoError(err)
	require.Equal(model.GetKey(), s.GetKey())
	require.Equal(model.GetURL(), s.GetURL())

	values, err := repoImpl.FindAll(ctx, model.GetAuthor())
	require.NoError(err)
	require.Len(values, 1)
	value := values[0]
	require.Equal(value.GetKey(), s.GetKey())
	require.Equal(value.GetURL(), s.GetURL())
	for _, v := range []time.Time{value.CreatedAt, value.UpdatedAt} {
		require.NotEmpty(v)
	}

	// Exists
	isExists, err := repoImpl.Exists(ctx, model.GetKey())
	require.NoError(err)
	require.True(isExists)

	// Del
	deleted, err := repoImpl.Del(ctx, model.GetKey(), model.GetAuthor())
	require.NoError(err)
	require.True(deleted)
	_, err = repoImpl.Get(ctx, model.GetKey())
	require.ErrorIs(err, repository.ErrRecordNotFound)
}
