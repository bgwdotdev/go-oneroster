package handlers

import (
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
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got %v", want, got)
	}
}
