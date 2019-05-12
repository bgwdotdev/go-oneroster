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

		// Select results from table
		statement := fmt.Sprintf("SELECT sourcedId, name FROM orgs WHERE %v ORDER BY '%v' LIMIT '%v' OFFSET '%v'",
			v["filter"], v["sort"], v["limit"], v["offset"])
		// replace with logging
		fmt.Println(r.URL.Query())
		fmt.Println(statement)
		rows, err := db.Query(statement)
		if err != nil {
			panic(err)
		}

		// Build results
		//var orgs []Org
		var orgs []map[string]interface{}
		for rows.Next() {
			org := FormatResults(rows)
			orgs = append(orgs, org)
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
