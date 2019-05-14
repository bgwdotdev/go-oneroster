package handlers

import (
	"database/sql"
	"github.com/go-chi/render"
	// "github.com/google/uuid"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

// Basic JSON response structure
type Out struct {
	Body string `json:"body"`
}

// Queries database connection for Orgs
func GetAllOrgs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		v := Query(q)
		fields, err := ValidateFields(q, publicCols)
		if err != nil {
			// TODO: status error: warning, invalidate_selection_field
		}
		// TODO: run filters["field"] through ValidateFields
		filters, logicalOp := ParseFilters(q, publicCols)

		// Select results from table
		statement := fmt.Sprintf("SELECT %v FROM orgs WHERE %v%v? %v %v%v? ORDER BY ? LIMIT ? OFFSET ?",
			fields,
			filters[0]["field"], filters[0]["predicate"],
			logicalOp,
			filters[1]["field"], filters[1]["predicate"])
		stmt, err := db.Prepare(statement)
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// TODO: replace with logging
		fmt.Println(r.URL.Query())

		rows, err := stmt.Query(filters[0]["value"], filters[1]["value"], v["sort"], v["limit"], v["offset"])
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		// Build results
		var orgs []map[string]interface{}
		for rows.Next() {
			org := FormatResults(rows)
			orgs = append(orgs, org)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}

		// Wrap results in object
		var output = struct {
			Orgs []map[string]interface{} `json:"orgs"`
		}{orgs}

		// Output results
		render.JSON(w, r, output)
	}
}

func GetOrg(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get object based off id from query
		id := chi.URLParam(r, "id")
		statement := fmt.Sprintf("SELECT sourcedId, name FROM orgs WHERE sourcedId='%v'", id)

		var org Org
		db.QueryRow(statement).Scan(&org.SourcedId, &org.Name)

		// Wrap result
		var output = struct {
			Org Org
		}{org}

		// Output result
		render.JSON(w, r, output)
	}
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user: bob"))
}

var publicCols = []string{"sourcedId",
	"status",
	"dateLastModified",
	"name",
	"type",
	"identifier",
	"parentSourcedId",
}

// JSON out per spec
type Org struct {
	SourcedId        string
	Status           string
	DateLastModified string
	Name             string
	Type             string
	Identifier       string
	Parent           struct {
		Href      string
		SourcedId string
		Type      string
	}
	Children []struct {
		Href      string
		SourcedId string
		Type      string
	}
}
