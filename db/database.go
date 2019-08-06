package db

import (
	"database/sql"
	"github.com/fffnite/go-oneroster/conf"
	_ "github.com/mattn/go-sqlite3"
)

var err error

func ConnectDatabase(c conf.Config) *sql.DB {
	var conn *sql.DB
	conn, err = sql.Open(c.Database.DatabaseDriver, c.Database.DataSourceName)
	if err != nil {
		panic(err)
	}
	return conn
}
