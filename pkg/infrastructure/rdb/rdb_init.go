package rdb

import (
	"database/sql"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

type Client struct {
	*ent.Client
	db *sql.DB
}

func NewDB(cfg config.DBConfig) (*Client, error) {
	dsn, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("rdb: failed to build connection string: %v\n", err)
	}
	db, err := otelsql.Open(cfg.SQLDriver(), dsn)
	if err != nil {
		return nil, err
	}
	return newClient(cfg.Dialect(), db), nil
}

func newClient(dialect string, db *sql.DB) *Client {
	drv := entsql.OpenDB(dialect, db)
	SetDB(db)
	return &Client{
		Client: ent.NewClient(ent.Driver(drv), ent.Debug()),
		db:     db,
	}
}
