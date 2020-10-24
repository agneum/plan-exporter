package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/agneum/plan-exporter/config"
	"github.com/agneum/plan-exporter/pgscanner"
	"github.com/agneum/plan-exporter/visualizer"
)

func main() {
	target := flag.String("target", "depesz", "type of an explain visualizer to export")
	postURL := flag.String("post-url", "", "absolute URL to HTML <form> element's `action`")
	autoConfirm := flag.Bool("auto-confirm", false, "send an execution plan automatically without additional confirmation")

	flag.Parse()

	ctx := context.Background()
	cfg := &config.Config{
		Target:  *target,
		PostURL: *postURL,
	}

	planner, err := visualizer.New(cfg)

	if err != nil {
		log.Fatalf("failed to init a query plan exporter: %v", err)
	}

	fmt.Println(generateWelcomeMessage(cfg, planner))

	scannerCfg := &pgscanner.Config{
		AutoConfirm: *autoConfirm,
	}

	pgScanner := pgscanner.New(scannerCfg, os.Stdin, os.Stdout, planner)
	pgScanner.Run(ctx)
}

func generateWelcomeMessage(cfg *config.Config, planner pgscanner.PlanExporter) string {
	welcome := strings.Builder{}

	welcome.WriteString("Welcome to the query plan exporter.")
	welcome.WriteString(fmt.Sprintf("\nTarget: %s", cfg.Target))
	welcome.WriteString(fmt.Sprintf("\nURL: %s", planner.Target()))

	return welcome.String()
}
