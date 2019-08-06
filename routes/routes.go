package routes

import (
	"database/sql"
	"github.com/fffnite/go-oneroster/conf"
	"github.com/fffnite/go-oneroster/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	log "github.com/sirupsen/logrus"
)

func Routes(db *sql.DB) *chi.Mux {
	var c conf.AuthConfig
	err := c.Load()
	if err != nil {
		log.Error(err)
	}
	tokenAuth := jwtauth.New(c.KeyAlg, []byte(c.Key), nil)
	r := chi.NewRouter()
	r.Post("/login", handlers.Login())
	r.Get("/", handlers.AllUsers)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/orgs", handlers.GetAllOrgs(db))
		r.Get("/orgs/{id}", handlers.GetOrg(db))
	})
	return r
}
