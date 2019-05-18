package parameters

import (
	"database/sql"
)

// Dynamically builds the format of the select query for JSON output
func FormatResults(rows *sql.Rows) map[string]interface{} {
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	out := make(map[string]interface{})

	cv := make([]interface{}, len(cols))
	cvp := make([]interface{}, len(cols))
	// Create pointer for row.Scan()
	for i, _ := range cv {
		cvp[i] = &cv[i]
	}

	err = rows.Scan(cvp...)
	if err != nil {
		panic(err)
	}

	for i, c := range cols {
		v := cv[i]
		out[c] = v
	}

	return out
}
