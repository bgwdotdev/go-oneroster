package routes

import (
    "net/http"
    "github.com/go-chi/chi"
)

func Routes() *chi.Mux {
    r := chi.NewRouter() 
    r.Get("/", allUsers)
    return r
}

func allUsers(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("user: bob"))
}

