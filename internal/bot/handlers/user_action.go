package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) DeleteMessageFromUnverifiedUser(m *tgbotapi.Message) {
	if m.From.ID == newUserID && !isUserHuman {
		h.DeleteMessage(
			newUserChatID,
			m.MessageID,
			DeleteMsgDelayZeroMin)
	}
}

func (h *Handlers) VerifyButtonQueryHandler(query *tgbotapi.CallbackQuery) {
	if query.From.ID == newUserID {
		h.DeleteMessage(
			newUserChatID,
			query.Message.MessageID,
			DeleteMsgDelayZeroMin)

		isUserHuman = true
	}
}
