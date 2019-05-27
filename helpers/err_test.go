package helpers

import (
	"encoding/json"
	//"fmt"
	"os"
)

type fatal interface {
	Fatal() bool
}

func IsFatal(e error) bool {
	t, ok := e.(fatal)
	return ok && t.Fatal()
}

type e struct {
	Err  string
	Code string
}

func (e *e) Error() string { return e.Err }

func (e *e) Fatal() bool {
	if e.Code > "1" {
		return true
	}
	return false
}

type p struct {
	Offset string
	Limit  string
}

func (p *p) offset(u map[string]string) error {
	v, err := validateOffset(u["Offset"])
	if err != nil {
		err.(*e).Code = "1"
		return err
	}

	p.Offset = v
	return nil
}

func validateOffset(s string) (string, error) {
	if s < "0" {
		return "0", &e{Err: "less than 0"}
	}
	return s, nil
}

func (p *p) limit(u map[string]string) error {
	if u["Limit"] < "10" {
		p.Limit = u["Limit"]
		return &e{"less than 10", "2"}
	}

	p.Limit = u["Limit"]
	return nil
}

type ErrorsPayload struct {
	Errors []error `json:"errors"`
}

func TestErr(t *testing.T) {
	errs := &ErrorsPayload{}
	d := map[string]string{"Offset": "-1", "Limit": "-8"}

	param := &p{}

	err := param.limit(d)
	if err != nil {
		errs.Errors = append(errs.Errors, err)
		if IsFatal(err) {
			b, err := json.Marshal(errs)
			if err != nil {
				panic(err)
			}
			os.Stdout.Write(b)
			return
		}
	}

	err = param.offset(d)
	if err != nil {
		errs.Errors = append(errs.Errors, err)
	}

	b, err := json.Marshal(errs)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(b)
}
