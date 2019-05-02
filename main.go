package main

import (
    "fmt"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "GoOneRoster/routes"
)

// To be moved to conf.hcl
const (
    dbDSN = "OneRoster.s3db"
    dbDriver = "sqlite3"
)

var (
    db  *sql.DB
    err error
)

// generic catch function for error handling
func catch(err error) {
    if err != nil {
        panic(err)
    }
}

// Create basic database 
func dbMake(db *sql.DB) {
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS test ( name string, age int )")
    catch (err)
    _, err = db.Exec("INSERT INTO test (name, age) VALUES ('bob', 1)")
    catch (err)
}

// Query database for name
func dbOut(db *sql.DB) string {
    var name string
    err := db.QueryRow("SELECT name FROM test").Scan(&name)
    catch(err)
    fmt.Print(name)
    return name
}

// Basic JSON response structure
type Out struct {
    Body string `json:"body"`
}

func main() {
    r := chi.NewRouter()

    // Create DB connection and execute
    db, err = sql.Open(dbDriver, dbDSN)
    catch (err)
    defer db.Close()
    dbMake(db)
    dbOut(db)

    // Creates a root endpoint with get method returning helloWorld func results
    r.Get("/", helloWorld)
    r.Get("/user", returnUser)
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

// Queries database and returns name as json
func returnUser(w http.ResponseWriter, r *http.Request) {
    n := dbOut(db)
    out := Out{
        Body: n,
    }
    render.JSON(w, r, out)
}

