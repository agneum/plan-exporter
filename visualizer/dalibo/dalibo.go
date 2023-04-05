// Package dalibo provides a query plan exporter for the Dalibo visualizer.
package dalibo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/agneum/plan-exporter/client"
)

// Visualizer constants.
const (
	VisualizerType    = "dalibo"
	defaultPostURL    = "https://explain.dalibo.com"
	newPlanRoute      = "new.json"
	planResponseRoute = "plan"

	planKey  = "plan"
	titleKey = "title"
	queryKey = "query"
)

// Dalibo defines a query plan exporter for the Dalibo visualizer.
type Dalibo struct {
	postURL string
}

// PlanResponse represents response body.
type PlanResponse struct {
	ID        string `json:"id"`
	DeleteKey string `json:"deleteKey"`
}

// New creates a new Dalibo exporter.
func New(postURL string) *Dalibo {
	if postURL == "" {
		postURL = defaultPostURL
	}

	return &Dalibo{
		postURL: postURL,
	}
}

// Target returns a post URL.
func (d *Dalibo) Target() string {
	return d.postURL
}

// Export posts plan to a visualizer and returns link to the visualization plan page.
func (d *Dalibo) Export(plan string) (string, error) {
	formVal := url.Values{
		planKey:  []string{strings.TrimSpace(plan)},
		titleKey: []string{fmt.Sprintf("Plan created on %s", time.Now().Format(time.RFC1123))},
		queryKey: []string{""},
	}

	requestURL, err := url.JoinPath(d.postURL, newPlanRoute)
	if err != nil {
		return "", fmt.Errorf("failed to build request URL: %w", err)
	}

	planBody, err := client.MakeRequestWithBody(requestURL, formVal)
	if err != nil {
		return "", fmt.Errorf("failed to build request URL: %w", err)
	}

	defer func() { _ = planBody.Close() }()

	planResponse, err := d.parseBody(planBody)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}

	return d.buildShareLink(planResponse.ID)
}

func (d *Dalibo) parseBody(body io.ReadCloser) (*PlanResponse, error) {
	var plan PlanResponse

	if err := json.NewDecoder(body).Decode(&plan); err != nil {
		return nil, err
	}

	return &plan, nil
}

func (d *Dalibo) buildShareLink(id string) (string, error) {
	return url.JoinPath(d.postURL, planResponseRoute, id)
}
