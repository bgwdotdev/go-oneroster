package handlers

import (
	"regexp"
	"strings"
)

// Takes a url query and converts it into map of filter parameters
func ParseFilters(q map[string][]string) ([]map[string]string, string) {
	// Set default output
	d := map[string]string{
		"field":     "'1'",
		"predicate": "=",
		"value":     "1",
	}
	ms := []map[string]string{
		d,
		d,
	}
	lo := " AND "
	if _, ok := q["filter"]; !ok {
		return ms, lo
	}

	s := q["filter"][0]
	// Checks for value like: Age=>'10'
	r := regexp.MustCompile(`([a-zA-Z]*[=><~!]{1,2}[']\S*['])`)
	// Checks for AND or OR logical operator
	r2 := regexp.MustCompile(`(\sAND|OR\s)`)
	lo = r2.FindString(s)
	if lo == "" {
		lo = " AND "
	}

	// Seperates filters
	var v string
	var vs []string
	i := 0
	for r.MatchString(s) && i < 2 {
		i++
		v = r.FindString(s)
		vs = append(vs, v)
		s = strings.Split(s, v)[1]
	}

	// Splits filters operations
	field := regexp.MustCompile(`([a-zA-Z]*)`)
	predicate := regexp.MustCompile(`([=><~!]{1,2})`)
	value := regexp.MustCompile(`([']\S*['])`)
	for i, v := range vs {
		m := make(map[string]string)
		// Parse parameters
		f := field.FindString(v)
		p := predicate.FindString(v)
		val := value.FindString(v)
		val = removeSingleQuotes(val)
		// Convert to SQL like statement
		if p == "~" {
			p = " LIKE "
			val = "%" + val + "%"
		}

		m["field"] = f
		m["predicate"] = p
		m["value"] = val
		ms[i] = m
	}
	return ms, lo
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
