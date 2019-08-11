package handlers

import (
	"database/sql"
	"github.com/fffnite/go-oneroster/parameters"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func GetAllEnrollments(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api := apiRequest{
			Request: Req{w, r},
			DB:      db,
			ORData: OneRoster{
				Table:      "enrollments",
				Columns:    enrollmentsCols,
				OutputName: "enrollments",
			},
			Params: parameters.Parameters{},
			Fks: []FK{
				FK{"classSourcedId", "classes", "sourcedId", "class"},
				FK{"schoolSourcedId", "orgs", "sourcedId", "school"},
				FK{"userSourcedId", "users", "sourcedId", "user"},
			},
		}
		api.invoke()
	}
}

var enrollmentsCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"classSourcedId",
	"schoolSourcedId",
	"userSourcedId",
	"role",
	"primary",
	"beginDate",
	"endDate",
}
