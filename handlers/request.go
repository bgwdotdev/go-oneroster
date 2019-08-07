package handlers

import (
	"database/sql"
	data "github.com/fffnite/go-oneroster/db"
	"github.com/fffnite/go-oneroster/parameters"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type apiRequest struct {
	Table   string
	Columns []string
	Request *http.Request
	DB      *sql.DB
	Params  parameters.Parameters
}

func (r *apiRequest) Parse() ([]error, error) {
	log.Info(r.Request)
	errp, err := r.Params.Resolve((r.Request.URL.Query()), r.Columns)
	if err != nil {
		return errp, err
	}
	return errp, nil
}

func (r *apiRequest) query(rows *sql.Rows) []map[string]interface{} {
	var results []map[string]interface{}
	for rows.Next() {
		result := data.FormatResults(rows)
		if strings.Contains(r.Params.Fields, "parent") {
			result["parent"] = data.QueryNestedProperty(r.Table, "sourcedId",
				result["parentSourcedId"], r.DB, r.Request.URL)
		}
		delete(result, "parentSourcedId")
		results = append(results, result)
	}
	err := rows.Err()
	if err != nil {
		log.Error(err)
		return results
	}
	return results
}
