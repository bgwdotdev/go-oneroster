package parameters

import (
	"net/url"
)

type Parameters struct {
	Sort            string
	Limit           string
	Offset          string
	Fields          string
	Filter1         filter
	LogicalOperator string
	Filter2         filter
}

type filter struct {
	Field     string
	Predicate string
	Value     string
}

// Parses parameter values from the query and returns structured, validated data
func ParseUrl(u *url.URL, c []string) Parameters {
	q := u.Query()

	slo := SortLimitOffset(q, c)
	fields, err := ValidateFields(q, c)
	if err != nil {
		// TODO: return; status error, warning, invalid_selection_field
	}
	filters, logicalOp := ParseFilters(q, c)

	var fs []filter
	for _, v := range filters {
		f := filter{
			v["field"],
			v["predicate"],
			v["value"],
		}
		fs = append(fs, f)
	}
	p := Parameters{
		Sort:            slo["sort"],
		Limit:           slo["limit"],
		Offset:          slo["offset"],
		Fields:          fields,
		Filter1:         fs[0],
		LogicalOperator: logicalOp,
		Filter2:         fs[1],
	}

	return p
}
