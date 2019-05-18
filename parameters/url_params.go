package handlers

import (
	"fmt"
	"net/url"
)

func ParseUrlParams(u *url.URL, c []string) []map[string][]string {
	q := u.URL.Query()

	fields, err := ValidateFields(q, c)
	if err != nil {
		// TODO: return; status error, warning, invalid_selection_field
	}

	filters, logicalOp := ParseFilter(q, c)

	return q
}
