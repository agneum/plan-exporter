package pgscanner

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"strings"
)

const QueryPlanSubstringDetector = "QUERY PLAN"

type PlanExporter interface {
	Export(string) (string, error)
}

type PgScanner struct {
	reader       io.Reader
	writer       io.Writer
	planExporter PlanExporter
}

func New(reader io.Reader, writer io.Writer, planExporter PlanExporter) *PgScanner {
	return &PgScanner{reader: reader, writer: writer, planExporter: planExporter}
}

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

			if len(explainLines) < 2 {
				log.Printf("Not enough lines in plan to post.")
				continue
			}

			plan := strings.Join(explainLines[2:len(explainLines)-1], "\n")

			fmt.Fprintln(s.writer, "Posting to the visualizer...")

			explainLines = []string{}

			url, err := s.planExporter.Export(plan)
			if err != nil {
				log.Printf("Failed to post query plan: %s.", err)
				continue
			}

			fmt.Fprintf(s.writer, "The plan has been posted successfully.\nURL: %s", url)
			continue
		}

		explainLines = append(explainLines, text)
	}
}
