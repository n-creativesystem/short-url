package rdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect/sql/schema"
	infra_config "github.com/n-creativesystem/short-url/pkg/infrastructure/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

type MigratorArgs struct {
	RetryCount   uint64
	RetryWait    uint64
	AllowMissing bool
	Table        string
	Dir          string
	Driver       string
}

func Migration(ctx context.Context, args []string, migratorArgs MigratorArgs) error {
	dbConfig := infra_config.NewDBConfig(infra_config.WithMigration())
	var (
		client *Client
		err    error
	)
	err = ExecuteWithRetry(migratorArgs.RetryCount, migratorArgs.RetryWait, func() error {
		if args[0] == "create" {
			client = GetMigrateClientWithConfig("migration", dbConfig)
			return nil
		} else {
			client, err = NewDB(dbConfig)
			if err == nil {
				return nil
			}
		}
		return fmt.Errorf("failed to open DB: %w", err)
	})

	if err != nil {
		return fmt.Errorf("migrator: %v\n", err)
	}

	defer func() {
		if err := client.db.Close(); err != nil {
			logging.FromContext(ctx).With(logging.WithErr(err)).ErrorContext(ctx, "migrator: failed to close DB")
		}
	}()

	err = ExecuteWithRetry(migratorArgs.RetryCount, migratorArgs.RetryWait, func() error {
		err = client.db.Ping()
		if err == nil {
			return nil
		}
		return fmt.Errorf("failed to ping DB: %w", err)
	})

	if err != nil {
		return fmt.Errorf("migrator: %v\n", err)
	}
	return MigrationWithDB(ctx, client, args, migratorArgs)
}

func MigrationWithDB(ctx context.Context, client *Client, args []string, migratorArgs MigratorArgs) error {
	var arguments []string
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}
	command := args[0]

	switch command {
	case "create":
		return createMigrateFile(ctx, client.Client, migratorArgs.Driver, arguments[0])
	default:
		_ = goose.SetDialect(migratorArgs.Driver)
		var gooseOpts []goose.OptionsFunc
		if migratorArgs.AllowMissing {
			gooseOpts = append(gooseOpts, goose.WithAllowMissing())
		}

		if migratorArgs.Table != "" {
			goose.SetTableName(migratorArgs.Table)
		}

		goose.SetLogger(logging.NewGooseLogger())
		if err := goose.RunWithOptions(command, client.db, migratorArgs.Dir, arguments, gooseOpts...); err != nil {
			return err
		}

		return nil
	}

}

// ExecuteWithRetry execute fn in specific retry count and wait.
func ExecuteWithRetry(count, wait uint64, fn func() error) error {
	retry := uint64(0)
	for {
		err := fn()
		if err == nil {
			return nil
		}
		if retry >= count {
			return fmt.Errorf("retry count exceeded. error: %w", err)
		}
		log.Printf("migrator: retry:%d error: %+v\n", retry, err)

		retry++
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
}

func createMigrateFile(ctx context.Context, client *ent.Client, driver string, name string) error {
	if name == "" {
		return errors.New("migration name is required. Use: 'go run -mod=mod main.go migrator create <name>'")
	}
	migrationDir := filepath.Join("db", "migrations", driver)
	_ = os.MkdirAll(migrationDir, 0744)
	dir, err := sqltool.NewGooseDir(migrationDir)
	if err != nil {
		return errors.Wrap(err, "failed creating atlas migration directory")
	}
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeReplay),
	}

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = client.Schema.NamedDiff(ctx, name, opts...)
	if err != nil {
		return errors.Wrap(err, "failed generating migration file")
	}
	return nil
}
