// Package depesz provides a query plan exporter for the Depesz visualizer.
package depesz

import (
	"fmt"
	"net/url"

	"github.com/agneum/plan-exporter/client"
)

// visualizer constants
const (
	VisualizerType = "depesz"
	visualizerURL  = "https://explain.depesz.com/"
)

// Depesz defines a query plan exporter for the Depesz visualizer.
type Depesz struct {
}

// New creates a new Depesz exporter.
func New() *Depesz {
	return &Depesz{}
}

//  Export posts plan to a visualizer and returns link to the visualization plan page.
func (d *Depesz) Export(plan string) (string, error) {
	formVal := url.Values{client.PlanKey: []string{plan}}

	explainURL, err := client.MakeRequest(visualizerURL, formVal)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}

	return explainURL, nil
}
