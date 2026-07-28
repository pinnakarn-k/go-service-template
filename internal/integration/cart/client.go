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
	const operation = "get carts client"

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

	headers, err := httpclient.GenerateTokenHeaders(
		httpclient.TokenConfig{
			Application: c.cfg.Integration.GetCarts.Application,
			Requester:   c.cfg.Integration.GetCarts.Requester,
			Key:         c.cfg.Integration.GetCarts.Key,
			Base64Mode:  c.cfg.Integration.GetCarts.Base64Mode,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Application", headers.Application)
	httpReq.Header.Set("Requester", headers.Requester)
	httpReq.Header.Set("Pretoken", headers.Pretoken)
	httpReq.Header.Set("Token", headers.Token)

	var carts ClientGetCartsResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		httpReq,
		&carts,
	); err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return &carts, nil
}
