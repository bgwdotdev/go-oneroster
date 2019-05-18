package db

import (
	"GoOneRoster/parameters"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

// Queries the database based off endpoint
func QueryProperties(t string, c []string, p *parameters.Parameters, db *sql.DB) *sql.Rows {
	w := fmt.Sprintf("%v%v? %v %v%v?", p.F1.field, p.F1.Predicate, p.LogicalOperator, p.F2.field, p.F2.Predicate)
	s := squirrel.
		Select(p.Fields).
		From(t).
		Where(w, p.F1.value, p.F1.value).
		OrderBy(p.Sort).
		Limit(p.Limit).
		Offset(p.Offset).
		ToSql()

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
