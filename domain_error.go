package timer

// domainError is a custom error type
type domainError struct {
	Err string
}

// Error returns the error message
func (e *domainError) Error() string {
	return e.Err
}

// ErrInvalidDuration is returned when the duration is longer than 6 hours
var ErrInvalidDuration = &domainError{Err: "duration exceeds 6 hours"}
