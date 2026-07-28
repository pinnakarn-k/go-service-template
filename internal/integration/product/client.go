package product

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

func (c *Client) GetProducts(
	ctx context.Context,
) (*ClientGetProductsResponse, error) {
	const operation = "get products client"

	url := c.cfg.Integration.GetProducts.URL

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
			Application: c.cfg.Integration.GetProducts.Application,
			Requester:   c.cfg.Integration.GetProducts.Requester,
			Key:         c.cfg.Integration.GetProducts.Key,
			Base64Mode:  c.cfg.Integration.GetProducts.Base64Mode,
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

	var products ClientGetProductsResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		httpReq,
		&products,
	); err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return &products, nil
}
