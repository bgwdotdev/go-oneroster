package handlers

import (
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

// Queries database connection for Orgs
func getAllOrgs(w http.ResponseWriter, r *http.Request) {
	var (
		sourcedId uuid.UUID
		name      string
	)
	db.Query("SELECT sourcedId, name FROM orgs").Scan(&sourcedId, &name)
	out := Out{
		Body: sourcedId, name,
	}
	render.JSON(w, r, out)
}

func getOrg(w http.ResponseWriter, r *http.Request) {
	// stuff here
	id := chi.URLParam(r, "id")
	db.QueryRow("SELECT sourceId, name FROM orgs WHERE sourcedId=id")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user: bob"))
}
