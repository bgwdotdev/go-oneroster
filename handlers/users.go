package handlers

import (
	"database/sql"
	"github.com/fffnite/go-oneroster/parameters"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api := apiRequest{
			Request: Req{w, r},
			DB:      db,
			ORData: OneRoster{
				Table:      "users",
				Columns:    userCols,
				OutputName: "Users",
			},
			Params: parameters.Parameters{},
			Fks: []FK{
				// new 1-N agents tables
				FK{"agentSourcedIds", "users", "sourcedId", "agents"},
				FK{"orgSourcedIds", "orgs", "sourcedId", "orgs"},
			},
		}
		api.invoke()
	}
}

var userCols = []string{
	"sourcedId",
	"status",
	"dateLastModified",
	"enabledUser",
	"orgSourcedIds",
	"role",
	"username",
	"userIds",
	"givenName",
	"familyName",
	"middleName",
	"identifier",
	"email",
	"sms",
	"phone",
	"agentSourcedIds",
	"grades",
	"password",
}
