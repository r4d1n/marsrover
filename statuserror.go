package marsrover

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Msg  string
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Msg
}

// Returns HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}
