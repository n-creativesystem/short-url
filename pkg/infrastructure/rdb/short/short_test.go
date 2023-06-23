package short

import (
	"context"
	"errors"
	"testing"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/tests"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/shared_test/short"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto"
	"github.com/stretchr/testify/require"
)

func TestShort(t *testing.T) {
	require := require.New(t)
	err := crypto.Init()
	require.NoError(err)
	t.Parallel()
	for _, d := range config.RDB {
		d := d
		t.Run(d.String(), func(t *testing.T) {
			config.SetDriver(d)
			client := tests.GetTestDBAndMigrate("shorturl_by_short")
			rdb.SetClient(client)
			ctx := context.Background()
			repoImpl := NewRepository()
			short.TestShort(t, ctx, repoImpl)
		})
	}
}

func TestShortError(t *testing.T) {
	t.Parallel()
	for _, d := range config.RDB {
		d := d
		t.Run(d.String(), func(t *testing.T) {
			config.SetDriver(d)
			ctx := context.Background()
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			oldDB := rdb.GetDB()
			defer func() {
				db.Close()
				rdb.SetDB(oldDB)
			}()
			rdb.SetDB(db)
			drv := entsql.OpenDB(config.GetDriver().String(), db)
			rdb.SetClient(ent.NewClient(ent.Driver(drv)))
			mock.ExpectExec("DELETE").WillReturnError(errors.New("Error"))
			repoImpl := newShortImpl()
			deleted, err := repoImpl.Del(ctx, "", "")
			require.Error(t, err)
			require.False(t, deleted)

			mock.ExpectQuery("SELECT").WillReturnError(errors.New("Error"))
			v, err := repoImpl.findOne(ctx)
			require.Error(t, err)
			require.Nil(t, v)

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}
