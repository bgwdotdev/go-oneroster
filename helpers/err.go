package helpers

import (
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type invalid interface {
	Invalid() bool
}

func IsInvalid(e error) bool {
	ie, ok := e.(invalid)
	return ok && ie.Invalid()
}

type Error struct {
	StatusCode  int    `json:"-"`
	CodeMajor   string `json:"codeMajor"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	CodeMinor   string `json:"codeMinor"`
	IsInvalid   bool   `json:"-"`
}

func (e *Error) Error() string { return e.Description }

func (e *Error) Invalid() bool { return e.IsInvalid }

func (e *Error) Payload(w http.ResponseWriter, r *http.Request) {
	e = validateCodeMinor(e)
	log.Info(e)
	render.Status(r, e.StatusCode)
	render.JSON(w, r, e)
}

func validateCodeMinor(e *Error) *Error {
	switch e.CodeMinor {
	case "full success":
		e.StatusCode = http.StatusOK
		e.CodeMajor = "success"
		e.Severity = "status"
	case "invalid_sort_field", "invalid_selection_field":
		e.StatusCode = http.StatusOK
		e.CodeMajor = "success"
		e.Severity = "warning"
	case "invalid data", "invalid_filter_field", "invalid_blank_selection_field":
		e.StatusCode = http.StatusBadRequest
		e.CodeMajor = "failure"
		e.Severity = "error"
	case "unauthorized":
		e.StatusCode = http.StatusUnauthorized
		e.CodeMajor = "failure"
		e.Severity = "error"
	case "forbidden":
		e.StatusCode = http.StatusForbidden
		e.CodeMajor = "failure"
		e.Severity = "error"
	case "unknown object":
		e.StatusCode = http.StatusNotFound
		e.CodeMajor = "failure"
		e.Severity = "error"
	case "server_busy":
		e.StatusCode = http.StatusTooManyRequests
		e.CodeMajor = "failure"
		e.Severity = "error"
	}

	switch e.CodeMajor {
	case "failure":
		e.IsInvalid == true
	default:
		e.IsInvalid == false
	}

	return e
}
