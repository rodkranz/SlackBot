// Package hydrate is for fill slackPayload of github infromation
// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package hydrate

import (
	"github.com/rodkranz/webhook/module/slack"
	"github.com/rodkranz/webhook/module/github"
)

// NewSlackFromGitHub returns slack.Payload with information about github.
func NewSlackFromGitHub(gp *github.Payload) *slack.Payload {
	sp          := &slack.Payload{}
	sp.Username = "GoBot"
	sp.Channel  = "#general"
	sp.Emotion  = "ghost"

	a            := &slack.Attachment{}
	a.AuthorName = gp.Pusher.GetName()
	a.AuthorIcon = gp.Sender.AvatarURL
	a.AuthorLink = gp.Sender.HTMLURL
	a.Title      = gp.Repository.Name
	a.TitleLink  = gp.Repository.URL
	a.Text       = gp.HeadCommit.Message
	a.Color      = "#36a64f"
	a.Footer     = "by Gobot"
	a.FooterIcon = "https://golang.org/favicon.ico"
	sp.AppendAttachment(a)

	return sp
}
