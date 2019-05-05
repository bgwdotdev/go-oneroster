package main

import (
	"GoOneRoster/conf"
	"GoOneRoster/routes"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var (
	db  *sql.DB
	err error
)

// generic catch function for error handling
func catch(err error) {
	if err != nil {
		panic(err)
	}
}

// Basic JSON response structure
type Out struct {
	Body string `json:"body"`
}

func main() {
	r := chi.NewRouter()

	// Create DB connection and execute
	db, err = sql.Open(dbDriver, dbDSN)
	catch(err)
	defer db.Close()

	// Creates a users endpoint that can have different methods attached to it
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/users", routes.Routes())
		r.Mount("/orgs", routes.Orgs())
	})

	// Starts the webserver with the Router
	http.ListenAndServe(":3000", r)
}
