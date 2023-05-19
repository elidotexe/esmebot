package handlers

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const WaitToReplyTime = time.Second * 4

func (h *Handlers) OnPrivateMessage(m *tgbotapi.Message) {
	chatID := m.Chat.ID

	h.logger.Infof("Received private message from %s", m.From.UserName)
	var msgText string

	msgText = "Greetings, my purpose in this digital realm is to cater to your every " +
		"need within the ravenuexus group chat!"
	msg := tgbotapi.NewMessage(chatID, msgText)
	msg.ReplyToMessageID = m.MessageID

	h.sendPrivateMessage(msg)

	time.Sleep(WaitToReplyTime)
	msgText = "If thou dost possess queries, seeketh answers from thy earthly " +
		"comrade, who hath a companion acquainted with another acquaintance linked to " +
		"mine own celestial creator.👽✌️"
	secMsg := tgbotapi.NewMessage(chatID, msgText)

	h.sendPrivateMessage(secMsg)

	time.Sleep(WaitToReplyTime)
	msgText = "Farewell, puny human! May your earthly rotations be filled with " +
		"delightfully bizarre encounters and the occasional extraterrestrial mischief!🛸😉"
	thirdMsg := tgbotapi.NewMessage(chatID, msgText)

	h.sendPrivateMessage(thirdMsg)
}

func (h *Handlers) sendPrivateMessage(msgConfig tgbotapi.MessageConfig) {
	_, err := h.bot.Send(msgConfig)
	if err != nil {
		h.logger.Errorf("Error sending message: %v", err)
		return
	}
}
