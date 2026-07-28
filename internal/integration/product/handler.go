package product

import (
	"go-service-template/internal/response"

	"github.com/gofiber/fiber/v2"

	appvalidator "go-service-template/internal/validator"
)

type Handler struct {
	service   *Service
	validator *appvalidator.Validator
}

func NewHandler(
	service *Service,
	validator *appvalidator.Validator,
) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) GetProducts(c *fiber.Ctx) error {
	products, err := h.service.GetProducts(
		c.UserContext(),
	)
	if err != nil {
		return err
	}

	return response.OK(c, products)
}
