package parameters

import (
	"net/url"
)

// Parses the url query values for sort, limit and offset settings otherwise returns defaults
func SortLimitOffset(q url.Values) map[string]string {
	d := map[string]string{
		"sort":   "sourcedId",
		"limit":  "100",
		"offset": "0",
	}

	for k, _ := range d {
		if v, ok := q[k]; ok {
			d[k] = v[0]
		}
	}

	return d
}
