package cart

import (
	"context"
	"fmt"
	"go-service-template/internal/config"
	"go-service-template/internal/httpclient"
	"net/http"
)

type Client struct {
	cfg        *config.Config
	httpClient *http.Client
}

func NewClient(
	cfg *config.Config,
	httpClient *http.Client,
) *Client {
	return &Client{
		cfg:        cfg,
		httpClient: httpClient,
	}
}

func (c *Client) GetCarts(
	ctx context.Context,
) (*ClientGetCartsResponse, error) {
	url := c.cfg.Integration.GetCarts.URL

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Application", c.cfg.Integration.GetPosts.Application)
	httpReq.Header.Set("Requester", c.cfg.Integration.GetPosts.Requester)

	var carts ClientGetCartsResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		httpReq,
		&carts,
	); err != nil {
		return nil, fmt.Errorf(
			"get carts client: %w",
			err,
		)
	}

	return &carts, nil
}
