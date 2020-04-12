package pgscanner

import (
	"testing"
)

func TestScannerStateMode(t *testing.T) {
	scannerState := ScannerState{}

	if scannerState.mode != NormalMode {
		t.Errorf("invalid initial state mode. Expected: %v,  given: %v", NormalMode, scannerState.mode)
	}

	scannerState.SetMode(ScanningMode)
	if scannerState.mode != ScanningMode {
		t.Errorf("invalid state mode. Expected: %v,  given: %v", ScanningMode, scannerState.mode)
	}

	if !scannerState.IsMode(ScanningMode) {
		t.Errorf("invalid state mode. Expected: %v,  given: %v", ScanningMode, scannerState.mode)
	}

	scannerState.Reset()

	if scannerState.mode != NormalMode {
		t.Errorf("invalid initial state mode. Expected: %v,  given: %v", NormalMode, scannerState.mode)
	}
}

func TestScannerStateBuffer(t *testing.T) {
	scannerState := ScannerState{}

	if scannerState.buffer != "" {
		t.Errorf("state buffer is not empty: %v", scannerState.buffer)
	}

	scannerState.SetBuffer("test buffer")
	if scannerState.buffer != "test buffer" {
		t.Errorf("invalid state buffer. Expected: %v,  given: %v", "test buffer", scannerState.buffer)
	}

	scannerState.CleanBuffer()
	if scannerState.buffer != "" {
		t.Errorf("state buffer is not empty: %v", scannerState.buffer)
	}

	scannerState.SetBuffer("test buffer")
	scannerState.Reset()

	if scannerState.buffer != "" {
		t.Errorf("state buffer is not empty: %v", scannerState.buffer)
	}
}
