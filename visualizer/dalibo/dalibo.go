package dalibo

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const VisualizerType = "dalibo"

const daliboURL = "https://explain.dalibo.com"

type Dalibo struct {
}

func New() *Dalibo {
	return &Dalibo{}
}

func (d *Dalibo) Export(plan string) (string, error) {
	response, err := makeRequest(plan)
	if err != nil {
		return "", fmt.Errorf("failed to make a request: %w", err)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	// Load the HTML document.
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to load the HTML document: %w", err)
	}

	const titleCount = 2

	// Find the plan ID.
	pageTitle := doc.Find("title").Text()
	pageTitle = strings.TrimLeft(strings.Join(strings.Fields(pageTitle), ""), "Plan")
	titleParts := strings.SplitN(pageTitle, "|", titleCount)

	if len(titleParts) < titleCount {
		return "", errors.New("failed to parse a plan ID")
	}

	return fmt.Sprintf("%s/plan/%s", daliboURL, titleParts[0]), nil
}

func makeRequest(plan string) (*http.Response, error) {
	formVal := url.Values{}
	formVal.Add("plan", plan)

	response, err := http.PostForm(fmt.Sprintf("%s/new", daliboURL), formVal)
	if err != nil {
		return nil, fmt.Errorf("failed to post form: %w", err)
	}

	return response, nil
}
