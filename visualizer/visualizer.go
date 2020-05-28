// Package visualizer provides an exporter factory.
package visualizer

import (
	"fmt"
	"net/url"

	"github.com/agneum/plan-exporter/client"
	"github.com/agneum/plan-exporter/pgscanner"
)

// New creates a new query plan exporter by visualizer type.
func New(visualizer string, url string) (pgscanner.PlanExporter, error) {
	switch visualizer {
	case dalibo.VisualizerType:
		return dalibo.New(url), nil

	case depesz.VisualizerType:
		return depesz.New(url), nil

	case tensor.VisualizerType:
		return tensor.New(url), nil
	}

	return nil, fmt.Errorf("unknown visualizer given %q", visualizer)
}
