package httptransport

import "github.com/gofiber/fiber/v2"

type Success[T any] struct {
	Data T `json:"data"`
}

func OK[T any](c *fiber.Ctx, data T) error {
	return c.Status(fiber.StatusOK).JSON(Success[T]{
		Data: data,
	})
}
