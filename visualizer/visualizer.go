// Package visualizer provides an exporter factory.
package visualizer

import (
	"fmt"

	"github.com/agneum/plan-exporter/pgscanner"
	"github.com/agneum/plan-exporter/visualizer/dalibo"
	"github.com/agneum/plan-exporter/visualizer/depesz"
	"github.com/agneum/plan-exporter/visualizer/tensor"
)

// Visualizer has the following properties
type Visualizer struct {
	Target  string
	PostURL string
}

// New creates new Visualizer object
func New() *Visualizer {
	return &Visualizer{}
}

// BuildExporter buidls a new query plan exporter by visualizer type.
func (v Visualizer) BuildExporter() (pgscanner.PlanExporter, error) {
	switch v.Target {
	case dalibo.VisualizerType:
		return dalibo.New(v.PostURL), nil

	case depesz.VisualizerType:
		return depesz.New(v.PostURL), nil

	case tensor.VisualizerType:
		return tensor.New(v.PostURL), nil
	}

	return nil, fmt.Errorf("unknown visualizer given %q", v.Target)
}
