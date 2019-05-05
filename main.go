package main

import (
	"GoOneRoster/conf"
	"GoOneRoster/routes"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var (
	err error
)

// generic catch function for error handling
func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	c, err := conf.Read()
	catch(err)
	r := chi.NewRouter()

	// Create DB connection and execute
	db := conf.ConnectDatabase(c)
	defer db.Close()

	// Creates a users endpoint that can have different methods attached to it
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/", routes.Routes(db))
	})

	// Starts the webserver with the Router
	http.ListenAndServe(":3000", r)
}
