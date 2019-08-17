package handlers

import (
	"database/sql"
	"github.com/fffnite/go-oneroster/parameters"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func GetAllCourses(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api := apiRequest{
			Request: Req{w, r},
			DB:      db,
			ORData: OneRoster{
				Table:      "courses",
				Columns:    courseCols,
				OutputName: "Courses",
			},
			Params: parameters.Parameters{},
			Fks: []FK{
				FK{"schoolYearSourcedId", "academicSessions", "sourcedId", "sourcedId", "schoolYear"},
				FK{"orgSourcedId", "orgs", "sourcedId", "sourcedId", "org"},
			},
		}
		api.invoke()
	}
}

var courseCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"schoolYearSourcedId",
	"title",
	"courseCode",
	"grades",
	"orgSourcedId",
	"subjects",
	"subjectCodes",
}
