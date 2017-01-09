// Package github is for return a payload parsed.
package github

import "encoding/json"

// NewPayload returns a new struct of github slack-bot.
func NewPayload(b []byte) (*Payload, error) {
	p := &Payload{}

	if err := json.Unmarshal(b, p); err != nil {
		return p, ErrGitHubUnmarshal{Err: err}
	}

	return p, nil
}
