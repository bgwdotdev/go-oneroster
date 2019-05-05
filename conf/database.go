package conf

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var err error

func ConnectDatabase(c Config) *sql.DB {
	var conn *sql.DB
	conn, err = sql.Open(c.Database.DatabaseDriver, c.Database.DataSourceName)
	if err != nil {
		panic(err)
	}
	return conn
}
