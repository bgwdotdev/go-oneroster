package handlers

import (
	"database/sql"
	"github.com/fffnite/go-oneroster/parameters"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func GetAllClasses(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api := apiRequest{
			Request: Req{w, r},
			DB:      db,
			ORData: OneRoster{
				Table:      "classes",
				Columns:    classesCols,
				OutputName: "Classes",
			},
			Params: parameters.Parameters{},
			Fks: []FK{
				FK{"courseSourcedId", "courses", "sourcedId", "course"},
				FK{"termSourcedIds", "classAcademicSessions", "sourcedId", "terms"},
				FK{"schoolSourcedId", "orgs", "sourcedId", "school"},
			},
		}
		api.invoke()
	}
}

var classesCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"title",
	"grades",
	"courseSourcedId",
	"classCode",
	"classType",
	"location",
	"schoolSourcedId",
	"termSourcedIds",
	"subjects",
	"subjectCodes",
	"periods",
}
