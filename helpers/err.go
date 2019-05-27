package helpers

import ()

type invalid interface {
	Invalid() bool
}

func IsInvalid(e error) bool {
	t, ok := e.(invalid)
	return ok && t.Invalid()
}

type ErrorObject struct {
	StatusCode  int    `json:"-"`
	CodeMajor   string `json:"codeMajor"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	CodeMinor   string `json:"codeMinor"`
	IsInvalid   bool   `json:"-"`
}

func (e *ErrorObject) Error() string { return e.Description }

func (e *ErrorObject) Invalid() bool { return e.IsInvalid }

func (e *ErrorObject) Populate() {
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
	default:
		e.StatusCode = http.StatusInternalServerError
		e.CodeMajor = "failure"
		e.Severity = "erorr"
	}
	switch e.CodeMajor {
	case "failure":
		e.IsInvalid = true
	default:
		e.IsInvalid = false
	}
}
