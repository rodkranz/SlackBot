// Package conf is configuration of webHook
package conf

var (
	// HTTPHost configuration of address
	HTTPHost = "0.0.0.0"

	// HTTPPort configuration of port
	HTTPPort = 8080

	// SecretWord is text for check when receive webHook
	SecretWord = "LetMeInBananaGames"

	// SlackAPI incoming webHooks
	SlackAPI = "https://hooks.slack.com/services/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"

	// DateFormat data of bot.
	DateFormat = "01/02/2006 15:04"
)
