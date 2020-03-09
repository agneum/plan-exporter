package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const QueryPlanSubstringDetector = "QUERY PLAN"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	explainLines := []string{}
	isScanningMode := false

	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, QueryPlanSubstringDetector) {
			isScanningMode = true
		}

		if !isScanningMode {
			continue
		}

		if text == "" {
			isScanningMode = false

			if len(explainLines) < 2 {
				log.Printf("Not enough lines in plan to post.")
				continue
			}

			plan := strings.Join(explainLines[2:len(explainLines)-1], "\n")

			explainLines = []string{}
			url, err := postToDalibo(plan)
			if err != nil {
				log.Printf("Failed to post query plan: %s.", err)
				continue
			}

			fmt.Printf("The plan has been posted successfully.\nURL: %s", url)
			continue
		}

		explainLines = append(explainLines, text)
	}

	fmt.Println("Welcome to the query sender.")
}
