package db

import (
	"database/sql"
	"fmt"
	"net/url"
)

// Queries a nested item from a table
func QueryNestedProperty(t, c string, id interface{}, db *sql.DB, u *url.URL) []map[string]interface{} {
	statement := fmt.Sprintf("SELECT sourcedId FROM %v WHERE %v='%v'", t, c, id)

	stmt, err := db.Prepare(statement)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var rs []map[string]interface{}
	for rows.Next() {
		r := FormatResults(rows)
		r["type"] = t
		r["href"] = u.Host + "/v1/" + t + "/" + id.(string)
		rs = append(rs, r)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return rs
}
