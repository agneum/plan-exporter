package pgscanner

// ScannerState modes.
const (
	NormalMode = iota
	ScanningMode
	ConfirmingMode
)

// ScannerState contains the current state of the PgScanner.
type ScannerState struct {
	buffer string
	mode   int
}

// SetBuffer sets a buffer value.
func (st *ScannerState) SetBuffer(buffer string) {
	st.buffer = buffer
}

// CleanBuffer removes a plan buffer content.
func (st *ScannerState) CleanBuffer() {
	st.buffer = ""
}

// SetMode changes the scanner state mode.
func (st *ScannerState) SetMode(mode int) {
	st.mode = mode
}

// IsMode checks if the scanner state is in a specific mode.
func (st *ScannerState) IsMode(mode int) bool {
	return st.mode == mode
}

// Reset resets the scanner state.
func (st *ScannerState) Reset() {
	st.mode = NormalMode
	st.CleanBuffer()
}
