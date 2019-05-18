package handlers

import (
	"database/sql"
	"github.com/go-chi/render"
	// "github.com/google/uuid"
	data "GoOneRoster/db"
	"GoOneRoster/parameters"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

// Queries database connection for Orgs
func GetAllOrgs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := "orgs"
		q := r.URL
		params := parameters.ParseUrl(q, publicCols)
		rows := data.QueryProperties(t, publicCols, params, db)
		defer rows.Close()
		// TODO: replace with logging
		fmt.Println(r.URL.Query())

		// Build results
		var orgs []map[string]interface{}
		for rows.Next() {
			org := parameters.FormatResults(rows)
			org["children"] = data.QueryNestedProperty("orgs", "parentSourcedId", org["sourcedId"], db)
			org["parent"] = data.QueryNestedProperty("orgs", "sourcedId", org["parentSourcedId"], db)
			orgs = append(orgs, org)
		}
		err := rows.Err()
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
