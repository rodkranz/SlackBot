// Package slack contain payload
// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package slack

import "encoding/json"

// Payload struct
type Payload struct {
	Channel     string       `json:"channel,omitempty"`
	Username    string       `json:"username,omitempty"`
	Text        string       `json:"text"`
	Emotion     string       `json:"icon_emoji,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// AppendAttachment add new Attachment at Payload.
func (p *Payload) AppendAttachment(a *Attachment) {
	p.Attachments = append(p.Attachments, a)
}

// JSON returns array of bytes with json structure and error
func (p *Payload) JSON() (b []byte, err error) {
	if b, err = json.Marshal(p); err != nil {
		return b, ErrSlackPayload{Err: err}
	}

	return b, nil
}
