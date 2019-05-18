package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// Queries the database based off endpoint
func QueryProperties(t string, c []string, q []map[string]string, db *sql.DB) *sql.Rows {
	s := fmt.Sprintf("SELECT %v FROM %v WHERE %v%v? %v %v%v? ORDER BY ? LIMIT ? OFFSET ?",
		q["columns"], t,
		q["filter"][0]["field"], q["filter"][0]["predicate"],
		q["filter"]["logicalOperator"],
		q["filter"][1]["field"], q["filter"][1]["predicate"],
	)

	stmt, err := db.Prepare(s)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(q["filter"][0]["value"], q["filter"][1]["value"],
		q["sort"], q["limit"], q["offset"])
	if err != nil {
		panic(err)
	}

	return rows
}
