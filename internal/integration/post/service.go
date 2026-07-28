package post

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

func (s *Service) GetPosts(
	ctx context.Context,
) ([]PostResponse, error) {
	posts, err := s.client.GetPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"get posts service: %w",
			err,
		)
	}

	return mapPostsResponse(posts), nil
}

func (s *Service) GetPostByID(
	ctx context.Context,
	id int,
) (*PostResponse, error) {
	post, err := s.client.GetPostByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf(
			"get post by id service: %w",
			err,
		)
	}

	return mapPostResponse(post), nil
}
