package visualizer

import (
	"fmt"

	"github.com/agneum/plan-exporter/pgscanner"
	"github.com/agneum/plan-exporter/visualizer/dalibo"
)

func New(visualizer string) (pgscanner.PlanExporter, error) {
	switch visualizer {
	case dalibo.VisualizerType:
		return dalibo.New(), nil
	}

	return nil, fmt.Errorf("unknown visualizer given %q", visualizer)
}
