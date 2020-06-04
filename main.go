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
	postURL := flag.String("post_url", "", "Absolute URL to HTML <form> element's `action`")

	flag.Parse()

	cfg := &config.Config{
		Target:  *target,
		PostURL: *postURL,
	}

	ctx := context.Background()
	planner, err := visualizer.New(cfg)

	if err != nil {
		log.Fatalf("failed to init a query plan exporter: %v", err)
	}

	fmt.Printf("Welcome to the query plan exporter. Target: %s.\n", *target)

	pgScanner := pgscanner.New(os.Stdin, os.Stdout, planner)
	pgScanner.Run(ctx)
}
