package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

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
	planner, err := visualizer.New(&config.Config{
		Target:  *target,
		PostURL: *postURL,
	})

	if err != nil {
		log.Fatalf("failed to init a query plan exporter: %v", err)
	}

	fmt.Printf("Welcome to the query plan exporter. Target: %s.\n", *target)

	scannerCfg := &pgscanner.Config{
		AutoConfirm: *autoConfirm,
	}

	pgScanner := pgscanner.New(scannerCfg, os.Stdin, os.Stdout, planner)
	pgScanner.Run(ctx)
}
