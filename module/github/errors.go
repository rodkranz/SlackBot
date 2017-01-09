// Package github contains the errors available
// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package github

import "fmt"

// ErrGitHubUnmarshal error if can't possible unmarshal gitHub input webhook
type ErrGitHubUnmarshal struct {
	Err error
}

// IsErrGitHubUnmarshal returns if the error type is ErrGitHubUnmarshal or not.
func IsErrGitHubUnmarshal(err error) bool {
	_, ok := err.(ErrGitHubUnmarshal)
	return ok
}

// Error interface for error.
func (e ErrGitHubUnmarshal) Error() string {
	return fmt.Sprintf("Error to parse github payload: %v", e.Err.Error())
}

