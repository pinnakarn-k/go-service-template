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

	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Application", c.cfg.Integration.GetPostByID.Application)
	httpReq.Header.Set("Requester", c.cfg.Integration.GetPostByID.Requester)

	var post ClientPostResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		httpReq,
		&post,
	); err != nil {
		return nil, fmt.Errorf(
			"get post by id client: %w",
			err,
		)
	}

	return &post, nil
}
