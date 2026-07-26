package httptransport

import (
	"errors"
	"go-service-template/internal/apperror"
	"go-service-template/internal/middleware"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewErrorHandler(logger *slog.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		status := http.StatusInternalServerError

		response := apperror.Response{
			Code:    apperror.ErrInternal.Code,
			Message: apperror.ErrInternal.Message,
		}

		var appErr *apperror.Error

		if errors.As(err, &appErr) {
			status = appErr.Status

			response = apperror.Response{
				Code:    appErr.Code,
				Message: appErr.Message,
				Details: appErr.Details,
			}
		}

		requestID, _ := c.Locals(middleware.RequestIDLocal).(string)

		if status >= http.StatusInternalServerError {
			logger.Error(
				"request failed",
				"request_id", requestID,
				"error", err,
				"method", c.Method(),
				"path", c.Path(),
			)
		}

		return c.Status(status).JSON(response)
	}
}
