package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

// parses the logical operator from a filter string
// supports OR AND
func parseFilterLo(s string) string {
	r := regexp.MustCompile(`(\sAND|OR\s)`)
	lo := r.FindString(s)
	switch lo := r.FindString(s); lo {
	case " OR ":
		return "$or"
	default:
		return "$and"
	}
}

// parses the field name from a filter string
// validates against allowed columns
func parseFilterField(s string, safeCol []string) (string, error) {
	r := regexp.MustCompile(`([a-zA-Z]*)`)
	v := r.FindString(s)
	field, err := validateField(v, safeCol)
	if err != nil {
		err.(*ErrorObject).CodeMinor = "invalid_filter_field"
		err.(*ErrorObject).Populate()
		return "", err
	}
	return field, nil
}

// parses the comparison operator from a filter string
// returns mongodb predicate
func parseFilterPredicate(s string) string {
	r := regexp.MustCompile(`([=><~!]{1,2})`)
	var r string
	switch p := r.FindString(s); p {
	case "=":
		r = "$eq"
	case "!=":
		r = "$ne"
	case ">":
		r = "$gt"
	case ">=":
		r = "$gte"
	case "<":
		r = "$lt"
	case "<=":
		r = "$lte"
	case "~":
		r = "$regex"
	default:
		r = "$eq"
	}
	return r
}

// parses the target value of a filter string held single quotes
func parseFilterValue(s string) string {
	r := regexp.MustCompile(`([']\S*['])`)
	return removeSingleQuotes(r.FindStrings(s))
}

// removes single straight quotes from the start and end of a string
func removeSingleQuotes(s string) string {
	b := regexp.MustCompile(`(^['])`)
	e := regexp.MustCompile(`([']$)`)
	s = b.Split(s, -1)[1]
	s = e.Split(s, -1)[0]
	return s
}

// Compares the requested field against a list of field names
// Returns matched field name
func validateField(s string, safeFields []string) (string, error) {
	var f string
	for _, v := range safeFields {
		if s == v {
			f = v
		}
	}
	if f == "" {
		err := fmt.Sprintf("Unknown field: %v", s)
		return "", &helpers.ErrorObject{Description: err}
	}
	return f, nil
}
