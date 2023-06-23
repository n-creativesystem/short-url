package oauth2

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	oauth2token "github.com/n-creativesystem/short-url/pkg/domain/oauth2_token"
	"github.com/n-creativesystem/short-url/pkg/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOAuth2Token(t *testing.T, ctx context.Context, repoImpl oauth2token.Repository) {
	defer repoImpl.Close()
	runTokenStoreTest(t, ctx, repoImpl)
}

func runTokenStoreTest(t *testing.T, ctx context.Context, store oauth2.TokenStore) {
	runTokenStoreCodeTest(t, ctx, store)
	runTokenStoreAccessTest(t, ctx, store)
	runTokenStoreRefreshTest(t, ctx, store)

	time.Sleep(3 * time.Second)
}

func runTokenStoreCodeTest(t *testing.T, ctx context.Context, store oauth2.TokenStore) {
	code := fmt.Sprintf("code %s", time.Now().String())

	tokenCode := models.NewToken()
	tokenCode.SetCode(code)
	tokenCode.SetCodeCreateAt(time.Now())
	tokenCode.SetCodeExpiresIn(time.Minute)
	require.NoError(t, store.Create(ctx, tokenCode))

	token, err := store.GetByCode(ctx, code)
	require.NoError(t, err)
	assert.Equal(t, code, token.GetCode())

	require.NoError(t, store.RemoveByCode(ctx, code))

	_, err = store.GetByCode(ctx, code)
	assert.Equal(t, repository.ErrRecordNotFound, err)
}

func runTokenStoreAccessTest(t *testing.T, ctx context.Context, store oauth2.TokenStore) {
	code := fmt.Sprintf("access %s", time.Now().String())

	tokenCode := models.NewToken()
	tokenCode.SetAccess(code)
	tokenCode.SetAccessCreateAt(time.Now())
	tokenCode.SetAccessExpiresIn(time.Minute)
	require.NoError(t, store.Create(ctx, tokenCode))

	token, err := store.GetByAccess(ctx, code)
	require.NoError(t, err)
	assert.Equal(t, code, token.GetAccess())

	require.NoError(t, store.RemoveByAccess(ctx, code))

	_, err = store.GetByAccess(ctx, code)
	assert.Equal(t, repository.ErrRecordNotFound, err)
}

func runTokenStoreRefreshTest(t *testing.T, ctx context.Context, store oauth2.TokenStore) {
	code := fmt.Sprintf("refresh %s", time.Now().String())

	tokenCode := models.NewToken()
	tokenCode.SetRefresh(code)
	tokenCode.SetRefreshCreateAt(time.Now())
	tokenCode.SetRefreshExpiresIn(time.Minute)
	require.NoError(t, store.Create(ctx, tokenCode))

	token, err := store.GetByRefresh(ctx, code)
	require.NoError(t, err)
	assert.Equal(t, code, token.GetRefresh())

	require.NoError(t, store.RemoveByRefresh(ctx, code))

	_, err = store.GetByRefresh(ctx, code)
	assert.Equal(t, repository.ErrRecordNotFound, err)
}
