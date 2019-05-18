package parameters

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestSortQuery(t *testing.T) {
	// setup
	q := make(map[string][]string)
	q["sort"] = []string{"name"}
	var s string

	// Execute
	if v, ok := q["sort"]; ok {
		fmt.Println(v)
		s = v[0]
	} else {
		s = "Id"
	}
	statement := fmt.Sprintf("SELECT id FROM table ORDER BY '%v'", s)

	// Validation
	got := statement
	want := "SELECT id FROM table ORDER BY 'name'"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got %v", want, got)
	}

	// Cleanup
}

func TestQuery(t *testing.T) {
	// setup
	q := map[string][]string{
		"sort":   []string{"name"},
		"limit":  []string{"50"},
		"offset": []string{"10"},
		"random": []string{"value"},
		"filter": []string{"id='2'"},
	}

	d := map[string]string{
		"sort":   "sourcedId",
		"limit":  "100",
		"offset": "0",
		"filter": "'1' = '1'",
		"field":  "sourcedId,name",
	}

	// Execution
	for k, _ := range d {
		if v, ok := q[k]; ok {
			d[k] = v[0]
		}
	}

	// validation
	got := d
	want := map[string]string{
		"sort":   "name",
		"limit":  "50",
		"offset": "10",
		"filter": "id='2'",
		"field":  "sourcedId,name",
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got %v", want, got)
	}
}

func TestFormatResults(t *testing.T) {
	// setup
	// Results from an SQL row
	cols := []string{"id", "sourcedId"}
	rows := struct {
		id        string
		sourcedId string
	}{
		"1",
		"ab-2",
	}

	// execute
	// results store
	r := make(map[string]interface{})
	// column
	cv := make([]interface{}, len(cols))
	// column pointer
	cvp := make([]interface{}, len(cols))
	// assign col to point
	for i, _ := range cv {
		cvp[i] = &cv[i]
	}

	// rows.Scan(&cvp)
	cvp[0] = rows.id
	cvp[1] = rows.sourcedId

	// set result col name and value
	for i, c := range cols {
		v := cvp[i]
		r[c] = v
	}

	_, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	// validate
	got := r
	want := map[string]interface{}{
		"id":        "1",
		"sourcedId": "ab-2",
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got %v", want, got)
	}

	// cleanup

}
