// Package conf is configuration of webHook
// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package conf

var (
	// HTTPHost configuration of address
	HTTPHost = "0.0.0.0"

	// HTTPPort configuration of port
	HTTPPort = 8080

	// SecretWord is text for check when receive webHook
	SecretWord = "LetMeInBananaGames"

	// SlackAPI incoming webHooks
	SlackAPI = "https://hooks.slack.com/services/?????????????"

	// DateFormat data of bot.
	DateFormat = "01/02/2006 15:04"
)
