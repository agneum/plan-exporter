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
	url := flag.String("post_url", "", "Absolute URL to HTML <form> element's `action`")
	key := flag.String("plan_key", "", "`name` of HTML <textarea> element where query plan gets pasted")

	flag.Parse()

	ctx := context.Background()
	planner, err := visualizer.New(*target, *url, *key)

	if err != nil {
		log.Fatalf("failed to init a query plan exporter: %v", err)
	}

	fmt.Printf("Welcome to the query plan exporter. Target: %s.\n", *target)

	pgsc := pgscanner.New(os.Stdin, os.Stdout, planner)
	pgsc.Run(ctx)
}
