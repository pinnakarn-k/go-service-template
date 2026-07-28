package post

import (
	"context"
	"fmt"
	"net/http"

	"go-service-template/internal/config"
	"go-service-template/internal/httpclient"
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

func (c *Client) GetPosts(
	ctx context.Context,
) ([]ClientPostResponse, error) {
	url := c.cfg.Integration.GetPosts.URL

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

	var posts []ClientPostResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		httpReq,
		&posts,
	); err != nil {
		return nil, fmt.Errorf(
			"get posts client: %w",
			err,
		)
	}

	return posts, nil
}

func (c *Client) GetPostByID(
	ctx context.Context,
	id int,
) (*ClientPostResponse, error) {
	const operation = "get post by id client"

	url := fmt.Sprintf(c.cfg.Integration.GetPostByID.URL, id)

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
			Application: c.cfg.Integration.GetPostByID.Application,
			Requester:   c.cfg.Integration.GetPostByID.Requester,
			Key:         c.cfg.Integration.GetPostByID.Key,
			Base64Mode:  c.cfg.Integration.GetPostByID.Base64Mode,
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

	var post ClientPostResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		httpReq,
		&post,
	); err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return &post, nil
}
