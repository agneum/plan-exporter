// Package client provides a http helper to post query plans.
package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// MakeRequest performs request and returns the redirected request URL.
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

// MakeRequestWithBody performs request and returns response.
func MakeRequestWithBody(targetURL string, formVal url.Values) (io.ReadCloser, error) {
	response, err := http.PostForm(targetURL, formVal) // nolint:gosec
	if err != nil {
		return nil, fmt.Errorf("failed to post form: %w", err)
	}

	return response.Body, nil
}
