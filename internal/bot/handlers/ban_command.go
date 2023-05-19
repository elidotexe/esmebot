package handlers

import (
	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) OnBanUserCommand(m *tgbotapi.Message) {
	userIsAdmin := c.IsAdmin(m, h.bot, h.logger)
	if !userIsAdmin {
		return
	}

	h.logger.Infof("Received ban user command for user %s", m.CommandArguments())
	userToBan := m.CommandArguments()

	chatMemberConfig := tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID:             m.Chat.ID,
			SuperGroupUsername: userToBan,
		},
	}

	_, err := h.bot.GetChatMember(chatMemberConfig)
	if err != nil {
		h.logger.Errorf("Error getting chat info: %s", err)
		return
	}
}
