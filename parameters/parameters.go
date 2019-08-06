package parameters

import (
	"fmt"
	"github.com/fffnite/go-oneroster/helpers"
	log "github.com/sirupsen/logrus"
	"net/url"
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

func (p *Parameters) Resolve(q url.Values, col []string) ([]error, error) {
	var ep []error
	var err error
	p.SetLimit(q)
	p.SetOffset(q)
	err = p.SetSort(q, col)
	if err != nil {
		ep = append(ep, err)
	}
	err = p.SetFields(q, col)
	if err != nil {
		ep = append(ep, err)
	}
	err = p.SetFilters(q, col)
	if err != nil {
		ep = append(ep, err)
	}
	for _, e := range ep {
		if helpers.IsInvalid(e) {
			return ep, e
		}
	}
	return ep, nil
}

func (p *Parameters) SetSort(q url.Values, col []string) error {
	p.Sort = "sourcedId"
	qv := q.Get("sort")
	if qv != "" {
		v, err := validateField(qv, col)
		if err != nil {
			err.(*helpers.ErrorObject).CodeMinor = "invalid_sort_field"
			err.(*helpers.ErrorObject).Populate()
			return err
		}
		p.Sort = v
	}
	return nil
}

func (p *Parameters) SetLimit(q url.Values) {
	p.Limit = "100"
	qv := q.Get("limit")
	if qv != "" {
		p.Limit = qv
	}
}

func (p *Parameters) SetOffset(q url.Values) {
	p.Offset = "0"
	qv := q.Get("offset")
	if qv != "" {
		p.Offset = qv
	}
}

func (p *Parameters) SetFields(q url.Values, col []string) error {
	p.Fields = strings.Join(col, ", ")

	qv := q.Get("fields")
	if qv != "" {
		fs := strings.Split(qv, ",")
		var cols []string
		for _, f := range fs {
			c, err := validateField(f, col)
			if err != nil {
				err.(*helpers.ErrorObject).CodeMinor = "invalid_selection_field"
				err.(*helpers.ErrorObject).Populate()
				return err
			}
			cols = append(cols, c)
		}
		p.Fields = strings.Join(cols, ",")
	}

	return nil
}

func (p *Parameters) SetFilters(q url.Values, col []string) error {
	fltd := defaultFilter()
	flts := [2]filter{fltd, fltd}
	p.LogicalOperator = " AND "
	qv := q.Get("filter")
	if qv != "" {
		op := resolveLogicalOperator(qv)
		fs := splitFilterQuery(qv)
		for i, f := range fs {
			var flt filter
			err := flt.field(f, col)
			if err != nil {
				return err
			}
			flt.predicate(f)
			flt.value(f)
			flt.like()
			flts[i] = flt
		}
		p.LogicalOperator = op
	}
	p.Filter1 = flts[0]
	p.Filter2 = flts[1]
	return nil
}

// Compares the field against a list of column names
// Returns column name on match
func validateField(s string, col []string) (string, error) {
	var f string
	for _, c := range col {
		if s == c {
			f = c
		}
	}
	if f == "" {
		err := fmt.Sprintf("Unknown field: %v", s)
		log.Info(err)
		return "", &helpers.ErrorObject{Description: err}
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
		err.(*helpers.ErrorObject).CodeMinor = "invalid_filter_field"
		err.(*helpers.ErrorObject).Populate()
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
