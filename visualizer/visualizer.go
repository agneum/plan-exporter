// Package visualizer provides an exporter factory.
package visualizer

import (
	"fmt"

	"github.com/agneum/plan-exporter/config"
	"github.com/agneum/plan-exporter/pgscanner"
	"github.com/agneum/plan-exporter/visualizer/dalibo"
	"github.com/agneum/plan-exporter/visualizer/depesz"
	"github.com/agneum/plan-exporter/visualizer/tensor"
)

// New creates a new query plan exporter by visualizer type.
func New(cfg *config.Config) (pgscanner.PlanExporter, error) {
	switch cfg.Target {
	case dalibo.VisualizerType:
		return dalibo.New(cfg.PostURL), nil

	case depesz.VisualizerType:
		return depesz.New(cfg.PostURL), nil

	case tensor.VisualizerType:
		return tensor.New(cfg.PostURL), nil
	}

	return nil, fmt.Errorf("unknown visualizer given %q", cfg.Target)
}
