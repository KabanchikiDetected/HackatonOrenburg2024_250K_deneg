package errors

import "errors"

var (
	InternalServerError = errors.New("internal server error")
	BadRequest          = errors.New("bad request")
	NotFound            = errors.New("not found")
)
