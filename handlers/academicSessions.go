package handlers

import (
	"database/sql"
	"github.com/fffnite/go-oneroster/parameters"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func GetAllAcademicSessions(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api := apiRequest{
			Request: Req{w, r},
			DB:      db,
			ORData: OneRoster{
				Table:      "academicSessions",
				Columns:    asCols,
				OutputName: "academicSessions",
			},
			Params: parameters.Parameters{},
			Fks:    []FK{FK{"parentSourcedId", "academicSessions", "sourcedId", "sourcedId", "parent"}},
		}
		api.invoke()
	}
}

var asCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"title",
	"type",
	"startDate",
	"endDate",
	"parentSourcedId",
	"schoolYear",
}
