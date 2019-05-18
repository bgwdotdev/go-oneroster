package db

import (
	"GoOneRoster/parameters"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

// Queries the database based off endpoint
func QueryProperties(t string, c []string, p *parameters.Parameters, db *sql.DB) *sql.Rows {
	// Build Dynamic where query
	w := fmt.Sprintf("%v%v? %v %v%v?", p.Filter1.Field, p.Filter1.Predicate,
		p.LogicalOperator,
		p.Filter2.Field, p.Filter2.Predicate)
	// Convert string to uint64
	limit, err := strconv.ParseUint(p.Limit, 10, 64)
	if err != nil {
		panic(err)
	}
	offset, err := strconv.ParseUint(p.Offset, 10, 64)
	if err != nil {
		panic(err)
	}

	// Create sql query
	s, args, err := squirrel.
		Select(p.Fields).
		From(t).
		Where(w, p.Filter1.Value, p.Filter1.Value).
		OrderBy(p.Sort).
		Limit(limit).
		Offset(offset).
		ToSql()
	if err != nil {
		panic(err)
	}
	// TODO: remove after finding purpose
	fmt.Sprintf("squirrel: %v", args)

	// execute query
	stmt, err := db.Prepare(s)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	return rows
}
