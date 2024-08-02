package data

import (
	"ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz-admin/configs"
	"hertz-admin/data/ent"
	"log"
	"os"
)

const (
	sqliteDsnBase = "./admin.db"
)

// NewSqliteClient returns a new ent client.
func NewSqliteClient(config configs.Config) (client *ent.Client, err error) {
	if config.Database.Host == "" {
		config.Database.Host = sqliteDsnBase
	}
	client, err = ent.Open("sqlite3",
		fmt.Sprintf("%s%s", config.Database.Host, "?cache=shared&mode=rwc&_fk=1"))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Create a local migration directory.
	migrationsPath := "./data/ent/migrations"
	_, err = os.Stat(migrationsPath)
	if err != nil {
		if !os.IsNotExist(err) {
			hlog.Fatalf("failed to stat migrations path: %v", err)
			return nil, err
		}
		// Create the directory if it doesn't exist.
		err = os.MkdirAll(migrationsPath, os.ModePerm)
		if err != nil {
			hlog.Fatalf("failed creating migrations path: %v", err)
			return nil, err
		}
	}
	dir, err := migrate.NewLocalDir(migrationsPath)
	if err != nil {
		hlog.Fatalf("failed creating atlas migration directory: %v", err)
		return nil, err
	}
	// Write migration diff.
	err = client.Schema.Diff(context.Background(), schema.WithDir(dir), schema.WithForeignKeys(false))
	if err != nil {
		hlog.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), schema.WithForeignKeys(false)); err != nil {
		hlog.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}

	return
}
