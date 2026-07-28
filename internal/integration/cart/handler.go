package cart

import (
	"go-service-template/internal/apperror"
	"go-service-template/internal/httptransport"

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

func (h *Handler) GetCarts(c *fiber.Ctx) error {
	var req GetCartsRequest

	if err := c.QueryParser(&req); err != nil {
		return apperror.NewInvalidQueryError(err)
	}

	if err := h.validator.Validate(req); err != nil {
		return err
	}

	carts, err := h.service.GetCarts(
		c.UserContext(),
		req,
	)
	if err != nil {
		return err
	}

	return httptransport.OK(c, carts)
}
