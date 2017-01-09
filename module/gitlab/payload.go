// Package gitlab is for return a payload parsed.
package gitlab

import "encoding/json"

// NewPayload returns a new struct of github slack-bot.
func NewPayload(b []byte) (*Payload, error) {
	p := &Payload{}

	if err := json.Unmarshal(b, p); err != nil {
		return p, ErrGitLabUnmarshal{Err: err}
	}

	return p, nil
}
