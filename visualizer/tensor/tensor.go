// Package tensor provides a query plan exporter for the Tensor visualizer.
package tensor

import (
	"fmt"
	"net/url"

	"github.com/agneum/plan-exporter/client"
)

// Visualizer constants.
const (
	VisualizerType = "tensor"
	defaultPostURL = "https://explain.tensor.ru/explain"
	planKey        = "plan"
)

// Tensor defines a query plan exporter for the Tensor visualizer.
type Tensor struct {
	postURL string
}

// New creates a new Tensor exporter.
func New(postURL string) *Tensor {
	if postURL == "" {
		postURL = defaultPostURL
	}

	return &Tensor{postURL: postURL}
}

// Export posts plan to a visualizer and returns link to the visualization plan page.
func (t *Tensor) Export(plan string) (string, error) {
	formVal := url.Values{planKey: []string{plan}}

	explainURL, err := client.MakeRequest(t.postURL, formVal)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}

	return explainURL, nil
}

// Target returns a post URL.
func (t *Tensor) Target() string {
	return t.postURL
}
