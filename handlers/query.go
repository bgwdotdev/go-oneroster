package handlers

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	data "github.com/fffnite/go-oneroster/db"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

// Queries the database based off endpoint
func (a *apiRequest) queryProperties() *sql.Rows {
	// Build Dynamic where query
	w := fmt.Sprintf("(%v%v? %v %v%v?)",
		a.Params.Filter1.Field, a.Params.Filter1.Predicate,
		a.Params.LogicalOperator,
		a.Params.Filter2.Field, a.Params.Filter2.Predicate)
	// Convert string to uint64
	limit, err := strconv.ParseUint(a.Params.Limit, 10, 64)
	if err != nil {
		panic(err)
	}
	offset, err := strconv.ParseUint(a.Params.Offset, 10, 64)
	if err != nil {
		panic(err)
	}

	// Create sql query
	rows, err := sq.
		Select(a.Params.Fields).
		From(a.ORData.Table).
		Where(w, a.Params.Filter1.Value, a.Params.Filter2.Value).
		OrderBy(a.Params.Sort).
		Limit(limit).
		Offset(offset).
		RunWith(a.DB).
		Query()
	if err != nil {
		panic(err)
	}
	return rows
}

// Queries database for a tables foreign key object and
// inserts as sub array into json results
func (a *apiRequest) queryFk(fk FK, id interface{}) []map[string]interface{} {
	rows, err := sq.
		Select("sourcedId").
		From(fk.RefTable).
		Where(sq.Eq{fk.RefColumn: id}).
		RunWith(a.DB).
		Query()
	if err != nil {
		//TODO: placeholder
		panic(err)
	}
	var rs []map[string]interface{}
	for rows.Next() {
		r := data.FormatResults(rows)
		r["type"] = fk.RefTable
		r["href"] = "/define/endpoint/here" //TODO: url endpoint
		rs = append(rs, r)
	}
	err = rows.Err()
	if err != nil {
		//TODO: placeholder
		panic(err)
	}
	return rs
}
