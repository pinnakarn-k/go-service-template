package apperror

import (
	"fmt"
)

type Error struct {
	Status  int
	Code    string
	Message string
	Details []Detail
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Code, e.Cause)
	}

	return e.Code
}

func (e *Error) Unwrap() error {
	return e.Cause
}
