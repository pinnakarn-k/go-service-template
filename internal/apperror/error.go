package apperror

import "net/http"

var (
	ErrUnauthorized = &Error{
		Status:  http.StatusUnauthorized,
		Code:    "UNAUTHORIZED",
		Message: "unauthorized",
	}

	ErrForbidden = &Error{
		Status:  http.StatusForbidden,
		Code:    "FORBIDDEN",
		Message: "forbidden",
	}

	ErrNotFound = &Error{
		Status:  http.StatusNotFound,
		Code:    "NOT_FOUND",
		Message: "resource not found",
	}

	ErrInternal = &Error{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_ERROR",
		Message: "internal server error",
	}
)

func NewInvalidQueryError(cause error) *Error {
	return &Error{
		Status:  http.StatusBadRequest,
		Code:    "INVALID_QUERY",
		Message: "invalid query parameters",
		Cause:   cause,
	}
}
