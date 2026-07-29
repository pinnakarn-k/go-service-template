package transaction

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"go-service-template/internal/apperror"
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

func (h *Handler) SearchTransactions(
	c *fiber.Ctx,
) error {
	var request SearchTransactionsRequest

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := h.validator.Validate(request); err != nil {
		return err
	}

	transactions, err := h.service.SearchTransactions(
		c.UserContext(),
		request,
	)
	if err != nil {
		return err
	}

	return httptransport.OKPage(c, *transactions)
}

func (h *Handler) GetTransactionByID(
	c *fiber.Ctx,
) error {
	id, err := strconv.ParseInt(
		c.Params("id"),
		10,
		64,
	)
	if err != nil {
		return apperror.NewInvalidPathParameterError(err)
	}

	if id <= 0 {
		return apperror.ErrInvalidID
	}

	transaction, err := h.service.GetTransactionByID(
		c.UserContext(),
		id,
	)
	if err != nil {
		return err
	}

	return httptransport.OK(
		c,
		transaction,
	)
}
