package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocms-io/plugin-event-logger/init/database/sql"
)

type Database struct {
	SQL           *sql.SQL
}

func DefaultSQL() *Database {

	database := Database{
		SQL: sql.DefaultSQL(),
	}

	return &database
}
