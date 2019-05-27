package main

import (
	"regexp"
	"strings"
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

func (p *Parameters) Sort(q url.Values, col []string) error {
	p.Sort = "sourcedId"
	v, err := validateField(q, col)
	if err != nil {
		err.(*ErrorObject).CodeMinor = "invalid_sort_field"
		return err
	}

	p.Sort = v
	return nil
}

func (p *Parameters) Limit(q url.Values) {
	p.Limit = "100"
	p.Limit = q
}

func (p *Parameters) Offset(q url.Values) {
	p.Offset = "0"
	p.Offset = q
}

func (p *Parameters) Fields(q url.Values, col []string) error {
	p.Fields = strings.Join(col, ", ")
	fs := string.Split(q, ",")
	for _, f := range fs {
		c, err := validateField(f, col)
		if err != nil {
			err.(*ErrorObject).CodeMinor = "invalid_select_field"
			return err
		}
		cols = append(cols, c)
	}
	p.Fields = strings.Join(cols, ",")
	return nil
}

func (p *Parameters) Filters(q url.Values, col []string) error {
	op := resolveLogicalOperator(q)
	fs := splitFilterQuery(q)
	fltd := defaultFilter()
	flts := [2]filter{fltd, fltd}
	for i, f := range fs {
		var flt filter
		err := flt.field(f, cols)
		if err != nil {
			return err
		}
		flt.predicate(f)
		flt.value(f)
		flt.like(f)
		flts[i] = flt
	}
	p.LogicalOperator = op
	p.Filter1 = flts[0]
	p.Filter2 = flts[1]
}

// Compares the field against a list of column names
// Returns column name on match
func validateFields(s string, col []string) (string, error) {
	var f string
	for _, c := range col {
		if s == c {
			f = c
		}
	}
	if f == "" {
		err := fmt.Sprintf("Unknown field: %v", s)
		return "", &ErrorObject{Description: err}
	}
	return f, nil
}

func resolveLogicalOperator(s string) string {
	r := regexp.MustCompile(`(\sAND|OR\s)`)
	op := r.FindString(s)
	if op == "" {
		op = " AND "
	}
	return op
}

// splits up to 2 queries up into seperate slices
// "name='bob' age=>'18'
func splitFilterQuery(s string) []string {
	r := regexp.MustCompile(`([a-zA-Z]*[=><~!]{1,2}[']\S*['])`)
	var fs []string
	for i := 0; i < 2; i++ {
		if r.MatchString(s) {
			f := r.FindString(s)
			fs = append(fs, f)
			s = strings.Split(s, f)[1]
		}
	}
	return fs
}

func defaultFilter() filter {
	return filter{
		Field:     "'1'",
		Predicate: "=",
		Value:     "1",
	}
}

// Resolves the field name in the filter query
func (f *filter) field(s string, col []string) error {
	r := regexp.MustCompile(`([a-zA-Z]*)`)
	v := r.FindString(s)
	fld, err := validateField(v, col)
	if err != nil {
		err.(*ErrorObject).CodeMinor = "invalid_filter_field"
		return err
	}
	f.Field = fld
	return nil
}

func (f *filter) predicate(s string) {
	r := regexp.MustCompile(`([=><~!]{1,2})`)
	f.Predicate = r.FindString(s)
}

func (f *filter) value(s string) {
	r := regexp.MustCompile(`([']\S*['])`)
	f.Value = removeSingleQuotes(r.FindString(s))
}

func removeSingleQuotes(s string) string {
	b := regexp.MustCompile(`(^['])`)
	e := regexp.MustCompile(`([']$)`)
	s = b.Split(s, -1)[1]
	s = e.Split(s, -1)[0]
	return s
}

func (f *filter) like() {
	if f.Predicate == "~" {
		f.Predicate = " LIKE "
		f.Value = "%" + f.Value + "%"
	}
}
