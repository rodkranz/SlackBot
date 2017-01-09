// Package slack contains the errors available
package slack

import "fmt"

// ErrSlackPayload error if can't convert struct to json.
type ErrSlackPayload struct {
	Err error
}

// IsErrSlackPayload returns if the error type is ErrSlackPayload or not.
func IsErrSlackPayload(err error) bool {
	_, ok := err.(ErrSlackPayload)
	return ok
}

// Error interface for error.
func (e ErrSlackPayload) Error() string {
	return fmt.Sprintf("Error slack payload: %v", e.Err.Error())
}

