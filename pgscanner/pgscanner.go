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

// PlanExporter defines the interface to post query plans.
type PlanExporter interface {
	Export(string) (string, error)
}

// PgScanner provides a psql output scanner.
type PgScanner struct {
	reader       io.Reader
	writer       io.Writer
	planExporter PlanExporter
}

// New creates a new Postgres scanner.
func New(reader io.Reader, writer io.Writer, planExporter PlanExporter) *PgScanner {
	return &PgScanner{reader: reader, writer: writer, planExporter: planExporter}
}

// Run starts the Postgres scanner.
func (s *PgScanner) Run(ctx context.Context) {
	scanner := bufio.NewScanner(s.reader)

	explainLines := []string{}
	isScanningMode := false

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
		}

		text := scanner.Text()

		if strings.Contains(text, QueryPlanSubstringDetector) {
			isScanningMode = true
		}

		if !isScanningMode {
			fmt.Println(text)
			continue
		}

		if text == "" {
			isScanningMode = false

			if len(explainLines) < minimalQueryPlanLines {
				log.Printf("Not enough lines in plan to post.")
				continue
			}

			plan := strings.Join(explainLines[2:len(explainLines)-1], "\n")

			_, _ = fmt.Fprintln(s.writer, "Posting to the visualizer...")

			explainLines = []string{}

			url, err := s.planExporter.Export(plan)
			if err != nil {
				log.Printf("Failed to post query plan: %s.", err)
				continue
			}

			_, _ = fmt.Fprintf(s.writer, "The plan has been posted successfully.\nURL: %s", url)

			continue
		}

		explainLines = append(explainLines, text)
	}
}
