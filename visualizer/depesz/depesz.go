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
	defaultPostURL = "https://explain.depesz.com/"
	planKey        = "plan"
)

// Depesz defines a query plan exporter for the Depesz visualizer.
type Depesz struct {
	postURL string
}

// New creates a new Depesz exporter.
func New(postURL string) *Depesz {
	if postURL == "" {
		postURL = defaultPostURL
	}

	return &Depesz{postURL: postURL}
}

// Export posts plan to a visualizer and returns link to the visualization plan page.
func (d *Depesz) Export(plan string) (string, error) {
	formVal := url.Values{planKey: []string{plan}}

	explainURL, err := client.MakeRequest(d.postURL, formVal)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}

	return explainURL, nil
}
