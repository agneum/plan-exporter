package pgscanner

import (
	"bytes"
	"errors"
	"testing"
)

type Mock struct {
	url string
	err error
}

func (m Mock) Export(s string) (string, error) {
	return m.url, m.err
}

func TestSuccessfulPlanPosting(t *testing.T) {
	buf := &bytes.Buffer{}

	pgScanner := New(buf, buf, &Mock{url: "testURL"})
	pgScanner.state.SetBuffer("non-empty buffer")
	pgScanner.state.SetMode(ConfirmingMode)

	pgScanner.postPlan()

	actualResult := buf.String()
	expectedResult := `Posting to the visualizer...
The plan has been posted successfully.
URL: testURL`

	if actualResult != expectedResult {
		t.Errorf("failed to post a plan. Given result: %s", actualResult)
	}

	if !pgScanner.state.IsMode(NormalMode) {
		t.Errorf("invalid initial state mode. Expected: %v,  given: %v", NormalMode, pgScanner.state.mode)
	}

	if pgScanner.state.buffer != "" {
		t.Errorf("state buffer is not empty: %v", pgScanner.state.buffer)
	}
}

func TestPlanPostingWithEmptyBuffer(t *testing.T) {
	buf := &bytes.Buffer{}

	pgScanner := New(buf, buf, &Mock{url: "testURL"})
	pgScanner.state.SetMode(ConfirmingMode)

	pgScanner.postPlan()

	actualResult := buf.String()
	expectedResult := `Posting to the visualizer...
no plan to export
`
	if actualResult != expectedResult {
		t.Errorf("failed to post a plan. Given result: %s", actualResult)
	}

	if !pgScanner.state.IsMode(NormalMode) {
		t.Errorf("invalid initial state mode. Expected: %v,  given: %v", NormalMode, pgScanner.state.mode)
	}

	if pgScanner.state.buffer != "" {
		t.Errorf("state buffer is not empty: %v", pgScanner.state.buffer)
	}
}

func TestPlanPostingWithWrongMode(t *testing.T) {
	buf := &bytes.Buffer{}

	pgScanner := New(buf, buf, &Mock{url: "testURL"})
	pgScanner.state.SetBuffer("non-empty buffer")

	pgScanner.postPlan()

	actualResult := buf.String()
	expectedResult := `Posting to the visualizer...
cannot post because scanner is in invalid mode
`
	if actualResult != expectedResult {
		t.Errorf("failed to post a plan. Given result: %s", actualResult)
	}

	if !pgScanner.state.IsMode(NormalMode) {
		t.Errorf("invalid initial state mode. Expected: %v,  given: %v", NormalMode, pgScanner.state.mode)
	}

	if pgScanner.state.buffer != "" {
		t.Errorf("state buffer is not empty: %v", pgScanner.state.buffer)
	}
}

func TestPlanPostingWithFailedExport(t *testing.T) {
	buf := &bytes.Buffer{}

	pgScanner := New(buf, buf, &Mock{url: "testURL", err: errors.New("failed export")})
	pgScanner.state.SetBuffer("non-empty buffer")
	pgScanner.state.SetMode(ConfirmingMode)

	pgScanner.postPlan()

	actualResult := buf.String()
	expectedResult := `Posting to the visualizer...
Failed to post query plan: failed export.
`
	if actualResult != expectedResult {
		t.Errorf("failed to post a plan. Given result: %s", actualResult)
	}

	if !pgScanner.state.IsMode(NormalMode) {
		t.Errorf("invalid initial state mode. Expected: %v,  given: %v", NormalMode, pgScanner.state.mode)
	}

	if pgScanner.state.buffer != "" {
		t.Errorf("state buffer is not empty: %v", pgScanner.state.buffer)
	}
}
