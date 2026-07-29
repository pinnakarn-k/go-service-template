package transaction

import (
	"go-service-template/internal/apperror"
	"net/http"
)

var ErrTransactionNotFound = &apperror.Error{
	Status:  http.StatusNotFound,
	Code:    "TRANSACTION_NOT_FOUND",
	Message: "transaction not found",
}
