package main

import (
    "net/http"
    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", helloWorld)
    r.Route("/users", func (r chi.Router) {
        r.Get("/", allUsers)
    })

    http.ListenAndServe(":3000", r)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
}

func allUsers(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("user: bob"))
}
