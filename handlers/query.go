package handlers

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	data "github.com/fffnite/go-oneroster/db"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

// Build Dynamic where query
func (a *apiRequest) buildWhere() string {
	return fmt.Sprintf("(%v%v? %v %v%v?)",
		a.Params.Filter1.Field, a.Params.Filter1.Predicate,
		a.Params.LogicalOperator,
		a.Params.Filter2.Field, a.Params.Filter2.Predicate)
}

// Queries the database based off endpoint
func (a *apiRequest) queryProperties() *sql.Rows {
	w := a.buildWhere()
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
		Select(fk.RefSelect + " AS sourcedId").
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

func (a *apiRequest) queryTotalCount() string {
	var count string
	w := a.buildWhere()
	rows, err := sq.
		Select("Count()").
		From(a.ORData.Table).
		Where(w, a.Params.Filter1.Value, a.Params.Filter2.Value).
		RunWith(a.DB).
		Query()
	if err != nil {
		// TODO: handle error
		panic(err)
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		// TODO: handle error
		panic(err)
	}
	return count
}

// Defines the next and previous limit/offsets for the api request
// based off current and max
func (a *apiRequest) queryLinkHeaders(count string) string {
    var link string
	url := a.Request.R.URL
	u := url.Scheme + url.Host + url.Path
	limit := url.Query().Get("limit")
	offset := url.Query().Get("offset")
	if limit == "" {
		limit = "100"
	}
	if offset == "" {
		offset = "0"
	}

	ilimit, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		panic(err)
	}
	ioffset, err := strconv.ParseUint(offset, 10, 64)
	if err != nil {
		panic(err)
	}
	icount, err := strconv.ParseUint(count, 10, 64)
	if err != nil {
		panic(err)
	}

    nextLimit := ilimit
    if icount < ioffset + ilimit {
        nextLimit = icount - ioffset
    }
	nextOffset := ioffset + nextLimit
    if ioffset != icount {
        link = link + fmt.Sprintf("<%v?limit=%v&offset=%v>; rel=\"next\",\n", u, nextLimit, nextOffset)
    }

	var prevOffset uint64
	if ioffset > ilimit {
		prevOffset = ioffset - ilimit
	}
    prevLimit := ilimit
    if int(ioffset) - int(ilimit) <= 0 {
        prevLimit = ioffset
    }
    if ioffset != 0 {
        link = link + fmt.Sprintf("<%v?limit=%v&offset=%v>; rel=\"prev\",\n", u, prevLimit, prevOffset)
    }
	return link
}
