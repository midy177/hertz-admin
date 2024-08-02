package data

import (
	"fmt"
	"hertz-admin/configs"
	"hertz-admin/data/ent"
)

// NewClient returns a new ent client.
// Only support mysql and postgres
func NewClient(config configs.Config) (client *ent.Client, err error) {
	switch config.Database.Type {
	case "mysql":
		return NewMySQLClient(config)
	case "postgres":
		return NewPostgresClient(config)
	case "sqlite3":
		return NewSqliteClient(config)
	}
	return nil, fmt.Errorf("unsupported database: %s. only support mysql and postgres", config.Database.DBName)
}
