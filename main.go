package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/agneum/plan-exporter/pgscanner"
	"github.com/agneum/plan-exporter/visualizer"
)

func main() {
	target := flag.String("target", "depesz", "type of an explain visualizer to export")

	flag.Parse()

	ctx := context.Background()
	planner, err := visualizer.New(*target)

	if err != nil {
		log.Fatalf("failed to init a query plan exporter: %v", err)
	}

	fmt.Printf("Welcome to the query plan exporter. Target: %s.\n", *target)

	pgsc := pgscanner.New(os.Stdin, os.Stdout, planner)
	pgsc.Run(ctx)
}
