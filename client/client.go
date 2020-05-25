// Package client provides a http helper to post query plans.
package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// MakeRequest performs the request and returns the redirected request URL.
func MakeRequest(targetURL string, formVal url.Values) (string, error) {
	response, err := http.PostForm(targetURL, formVal) // nolint:gosec
	if err != nil {
		return "", fmt.Errorf("failed to post form: %w", err)
	}

	if response.Body != nil {
		defer func() { _ = response.Body.Close() }()
	}

	return response.Request.URL.String(), nil
}
