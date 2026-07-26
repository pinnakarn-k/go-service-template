package dummyjson

import "fmt"

type ClientError struct {
	StatusCode int
	Body       string
}

func (e *ClientError) Error() string {
	return fmt.Sprintf(
		"dummyjson returned status %d: %s",
		e.StatusCode,
		e.Body,
	)
}
