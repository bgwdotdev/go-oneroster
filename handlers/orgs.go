package handlers

import (
	"database/sql"
	"github.com/go-chi/render"
	// "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

// Basic JSON response structure
type Out struct {
	Body string `json:"body"`
}

// Queries database connection for Orgs
func GetAllOrgs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			//	sourcedId uuid.UUID
			name string
		)
		db.QueryRow("SELECT name FROM orgs").Scan(&name)
		out := Out{
			Body: name,
		}
		render.JSON(w, r, out)
	}
}

func GetOrg(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// stuff here
	// id := chi.URLParam(r, "id")
	db.QueryRow("SELECT sourceId, name FROM orgs WHERE sourcedId=id")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user: bob"))
}
