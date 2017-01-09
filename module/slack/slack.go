// Package slack contain send actions
// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package slack

import (
	"net/url"
	"net/http"

	"github.com/rodkranz/webhook/module/conf"
)

// Send payload to slack for publish
func Send(s *Payload) (*http.Response, error) {
	data, err := s.JSON()
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Set("payload", string(data))
	return http.PostForm(conf.SlackAPI, v)
}
