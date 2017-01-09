// Package slack contain send actions
package slack

import (
	"net/url"
	"net/http"

	"github.com/rodkranz/slack-bot/module/conf"
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
