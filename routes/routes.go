package routes

import (
	"database/sql"
	"github.com/fffnite/go-oneroster/conf"
	"github.com/fffnite/go-oneroster/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(db *sql.DB, db2 *mongo.Client) *chi.Mux {
	var c conf.AuthConfig
	err := c.Load()
	if err != nil {
		log.Error(err)
	}
	tokenAuth := jwtauth.New(c.KeyAlg, []byte(c.Key), nil)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/login", handlers.Login())
	r.Get("/", handlers.HelloWorld)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/orgs", handlers.GetAllOrgs(db2))
		r.Get("/orgs/{id}", handlers.GetOrg(db2))
		r.Get("/academicSessions", handlers.GetAllAcademicSessions(db))
		r.Get("/courses", handlers.GetAllCourses(db))
		r.Get("/classes", handlers.GetAllClasses(db))
		r.Get("/users", handlers.GetAllUsers(db))
		r.Get("/enrollments", handlers.GetAllEnrollments(db))
	})
	r.Get("/mongoOrgs", handlers.GetAllOrgs(db2))
	r.Get("/mongoOrgs/{id}", handlers.GetOrg(db2))
	r.Put("/mongoOrgs/{id}", handlers.PutOrg(db2))
	// r.Put("/orgs/{id}", handlers.PutOrg())
	return r
}

var orgsCols = []string{"sourcedId",
	"status",
	"dateLastModified",
	"name",
	"type",
	"identifier",
	"parentSourcedId",
}

var asCols = []string{"sourcedId",
	"status",
	"dateLastModified",
	"title",
	"type",
	"startDate",
	"endDate",
	"parentSourcedId",
	"schoolYear",
}
