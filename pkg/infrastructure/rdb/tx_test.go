package rdb

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
	"github.com/stretchr/testify/require"
)

func TestTransactionCommit(t *testing.T) {
	db, mock, _ := sqlmock.New()
	SetDB(db)
	mock.ExpectBegin()
	// ネストしてるけど1度しか呼ばれてないので1度のみ
	mock.ExpectCommit()
	beginner, err := NewBeginner()
	require.NoError(t, err)
	err = beginner.BeginTx(context.Background(), func(ctx context.Context) error {
		tx := GetExecutor(ctx)
		require.IsType(t, &ent.Client{}, tx)
		beginner2, _ := NewBeginner()
		err := beginner2.BeginTx(ctx, func(ctx context.Context) error {
			tx2 := GetExecutor(ctx)
			require.Equal(t, tx, tx2)
			return nil
		})
		require.NoError(t, err)
		return nil
	})
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestTransactionRollback(t *testing.T) {
	db, mock, _ := sqlmock.New()
	SetDB(db)
	mock.ExpectBegin()
	mock.ExpectRollback()
	beginner, err := NewBeginner()
	require.NoError(t, err)
	err = beginner.BeginTx(context.Background(), func(ctx context.Context) error {
		tx := GetExecutor(ctx)
		require.IsType(t, &ent.Client{}, tx)
		beginner2, _ := NewBeginner()
		err := beginner2.BeginTx(ctx, func(ctx context.Context) error {
			tx2 := GetExecutor(ctx)
			require.Equal(t, tx, tx2)
			return nil
		})
		require.NoError(t, err)
		return errors.New("Rollback")
	})
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
