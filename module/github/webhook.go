// Package github contain payload struct
package github

import "time"

// Payload struct of slack-bot
type Payload struct {
	Ref        string     `json:"ref"`
	Compare    string     `json:"compare"`
	Commits    []*Commit   `json:"commits"`
	HeadCommit *Commit     `json:"head_commit"`
	Repository *Repository `json:"repository"`
	Pusher     *Pusher     `json:"pusher"`
	Sender     *Sender     `json:"sender"`
}

// Commit struct of commit information
type Commit struct {
	ID        string    `json:"id"`
	TreeID    string    `json:"tree_id"`
	URL       string    `json:"url"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Author    *Pusher   `json:"author"`
	Committer *Pusher   `json:"committer"`
	Added     []string  `json:"added"`
	Modified  []string  `json:"modified"`
	Removed   []string  `json:"removed"`
}

// Repository struct with repository information
type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Owner       *Pusher `json:"owner"`
}

// Sender struct with information of sender
type Sender struct {
	AvatarURL string `json:"avatar_url"`
	Login     string `json:"login"`
	HTMLURL   string `json:"html_url"`
}

// Pusher struct with information about user
type Pusher struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// GetName retries the name of Pusher.
func (p *Pusher) GetName() string {
	if len(p.Name) > 0 {
		return p.Name
	}
	if len(p.Username) > 0 {
		return p.Username
	}
	if len(p.Email) > 0 {
		return p.Email
	}
	return "Unknown"
}
