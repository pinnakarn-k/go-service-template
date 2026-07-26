package auth

import (
	"go-service-template/internal/apperror"
	"net/http"
)

var (
	ErrMissingAuthCookie = &apperror.Error{
		Status:  http.StatusUnauthorized,
		Code:    "MISSING_AUTH_COOKIE",
		Message: "missing authentication cookie",
	}

	ErrInvalidToken = &apperror.Error{
		Status:  http.StatusUnauthorized,
		Code:    "INVALID_TOKEN",
		Message: "invalid token",
	}

	ErrTokenExpired = &apperror.Error{
		Status:  http.StatusUnauthorized,
		Code:    "TOKEN_EXPIRED",
		Message: "token expired",
	}

	ErrCustCodeNotFound = &apperror.Error{
		Status:  http.StatusUnauthorized,
		Code:    "CUST_CODE_NOT_FOUND",
		Message: "custCode not found",
	}

	ErrAccountIDNotFound = &apperror.Error{
		Status:  http.StatusUnauthorized,
		Code:    "ACCOUNT_ID_NOT_FOUND",
		Message: "accountId not found",
	}
)
