package post

import (
	"go-service-template/internal/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetPosts(c *fiber.Ctx) error {
	posts, err := h.service.GetPosts(c.UserContext())
	if err != nil {
		return err
	}

	return response.OK(c, posts)
}

func (h *Handler) GetPostByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return ErrInvalidPostIDRequest
	}

	post, err := h.service.GetPostByID(
		c.UserContext(),
		id,
	)
	if err != nil {
		return err
	}

	return response.OK(c, post)
}
