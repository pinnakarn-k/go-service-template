package middleware

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RequestLogger(logger *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		requestID, _ := c.Locals(RequestIDLocal).(string)

		logger.Info(
			"request completed",
			"request_id", requestID,
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"duration_ms", duration.Milliseconds(),
		)

		return err
	}
}
