package routes

import (
	"GoOneRoster/handlers"
	"database/sql"
	"github.com/go-chi/chi"
	// "net/http"
)

func Routes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handlers.AllUsers)
	r.Get("/orgs", handlers.GetAllOrgs(db))
	r.Get("/orgs/{id}", handlers.GetOrg(db))
	return r
}
