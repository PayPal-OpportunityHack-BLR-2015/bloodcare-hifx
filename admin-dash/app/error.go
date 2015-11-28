package app

import (
	"fmt"
)

// Err is the struct that holds system error
// messages and HTTP status code
type Err struct {
	ID         int
	Message    string
	HTTPStatus int
	ErrorMsg   string
	Strace     string
}

func (e Err) Error() string {
	return fmt.Sprintf("bloodcare.dash.err: %d: %s: %s", e.ID, e.Message, e.ErrorMsg)
}

// NewErr returns an app Error Instance
func NewErr(id int, message string, code int, err string) *Err {
	return &Err{ID: id, Message: message, HTTPStatus: code, ErrorMsg: err}
}

//SetErr returns a new object after setting the error string
func (e Err) SetErr(errStr string) *Err {
	e.ErrorMsg = errStr
	return &e
}

var (
	// InternalServerError denotes internal server errors.
	InternalServerError = &Err{ID: 1000, Message: "Internal server error", HTTPStatus: 500}
	// NotFoundError denotes content not found errors.
	NotFoundError = &Err{ID: 1002, Message: "Requested content not found", HTTPStatus: 404}
	// InvalidParametersError denotes invalid parametrs errors (what else can it be ;))
	InvalidParametersError = &Err{ID: 1003, Message: "Invalid request parameters", HTTPStatus: 400}
)
