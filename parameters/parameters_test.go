package parameters

import (
	"net/url"
	"testing"
)

type f struct {
	a string
	b string
}

func TestFilterAdd(t *testing.T) {
	// setup
	f1 := f{"a", "b"}
	f2 := f{"a", "b"}
	want := [2]f{f1, f2}

	// execute
	fd := defaults()
	got := [2]f{
		fd,
		fd,
	}

	//validate
	if got != want {
		t.Errorf("got: %v; want %v", got, want)
	}
}

func defaults() f {
	return f{
		a: "a",
		b: "b",
	}
}

func TestUrlGet(t *testing.T) {
	// setup
	q, _ := url.ParseQuery(`x=1&y=1`)
	want := "1"
	// execute
	qv := q.Get("x")
	var got string
	if qv != "" {
		got = qv
	}
	// validate
	if got != want {
		t.Errorf("got: %v; want %v", got, want)
	}
}
