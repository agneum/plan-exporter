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
	postURL        = "https://explain.dalibo.com/new"
)

// Dalibo defines a query plan exporter for the Dalibo visualizer.
type Dalibo struct {
}

// New creates a new Dalibo exporter.
func New() *Dalibo {
	return &Dalibo{}
}

//  Export posts plan to a visualizer and returns link to the visualization plan page.
func (d *Dalibo) Export(plan string) (string, error) {
	formVal := url.Values{client.PlanKey: []string{plan}}

	explainURL, err := client.MakeRequest(postURL, formVal)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}

	return explainURL, nil
}
