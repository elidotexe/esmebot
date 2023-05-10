package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) MessageHandler(m *tgbotapi.Message) {
	chatID := m.Chat.ID
	userID := m.From.ID

	_, ok := h.storage.Exist(chatID, userID)
	if !ok {
		return
	}

	if !isUserHuman {
		h.DeleteMessage(chatID, m.MessageID, DeleteMsgDelayZeroMin)
		return
	}

	h.storage.Remove(chatID, userID)

	go h.deleteMessages(m)
}

func (h *Handlers) deleteMessages(m *tgbotapi.Message) {
	// ..
}
