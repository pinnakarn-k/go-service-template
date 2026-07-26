package jsonplaceholder

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go-service-template/internal/httpclient"
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

func (c *Client) GetPosts(
	ctx context.Context,
) ([]PostClientResponse, error) {
	url := c.baseURL + "/posts"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create get posts request: %w",
			err,
		)
	}

	req.Header.Set("Accept", "application/json")

	var posts []PostClientResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		req,
		&posts,
	); err != nil {
		return nil, fmt.Errorf(
			"call jsonplaceholder get posts: %w",
			err,
		)
	}

	return posts, nil
}

func (c *Client) GetPostByID(
	ctx context.Context,
	id int,
) (*PostClientResponse, error) {
	url := c.baseURL + "/posts/" + strconv.Itoa(id)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create get post by id request: %w",
			err,
		)
	}

	req.Header.Set("Accept", "application/json")

	var post PostClientResponse

	if err := httpclient.DoJSON(
		c.httpClient,
		req,
		&post,
	); err != nil {
		return nil, fmt.Errorf(
			"call jsonplaceholder get post by id: %w",
			err,
		)
	}

	return &post, nil
}
