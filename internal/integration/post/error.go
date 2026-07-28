package post

import (
	"go-service-template/internal/apperror"
	"net/http"
)

var (
	ErrInvalidPostIDRequest = &apperror.Error{
		Status:  http.StatusBadRequest,
		Code:    "INVALID_POST_ID",
		Message: "invalid post id",
	}
)
