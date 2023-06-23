package rdb

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
	"github.com/n-creativesystem/short-url/pkg/utils"
)

var (
	currentDB     *sql.DB
	currentClient *ent.Client
)

func SetDB(db *sql.DB) {
	currentDB = db
	SetClient(newEntClient(db))
}

func GetDB() *sql.DB {
	return currentDB
}

func GetClient() *ent.Client {
	client := currentClient
	if client == nil {
		drv := entsql.OpenDB(config.GetDriver().String(), currentDB)
		client = ent.NewClient(ent.Driver(drv))
	}
	if utils.IsCIorTest() {
		client = client.Debug()
	}
	return client
}

func SetClient(client *ent.Client) {
	currentClient = client
}

func newEntClient(db *sql.DB) *ent.Client {
	drv := entsql.OpenDB(config.GetDriver().String(), currentDB)
	return ent.NewClient(ent.Driver(drv))
}
