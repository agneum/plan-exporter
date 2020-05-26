// Package visualizer provides an exporter factory.
package visualizer

import (
	"fmt"
	"net/url"

	"github.com/agneum/plan-exporter/pgscanner"
	"github.com/agneum/plan-exporter/client"
)

type Visualizer struct {
  postURL string
  planKey string
}

// Export posts plan to a visualizer and returns link to the visualization plan page.
func (this *Visualizer) Export(plan string) (string, error) {
	formVal := url.Values{this.planKey: []string{plan}}

	explainURL, err := client.MakeRequest(this.postURL, formVal)
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

	default: //depesz
    return &Visualizer{
		        postURL: "https://explain.depesz.com/",
            planKey: "plan",
    }, nil
	}

	return nil, fmt.Errorf("unknown visualizer given %q", visualizer)
}
