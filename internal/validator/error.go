package validator

import (
	"go-service-template/internal/apperror"
	"net/http"
)

func newValidationError(details ...apperror.Detail) *apperror.Error {
	return &apperror.Error{
		Status:  http.StatusBadRequest,
		Code:    "VALIDATION_ERROR",
		Message: "request validation failed",
		Details: details,
	}
}
