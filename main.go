package main

import (
    "net/http"
    "github.com/go-chi/chi"
)

func main() {
    r := chi.NewRouter()

    // Creates a root endpoint with get method returning helloWorld func results
    r.Get("/", helloWorld)
    // Creates a users endpoint that can have different methods attached to it
    r.Route("/users", func (r chi.Router) {
        r.Get("/", allUsers)
    })

    // Starts the webserver with the Router
    http.ListenAndServe(":3000", r)
}

// outputs hello world
func helloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
}

// outputs user information
func allUsers(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("user: bob"))
}
