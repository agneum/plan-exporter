// Package pgscanner provides a scanner of psql output.
package pgscanner

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"strings"
)

// QueryPlanSubstringDetector defines the label for query plan output.
const QueryPlanSubstringDetector = "QUERY PLAN"

// minimalQueryPlanLines defines minimal lines to post a query plan.
const minimalQueryPlanLines = 2

// confirmationResponse defines the expected response to post.
const confirmationResponse = "Y"

// PlanExporter defines the interface to post query plans.
type PlanExporter interface {
	Export(string) (string, error)
}

// PgScanner provides a psql output scanner.
type PgScanner struct {
	cfg          *Config
	reader       io.Reader
	writer       io.Writer
	planExporter PlanExporter
	state        ScannerState
}

// Config contains scanner configuration.
type Config struct {
	AutoConfirm bool
}

// New creates a new Postgres scanner.
func New(cfg *Config, reader io.Reader, writer io.Writer, planExporter PlanExporter) *PgScanner {
	return &PgScanner{
		cfg:          cfg,
		reader:       reader,
		writer:       writer,
		planExporter: planExporter,
	}
}

// Run starts the Postgres scanner.
func (s *PgScanner) Run(ctx context.Context) {
	scanner := bufio.NewScanner(s.reader)

	s.state.Reset()

	// explainLines is a temporary buffer for plan results.
	explainLines := []string{}

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
		}

		text := scanner.Text()

		if s.state.mode == ConfirmingMode {
			// Check for the posting confirmation.
			if text == confirmationResponse {
				s.postPlan()
				continue
			}

			s.state.Reset()
		}

		if strings.Contains(text, QueryPlanSubstringDetector) {
			s.state.SetMode(ScanningMode)
			fmt.Println()
		}

		fmt.Println(text)

		if !s.state.IsMode(ScanningMode) {
			continue
		}

		if text == "" {
			s.state.SetMode(NormalMode)

			if len(explainLines) < minimalQueryPlanLines {
				log.Printf("Not enough lines in plan to post.")
				continue
			}

			s.state.SetBuffer(strings.Join(explainLines[2:len(explainLines)-1], "\n"))

			explainLines = []string{}
			s.state.mode = ConfirmingMode

			if s.cfg.AutoConfirm {
				s.postPlan()

				continue
			}

			_, _ = fmt.Fprintln(s.writer, "Do you want to post this plan to the visualizer?\nSend '\\qecho Y' to confirm")

			continue
		}

		explainLines = append(explainLines, text)
	}
}

func (s *PgScanner) postPlan() {
	_, _ = fmt.Fprintln(s.writer, "Posting to the visualizer...")
	defer s.state.Reset()

	if !s.state.IsMode(ConfirmingMode) {
		_, _ = fmt.Fprintf(s.writer, ("cannot post because scanner is in invalid mode\n"))
		return
	}

	if s.state.buffer == "" {
		_, _ = fmt.Fprintf(s.writer, ("no plan to export\n"))
		return
	}

	url, err := s.planExporter.Export(s.state.buffer)
	if err != nil {
		_, _ = fmt.Fprintf(s.writer, "Failed to post query plan: %s.\n", err)
		return
	}

	_, _ = fmt.Fprintf(s.writer, "The plan has been posted successfully.\nURL: %s\n", url)
}
