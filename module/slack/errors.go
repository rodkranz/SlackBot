// Package slack contains the errors available
// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
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

