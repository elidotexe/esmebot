package common

import (
	"github.com/elidotexe/esme/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// IsAdmin is a function that takes an update, a bot API instance, and a logger,
// and returns a boolean value indicating whether the user associated with
// the update is an admin.
func IsAdmin(m *tgbotapi.Message, b *tgbotapi.BotAPI, l *logger.Logger) bool {
	admins, _ := getChatAdmins(m, b, l)

	userIsAdmin := false
	for _, admin := range admins {
		if m.From.ID == admin.User.ID {
			userIsAdmin = true
			break
		}
	}

	return userIsAdmin
}

// getChatAdmins is a function that takes a message, a bot API instance,
// and a logger, and returns a list of chat administrators and an error (if any).
func getChatAdmins(
	m *tgbotapi.Message,
	b *tgbotapi.BotAPI,
	l *logger.Logger) ([]tgbotapi.ChatMember, error) {
	adminConfig := tgbotapi.ChatAdministratorsConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID:             m.Chat.ID,
			SuperGroupUsername: m.Chat.ChatConfig().SuperGroupUsername,
		},
	}

	admins, err := b.GetChatAdministrators(adminConfig)
	if err != nil {
		l.Errorf("Error getting chat admins: %v", err)
		return nil, err
	}

	return admins, nil
}
