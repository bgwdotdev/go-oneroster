package routes

import (
	"GoOneRoster/handlers"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handlers.allUsers)
	return r
}

// Orgs endpoint
func Orgs() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handlers.getAllOrgs)
	r.Get("/{id}", handlers.getOrg)
	return r
}
