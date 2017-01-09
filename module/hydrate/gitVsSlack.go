// Package hydrate is for fill slackPayload of github information
package hydrate

import (
	"github.com/rodkranz/slack-bot/module/slack"
	"github.com/rodkranz/slack-bot/module/github"
	"github.com/rodkranz/slack-bot/module/gitlab"
)

// NewSlackFromGitHub returns slack.Payload with information about github.
func NewSlackFromGitHub(gp *github.Payload) *slack.Payload {
	sp := &slack.Payload{}
	sp.Username = "GoBot"
	sp.Channel = "#gobot"
	sp.Emotion = "ghost"

	a := &slack.Attachment{}
	a.AuthorName = gp.Pusher.GetName()
	a.AuthorIcon = gp.Sender.AvatarURL
	a.AuthorLink = gp.Sender.HTMLURL
	a.Title = gp.Repository.Name
	a.TitleLink = gp.Repository.URL
	a.Text = gp.HeadCommit.Message
	a.Color = "#36a64f"
	a.Footer = "by Gobot"
	a.FooterIcon = "https://golang.org/favicon.ico"
	sp.AppendAttachment(a)

	return sp
}

// NewSlackFromGitLab returns slack.Payload with information about github.
func NewSlackFromGitLab(gp *gitlab.Payload) *slack.Payload {
	sp := &slack.Payload{}
	sp.Username = "GoBot"
	sp.Channel = "#gobot"
	sp.Emotion = "ghost"
	sp.Text = "New push to master! " + gp.After[0:7]

	for _, commit := range gp.Commits {
		a := &slack.Attachment{}
		a.AuthorName = commit.Author.GetName()
		a.TitleLink = gp.Project.WebURL + "+commits/" + commit.ID
		a.Title = commit.Message

		a.Color = "#36a64f"
		a.Footer = "by Gobot"
		a.FooterIcon = "https://golang.org/favicon.ico"

		sp.AppendAttachment(a)
	}

	return sp
}
