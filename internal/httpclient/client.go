package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const maxErrorBodySize = 4 * 1024

func New(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

func DoJSON(
	client *http.Client,
	req *http.Request,
	out any,
) error {
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("execute http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK ||
		resp.StatusCode >= http.StatusMultipleChoices {

		body, readErr := io.ReadAll(
			io.LimitReader(resp.Body, maxErrorBodySize),
		)
		if readErr != nil {
			return fmt.Errorf(
				"unexpected http status %d: read response body: %w",
				resp.StatusCode,
				readErr,
			)
		}

		return fmt.Errorf(
			"unexpected http status %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	if out == nil || resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}

	return nil
}
