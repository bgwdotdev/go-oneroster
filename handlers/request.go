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
	Request Req
	DB      *sql.DB
	ORData  OneRoster
	Params  parameters.Parameters
	Ids     []ID
	Fks     []FK
}

type Req struct {
	W http.ResponseWriter
	R *http.Request
}

type OneRoster struct {
	Table      string
	Columns    []string
	ObjectType string
	OutputName string
}

type ID struct {
	SourcedId string
	Column    string
}

type FK struct {
	KeyColumn string
	RefTable  string
	RefColumn string
}

/*
func (r *apiRequest) validateId() {
	if r.Ids != nil {
		for _, v := range r.Ids {
			// do thing
		}
	}
}

func (r *apiRequest) validateFk() {
	if r.Fks != nil {
		for _, v := range r.Fks {
			// do things
		}
	}
}
*/

// sets and validates query parameters
// returns oneroster api error payload if invalid
func (r *apiRequest) validateParams() ([]error, error) {
	log.Info(r.Request)
	errp, err := r.Params.Resolve((r.Request.R.URL.Query()), r.ORData.Columns)
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
			result["parent"] = data.QueryNestedProperty(r.ORData.Table, "sourcedId",
				result["parentSourcedId"], r.DB, r.Request.R.URL)
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

func (a *apiRequest) invoke() {
	errP, err := a.validateParams()
	if err != nil {
		render.JSON(a.Request.W, a.Request.R, errP)
		return
	}
	rows := a.queryProperties()
	defer rows.Close()
	results := a.query(rows)
	jResults = a.ORData.OutputName
	o := output{errP, results}
	render.Status(a.Request.R, http.StatusOK)
	render.JSON(a.Request.W, a.Request.R, o)
}
