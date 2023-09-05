package validator

import (
	"fmt"
	"strings"
)

func NewValidationError(expected string, actual string, possibleCause ...string) ValidationError {
	cause := "Unknown"
	if len(possibleCause) != 0 {
		cause = strings.Join(possibleCause, ",")
	}

	return ValidationError{
		Expected:      expected,
		Actual:        actual,
		PossibleCause: cause,
	}
}

type ValidationError struct {
	Expected      string
	Actual        string
	PossibleCause string
}

func (receiver ValidationError) Error() string {
	return fmt.Sprintf(
		"Expected to recieve: %s\nBut Received: %s\nPossibleCause: %s", //
		receiver.Expected, receiver.Actual, receiver.PossibleCause)
}
