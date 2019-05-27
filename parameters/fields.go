package parameters

import (
	"GoOneRoster/helpers"
	log "github.com/sirupsen/logrus"
	"strings"
)

// Compares a slice of fields from a query against a slice of column names
// Returns the column names where an exact match was made
func ValidateFields(q map[string][]string, c []string) (string, error) {
	var columns []string
	all := strings.Join(c, ", ")

	// if no field key, return all/default columns
	if _, ok := q["fields"]; !ok {
		return all, nil
	}
	v := q["fields"][0]

	fields := strings.Split(v, ",")
	for _, s := range fields {
		col, err := validateField(s, c)
		if err != nil {
			err.(*helpers.Error).CodeMinor = "invalid_select_field"
			log.Info(err)
			return all, err
			/*
				CodeMajor : Success
				Severity : warning
				CodeMinor : invalid_select_field
				StatusCode : 200
				Description : Unknown field(s) %v
			*/
		}
		columns = append(columns, col)
	}
	output := strings.Join(columns, ", ")

	return output, nil
}

// Compares a single field from a query against a slice of column names
func validateField(s string, c []string) (string, error) {
	var f string
	for _, cn := range c {
		if cn == s {
			f = cn
		}
	}

	if f == "" {
		return s, &helpers.Error{Description: "No field match"}
	}

	return f, nil
}
