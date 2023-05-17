package handlers

import (
	"fmt"

	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) OnCmdsCommand(m *tgbotapi.Message) {
	username := c.GetUsername(m.From)
	chatID := m.Chat.ID

	h.logger.Infof("Received /cmds command from %s", username)

	msgText := fmt.Sprintf("Hi %s, here is a list of commands you can use:\n\n"+
		"'/info' - Get a list of upcoming events🎉\n"+
		"'/info (followed by your town 'Machester' etc)' - Get a list of upcoming "+
		"events for your town🏡\n"+
		"'/commands' - Get a list of commands📜\n"+
		"'/rules' - Read the rules of the chat⚠️ \n\n"+ 
		"📢Please take note that the list of new commands will be added soon. "+
		"Make sure to check the list of commands more frequently.",
		username)

	msg := tgbotapi.NewMessage(chatID, msgText)
	msg.ReplyToMessageID = m.MessageID
	msg.ParseMode = "Markdown"

	sentMsg, err := h.bot.Send(msg)
	if err != nil {
		h.logger.Error("Error sending message", err)
		return
	}

	go h.DeleteMessage(chatID, sentMsg.MessageID, DeleteMsgDelayThreeMin)
}
