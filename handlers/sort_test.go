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
