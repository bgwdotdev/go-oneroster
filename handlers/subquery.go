package handlers

import (
	"database/sql"
	"fmt"
)

// Queries a nested item from a table
func QueryNestedProperty(t, c string, id interface{}, db *sql.DB) []map[string]interface{} {
	statement := fmt.Sprintf("SELECT sourcedId, type FROM %v WHERE %v='%v'", t, c, id)

	fmt.Println(statement)
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
		rs = append(rs, r)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return rs
}
