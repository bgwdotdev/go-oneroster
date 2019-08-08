package handlers

import (
	"database/sql"
	"encoding/json"
	data "github.com/fffnite/go-oneroster/db"
	"github.com/fffnite/go-oneroster/parameters"
	"github.com/go-chi/render"
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

func (r *apiRequest) parse() ([]error, error) {
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

type output struct {
	Errors  []error
	Results []map[string]interface{}
}

var jError string = "statusSetInfo"
var jResults string = "results"

func (o output) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		jError:   o.Errors,
		jResults: o.Results,
	}
	return json.Marshal(data)
}

func GetAll(t string, c []string, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p parameters.Parameters
		api := apiRequest{
			Table:   t,
			Columns: c,
			Request: r,
			DB:      db,
			Params:  p,
		}
		errP, err := api.parse()
		if err != nil {
			render.JSON(w, r, errP)
			return
		}
		rows := data.QueryProperties(api.Table, api.Columns, api.Params, api.DB)
		defer rows.Close()
		results := api.query(rows)
		jResults = t
		o := output{errP, results}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, o)
	}
}
