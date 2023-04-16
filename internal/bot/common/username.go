package common

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// GetUsername returns the username of a Telegram user. If the user has a username,
// it is returned with a "@" prefix. If not, the user's first and last name are
// concatenated and returned as the username.
func GetUsername(m *tgbotapi.User) string {
	if m.UserName != "" {
		return "@" + m.UserName
	}

	username := ""
	username = m.FirstName
	if m.LastName != "" {
		username = username + " " + m.LastName
	}

	return username
}
