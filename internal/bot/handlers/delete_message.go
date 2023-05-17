package handlers

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// DeleteMsgDelayFiveMin is the time to wait before deleting a message
const DeleteMsgDelayZeroMin = time.Second * 0
const DeleteMsgDelayThirty = time.Second * 30
const DeleteMsgDelayThreeMin = time.Minute * 3
const DeleteMsgDelayFiveMin = time.Minute * 5

// DeleteMessage deletes a message by sending a request to the Telegram API
// with a DeleteMessage command. It sleeps for DeleteMessageTime before making
// the request to allow for a delay in message delivery. Returns an error if
// the request fails.
func (h *Handlers) DeleteMessage(
	chatID int64,
	messageID int,
	deleteMsgDelay time.Duration) error {
	time.Sleep(deleteMsgDelay)

	delMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := h.bot.Request(delMsg)
	if err != nil {
		h.logger.Errorf("Error deleting message: %v", err)
		return err
	}

	return nil
}
