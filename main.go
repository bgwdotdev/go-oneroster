package main

import (
	"github.com/fffnite/go-oneroster/internal/database"
	"github.com/fffnite/go-oneroster/routes"
	"github.com/go-chi/chi"
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
	r := chi.NewRouter()

	// Create DB connection and execute
	db := database.ConnectDb()

	// Creates a users endpoint that can have different methods attached to it
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/", routes.Routes(db))
	})

	// Starts the webserver with the Router
	http.ListenAndServe(":3000", r)
}
