package handlers

import (
	"fmt"

	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) OnInfoCommand(m *tgbotapi.Message) {
	username := c.GetUsername(m.From)
	chatID := m.Chat.ID

	h.logger.Infof("Info command received from %v", username)
	msgText := fmt.Sprintf("Hi, %v. I'm currently undergoing maintenance to improve my features, including the 'info' command which is not functioning at the moment.\n\nWhile waiting for the update, I suggest that you take the initiative to help me and post upcoming events on your own in this group. \n\nI apologise for any inconvenience this may cause ✌️👽", username)

	msg := tgbotapi.NewMessage(chatID, msgText)
	msg.ReplyToMessageID = m.MessageID

	sentMsg, err := h.bot.Send(msg)
	if err != nil {
		h.logger.Errorf("Error sending message: %v", err)
		return
	}

	if m.Command() == "info" {
		func() {
			delMsg := tgbotapi.NewDeleteMessage(chatID, m.MessageID)
			_, err := h.bot.Request(delMsg)
			if err != nil {
				h.logger.Errorf("Error deleting message: %v", err)
				return
			}
		}()
	}

	go h.DeleteMessage(
		chatID,
		sentMsg.MessageID,
		DeleteMsgDelayThirty)
}
