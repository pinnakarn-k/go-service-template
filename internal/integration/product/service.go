package product

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

func (s *Service) GetProducts(
	ctx context.Context,
) (*GetProductsResponse, error) {
	products, err := s.client.GetProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"get products service: %w",
			err,
		)
	}

	return mapGetProductsResponse(products), nil
}
