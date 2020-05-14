package routes

import (
	"github.com/fffnite/go-oneroster/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(db2 *mongo.Client) *chi.Mux {
	keyA := viper.GetString("auth_key_alg")
	key := viper.GetString("auth_key")
	tokenAuth := jwtauth.New(keyA, []byte(key), nil)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/login", handlers.Login())
	r.Get("/", handlers.HelloWorld)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/orgs", handlers.GetAllOrgs(db2))
		r.Get("/orgs/{id}", handlers.GetOrg(db2))
		r.Put("/orgs/{id}", handlers.PutOrg(db2))
		r.Get("/academicSessions", handlers.GetAllAcademicSessions(db2))
		r.Get("/academicSessions/{id}", handlers.GetAcademicSession(db2))
		r.Put("/academicSessions/{id}", handlers.PutAcademicSession(db2))
		r.Get("/courses", handlers.GetAllCourses(db2))
		r.Get("/courses/{id}", handlers.GetCourses(db2))
		r.Put("/courses/{id}", handlers.PutCourses(db2))
		r.Get("/classes", handlers.GetAllClasses(db2))
		r.Get("/classes/{id}", handlers.GetClasses(db2))
		r.Put("/classes/{id}", handlers.PutClasses(db2))
		r.Get("/enrollments", handlers.GetAllEnrollments(db2))
		r.Get("/enrollments/{id}", handlers.GetEnrollments(db2))
		r.Put("/enrollments/{id}", handlers.PutEnrollments(db2))
		r.Get("/users", handlers.GetAllUsers(db2))
		r.Get("/users/{id}", handlers.GetUser(db2))
		r.Put("/users", handlers.PutBulkUser(db2))
		r.Put("/users/{id}", handlers.PutUser(db2))
		r.Put("/users/{id}/userIds/{subId}", handlers.PutUserId(db2))
	})
	return r
}
