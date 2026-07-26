package dummyjson

import (
	"context"
	"fmt"
	"go-service-template/internal/httpclient"
	"net/http"
	"strings"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(
	baseURL string,
	httpClient *http.Client,
) *Client {
	return &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
	}
}

func (c *Client) GetCarts(
	ctx context.Context,
) (*GetCartsClientResponse, error) {
	url := c.baseURL + "/carts"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create get carts request: %w",
			err,
		)
	}

	req.Header.Set("Accept", "application/json")

	var carts GetCartsClientResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		req,
		&carts,
	); err != nil {
		return nil, fmt.Errorf(
			"call dummyjson get carts: %w",
			err,
		)
	}

	return &carts, nil
}

func (c *Client) GetProducts(
	ctx context.Context,
) (*GetProductsClientResponse, error) {
	url := c.baseURL + "/products"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create get products request: %w",
			err,
		)
	}

	req.Header.Set("Accept", "application/json")

	var products GetProductsClientResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		req,
		&products,
	); err != nil {
		return nil, fmt.Errorf(
			"call dummyjson get products: %w",
			err,
		)
	}

	return &products, nil
}
