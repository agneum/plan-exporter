// Package tensor provides a query plan exporter for the Tensor visualizer.
package tensor

import (
	"fmt"
	"net/url"

	"github.com/agneum/plan-exporter/client"
)

// visualizer constants
const (
	VisualizerType = "tensor"
	visualizerURL  = "https://explain.tensor.ru/explain"
	planKey        = "explain"
)

// Tensor defines a query plan exporter for the Tensor visualizer.
type Tensor struct {
}

// New creates a new Tensor exporter.
func New() *Tensor {
	return &Tensor{}
}

// Export posts plan to a visualizer and returns link to the visualization plan page.
func (d *Tensor) Export(plan string) (string, error) {
	formVal := url.Values{planKey: []string{plan}}

	explainURL, err := client.MakeRequest(visualizerURL, formVal)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}

	return explainURL, nil
}
