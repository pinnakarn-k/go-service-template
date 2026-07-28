package httptransport

import (
	"go-service-template/internal/pagination"

	"github.com/gofiber/fiber/v2"
)

type Success[T any] struct {
	Data T `json:"data"`
}

func respond[T any](c *fiber.Ctx, status int, data T) error {
	return c.Status(status).JSON(Success[T]{
		Data: data,
	})
}

func OK[T any](c *fiber.Ctx, data T) error {
	return respond(c, fiber.StatusOK, data)
}

func Created[T any](c *fiber.Ctx, data T) error {
	return respond(c, fiber.StatusCreated, data)
}

func Accepted[T any](c *fiber.Ctx, data T) error {
	return respond(c, fiber.StatusAccepted, data)
}

func OKPage[T any](
	c *fiber.Ctx,
	response pagination.Response[T],
) error {
	return c.Status(fiber.StatusOK).JSON(response)
}
