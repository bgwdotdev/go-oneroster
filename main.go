package main

import (
    "net/http"
    "github.com/go-chi/chi"
    "GoOneRoster/routes"
)

// generic catch function for error handling
func catch(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    r := chi.NewRouter()

    // Creates a root endpoint with get method returning helloWorld func results
    r.Get("/", helloWorld)
    // Creates a users endpoint that can have different methods attached to it
    r.Route("/v1", func (r chi.Router) {
        r.Mount("/users", routes.Routes())
    })

    r.Mount("/schools", routes.Routes())
    // Starts the webserver with the Router
    http.ListenAndServe(":3000", r)
}

// outputs hello world
func helloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
}


