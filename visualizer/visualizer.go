// Package visualizer provides an exporter factory.
package visualizer

import (
	"fmt"

	"github.com/agneum/plan-exporter/pgscanner"
	"github.com/agneum/plan-exporter/visualizer/dalibo"
	"github.com/agneum/plan-exporter/visualizer/depesz"
)

// New creates a new query plan exporter by visualizer type.
func New(visualizer string) (pgscanner.PlanExporter, error) {
	switch visualizer {
	case dalibo.VisualizerType:
		return dalibo.New(), nil

	case depesz.VisualizerType:
		return depesz.New(), nil
	}

	return nil, fmt.Errorf("unknown visualizer given %q", visualizer)
}
