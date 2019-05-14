package handlers

import (
	"errors"
	"regexp"
	"strings"
)

var err error

// Takes a url query and converts it into map of filter parameters
func ParseFilters(q map[string][]string, cols []string) ([]map[string]string, string) {

	filters, lo := setFilterDefaults()
	// returns defaults if no user arguments
	if _, ok := q["filter"]; !ok {
		return filters, lo
	}

	s := q["filter"][0]
	lo = resolveLogicalOperator(s)
	vs := splitFilterQuery(s)

	for i, v := range vs {
		filters[i], err = parseFilter(v, cols)
	}

	return filters, lo
}

func parseFilter(v string, cols []string) (map[string]string, error) {
	m := make(map[string]string)

	f, err := parseFilterField(v, cols)
	if err != nil {
		//TODO: return web 4xx error no match
	}
	p := parseFilterPredicate(v)
	val := parseFilterValue(v)
	p, val = evaluateLikeStatement(p, val)

	m["field"] = f
	m["predicate"] = p
	m["value"] = val

	return m, nil
}

// Removes single quotes from start and end of value
// as this is handled by SQL parser
func removeSingleQuotes(s string) string {
	start := regexp.MustCompile(`(^['])`)
	end := regexp.MustCompile(`([']$)`)

	s = start.Split(s, -1)[1]
	s = end.Split(s, -1)[0]

	return s
}

// Checks for AND or OR logical operator
func resolveLogicalOperator(s string) string {
	r := regexp.MustCompile(`(\sAND|OR\s)`)

	lo := r.FindString(s)
	if lo == "" {
		lo = " AND "
	}

	return lo
}

// Checks for value like: Age=>'10'
// then splits into slices
func splitFilterQuery(s string) []string {
	var qs []string

	r := regexp.MustCompile(`([a-zA-Z]*[=><~!]{1,2}[']\S*['])`)

	for i := 0; i < 2; i++ {
		if r.MatchString(s) {
			q := r.FindString(s)
			qs = append(qs, q)
			s = strings.Split(s, q)[1]
		}

	}

	return qs
}

// Extracts values like: 'name'
// Compares to columns for match and returns col
func parseFilterField(s string, cols []string) (string, error) {
	r := regexp.MustCompile(`([a-zA-Z]*)`)

	v := r.FindString(s)

	f, err := validateField(v, cols)
	if err != nil {
		e := errors.New("No matching field")
		return "", e
	}

	return f, nil
}

// Extracts predicates/operators like '=' or '=>'
func parseFilterPredicate(s string) string {
	r := regexp.MustCompile(`([=><~!]{1,2})`)
	v := r.FindString(s)
	return v
}

// Extracts values encased in single quotes, strips and returns
func parseFilterValue(s string) string {
	r := regexp.MustCompile(`([']\S*['])`)
	v := r.FindString(s)
	v = removeSingleQuotes(v)
	return v
}

// If predicate is a '~' like, convert to SQL formatted statement
func evaluateLikeStatement(p string, v string) (string, string) {
	if p == "~" {
		p = " LIKE "
		v = "%" + v + "%"
	}

	return p, v
}

// Generates default values for filtering
func setFilterDefaults() ([]map[string]string, string) {

	d := map[string]string{
		"field":     "'1'",
		"predicate": "=",
		"value":     "1",
	}

	f := []map[string]string{
		d,
		d,
	}

	lo := " AND "

	return f, lo
}
