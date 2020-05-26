// Package visualizer provides an exporter factory.
package visualizer

import (
	"fmt"
	"net/url"

	"github.com/agneum/plan-exporter/client"
	"github.com/agneum/plan-exporter/pgscanner"
)

// Visualizer consists of a URL for HTTP POST call, along with HTML element name
type Visualizer struct {
	postURL string
	planKey string
}

// Export posts plan to a visualizer and returns link to the visualization plan page.
func (v *Visualizer) Export(plan string) (string, error) {
	formVal := url.Values{v.planKey: []string{plan}}

	explainURL, err := client.MakeRequest(v.postURL, formVal)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}

	return explainURL, nil
}

// New creates a new query plan exporter
func New(visualizer string) (pgscanner.PlanExporter, error) {
	switch visualizer {
	case "dalibo":
		return &Visualizer{
			postURL: "https://explain.dalibo.com/new",
			planKey: "plan",
		}, nil

	case "tensor":
		return &Visualizer{
			postURL: "https://explain.tensor.ru/explain",
			planKey: "explain",
		}, nil

	case "depesz":
		return &Visualizer{
			postURL: "https://explain.depesz.com/",
			planKey: "plan",
		}, nil
	}

	return nil, fmt.Errorf("unknown visualizer given %q", visualizer)
}
