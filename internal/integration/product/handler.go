package product

import (
	"github.com/gofiber/fiber/v2"

	"go-service-template/internal/httptransport"
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

	return httptransport.OK(c, products)
}
