package handlers

import (
	"database/sql"
	"github.com/go-chi/render"
	// "github.com/google/uuid"
	"fmt"
	data "github.com/fffnite/go-oneroster/db"
	"github.com/fffnite/go-oneroster/parameters"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// Queries database connection for Orgs
func GetAllOrgs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r)
		t := "orgs"
		q := r.URL.Query()
		var p parameters.Parameters
		ep, err := p.Resolve(q, publicCols)
		if err != nil {
			render.JSON(w, r, ep)
			return
		}
		rows := data.QueryProperties(t, publicCols, p, db)
		defer rows.Close()

		// Build results
		var orgs []map[string]interface{}
		for rows.Next() {
			org := data.FormatResults(rows)
			if strings.Contains(p.Fields, "parent") {
				org["parent"] = data.QueryNestedProperty("orgs", "sourcedId", org["parentSourcedId"], db, r.URL)
			}
			delete(org, "parentSourcedId")
			orgs = append(orgs, org)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}

		// Wrap results in object
		var output = struct {
			// TODO: links (HTTP link header / HTTP Header: x-total-count
			Errors []error                  `json:"statusInfoSet"`
			Orgs   []map[string]interface{} `json:"orgs"`
		}{ep, orgs}

		// Output results
		render.Status(r, http.StatusOK)
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
		//org["children"] = data.QueryNestedProperty("orgs", "parentSourcedId", org["sourcedId"], db)

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
