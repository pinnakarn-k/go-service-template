package cart

import (
	"context"
	"fmt"
)

type Service struct {
	client *Client
}

func NewService(client *Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) GetCarts(
	ctx context.Context,
	req GetCartsRequest,
) (*GetCartsResponse, error) {
	carts, err := s.client.GetCarts(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"get carts service: %w",
			err,
		)
	}

	return mapGetCartsResponse(carts), nil
}
