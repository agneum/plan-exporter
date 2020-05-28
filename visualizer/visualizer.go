// Package visualizer provides an exporter factory.
package visualizer

import (
	"fmt"

	"github.com/agneum/plan-exporter/pgscanner"
	"github.com/agneum/plan-exporter/visualizer/dalibo"
	"github.com/agneum/plan-exporter/visualizer/depesz"
	"github.com/agneum/plan-exporter/visualizer/tensor"
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
