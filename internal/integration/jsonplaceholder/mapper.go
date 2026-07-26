package jsonplaceholder

func mapPostsResponse(
	src []PostClientResponse,
) []PostResponse {
	items := make([]PostResponse, 0, len(src))

	for _, post := range src {
		//nolint:gosimple
		items = append(items, PostResponse{
			UserID: post.UserID,
			ID:     post.ID,
			Title:  post.Title,
			Body:   post.Body,
		})
	}

	return items
}

func mapPostResponse(
	src *PostClientResponse,
) *PostResponse {
	return &PostResponse{
		UserID: src.UserID,
		ID:     src.ID,
		Title:  src.Title,
		Body:   src.Body,
	}
}
