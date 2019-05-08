package handlers

import (
	"net/url"
)

func SortQuery(q url.Values) string {
	var s string

	if v, ok := q["sort"]; ok {
		s = v[0]
	} else {
		s = "sourcedId"
	}

	return s
}
