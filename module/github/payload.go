// Package github is for return a payload parsed.
// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package github

import "encoding/json"

// NewPayload returns a new struct of github webhook.
func NewPayload(b []byte) (*Payload, error) {
	p := &Payload{}

	if err := json.Unmarshal(b, p); err != nil {
		return p, ErrGitHubUnmarshal{Err: err}
	}

	return p, nil
}
