package parameters

import (
	"errors"
	"strings"
)

// Compares a slice of fields from a query against a slice of column names
// Returns the column names where an exact match was made
func ValidateFields(q map[string][]string, c []string) (string, error) {
	var columns []string
	all := strings.Join(c, ", ")

	// if no field key, return all/default columns
	if _, ok := q["field"]; !ok {
		return all, nil
	}
	v := q["fields"][0]

	fields := strings.Split(v, ",")
	for _, s := range fields {
		col, err := validateField(s, c)
		if err != nil {
			return all, err
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
		err := errors.New("No field match")
		return s, err
	} else {
		return f, nil
	}
}
