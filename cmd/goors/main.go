package main

import (
	"github.com/fffnite/go-oneroster/internal/conf"
	"github.com/fffnite/go-oneroster/internal/database"
	"github.com/fffnite/go-oneroster/internal/routes"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
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
	conf.LoadEnvs()

	r := chi.NewRouter()

	// Create DB connection and execute
	db := database.ConnectDb()

	// Creates a users endpoint that can have different methods attached to it
	r.Route("/ims/oneroster/v1p1", func(r chi.Router) {
		r.Mount("/", routes.Routes(db))
	})

	// Starts the webserver with the Router
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
