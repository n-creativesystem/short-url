package migrations

import (
	"embed"

	"github.com/pressly/goose/v3"
)

var (
	//go:embed mysql/* postgres/* sqlite/*
	Migration embed.FS
)

func init() {
	goose.SetBaseFS(Migration)
}
