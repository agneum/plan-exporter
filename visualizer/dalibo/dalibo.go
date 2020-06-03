// Package dalibo provides a query plan exporter for the Dalibo visualizer.
package dalibo

import (
	"fmt"
	"net/url"

	"github.com/agneum/plan-exporter/client"
)

// visualizer constants
const (
	VisualizerType = "dalibo"
	defaultPostURL = "https://explain.dalibo.com/new"
	planKey        = "plan"
)

// Dalibo defines a query plan exporter for the Dalibo visualizer.
type Dalibo struct {
	postURL string
}

// New creates a new Dalibo exporter.
func New(postURL string) *Dalibo {
	if postURL == "" {
		postURL = defaultPostURL
	}

	return &Dalibo{postURL: postURL}
}

// Export posts plan to a visualizer and returns link to the visualization plan page.
func (d *Dalibo) Export(plan string) (string, error) {
	formVal := url.Values{planKey: []string{plan}}

	explainURL, err := client.MakeRequest(d.postURL, formVal)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}

	return explainURL, nil
}
