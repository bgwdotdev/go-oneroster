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
	switch r.FindString(s) {
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
	err := validateField(v, safeCol)
	if err != nil {
		err.(*ErrorObject).CodeMinor = "invalid_filter_field"
		err.(*ErrorObject).Populate()
		return "", err
	}
	return v, nil
}

// parses the comparison operator from a filter string
// returns mongodb predicate
func parseFilterPredicate(s string) string {
	r := regexp.MustCompile(`([=><~!]{1,2})`)
	var res string
	switch p := r.FindString(s); p {
	case "=":
		res = "$eq"
	case "!=":
		res = "$ne"
	case ">":
		res = "$gt"
	case ">=":
		res = "$gte"
	case "<":
		res = "$lt"
	case "<=":
		res = "$lte"
	case "~":
		res = "$regex"
	default:
		res = "$eq"
	}
	return res
}

// parses the target value of a filter string held single quotes
func parseFilterValue(s string) string {
	r := regexp.MustCompile(`([']\S*['])`)
	return removeSingleQuotes(r.FindString(s))
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
// errors if no match
func validateField(s string, safeFields []string) error {
	for _, v := range safeFields {
		if s == v {
			return nil
		}
	}
	err := fmt.Sprintf("Unknown field: %v", s)
	return &ErrorObject{Description: err}
}

// splits up to 2 queries up into seperate slices
// "name='bob' age=>'18'
func splitFilterQuery(s string) []string {
	r := regexp.MustCompile(`([a-zA-Z]*[=><~!]{1,2}[']\S*['])`)
	var fs []string
	for i := 0; i < 2; i++ {
		if r.MatchString(s) {
			f := r.FindString(s)
			fs = append(fs, f)
			s = strings.Split(s, f)[1]
		}
	}
	return fs
}
