package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const daliboURL = "https://explain.dalibo.com"

func postToDalibo(plan string) (string, error) {
	fmt.Println("Posting to the Dalibo visualizer...")

	response, err := makeRequest(plan)
	if err != nil {
		return "", errors.Wrap(err, "failed to make a request")
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	// Load the HTML document.
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to load the HTML document")
	}

	// Find the plan ID.
	pageTitle := doc.Find("title").Text()
	pageTitle = strings.TrimLeft(strings.Join(strings.Fields(pageTitle), ""), "Plan")
	titleParts := strings.SplitN(pageTitle, "|", 2)

	if len(titleParts) < 2 {
		return "", errors.New("failed to parse a plan ID")
	}

	return fmt.Sprintf("%s/plan/%s", daliboURL, titleParts[0]), nil
}

func makeRequest(plan string) (*http.Response, error) {
	formVal := url.Values{}
	formVal.Add("plan", plan)

	response, err := http.PostForm(fmt.Sprintf("%s/new", daliboURL), formVal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to post form")
	}

	return response, nil
}
