package parameters

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

// compares a map of columns names against a query and returns columns if match
func TestValidateFields(t *testing.T) {
	// setup
	q := map[string][]string{
		"fields": []string{"a,b,c"},
	}

	c := []string{"a", "b", "c"}

	// execute
	var columns []string
	v := q["fields"][0]
	r := strings.Split(v, ",")
	for _, s := range r {
		col, err := validateField(c, s)
		if err != nil {
			panic(err)
		}
		columns = append(columns, col)
	}
	output := strings.Join(columns, ", ")

	//validate
	got := output
	want := "a, b, c"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got %v", want, got)
	}
}

// takes a query item and compares against columns, if no match, return error
func validateField(c []string, s string) (string, error) {

	// execute
	var output string
	for _, cv := range c {
		if cv == s {
			output = cv
		}
	}
	if output == "" {
		err := errors.New("No field match")
		return s, err
	} else {
		return output, nil
	}
}
