package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	RequestIDHeader = "X-Request-ID"
	RequestIDLocal  = "request_id"
)

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get(RequestIDHeader)

		if requestID == "" {
			requestID = uuid.NewString()
		}

		c.Locals(RequestIDLocal, requestID)
		c.Set(RequestIDHeader, requestID)

		return c.Next()
	}
}
