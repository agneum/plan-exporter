// Package client provides a http helper to post query plans.
package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// PlanKey defines the form key in a request for a query plan.
const PlanKey = "plan"

// MakeRequest performs the request and returns the redirected request URL.
func MakeRequest(targetURL string, formVal url.Values) (string, error) {
	response, err := http.PostForm(targetURL, formVal)
	if err != nil {
		return "", fmt.Errorf("failed to post form: %w", err)
	}

	if response.Body != nil {
		defer func() { _ = response.Body.Close() }()
	}

	return response.Request.URL.String(), nil
}
