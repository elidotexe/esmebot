package handlers

import (
	"fmt"

	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) OnRulesCommand(m *tgbotapi.Message) {
	username := c.GetUsername(m.From)
	chatID := m.Chat.ID
	groupName := m.Chat.Title

	h.logger.Infof("Received /rules command from %s", username)

	msgText := fmt.Sprintf(
		"Hi %s, here're the rules for the %s:\n\n"+
			"1. *No drug selling/promotion*:💊\n"+
			"Discussions related to drugs will be permitted but should focus on responsible "+
			"and educational aspects. Members can discuss topics such as harm reduction, "+
			"drug policy, or personal experiences with the intention of promoting awareness "+
			"and understanding.\n\n"+

			"2. *Be respectful*:🙏\n"+
			"Encourage members to treat each other with respect and refrain from using "+
			"offensive language, personal attacks, or discriminatory remarks. This rule "+
			"helps maintain a positive and inclusive environment.\n\n"+

			"3. *No spam or abuse of bot commands*:🤖\n"+
			"Prohibit the excessive or repetitive use of bot commands and any form of "+
			"spamming within the group. This includes but is not limited to sending "+
			"multiple repetitive messages, flooding the chat with irrelevant content, "+
			"or excessively triggering bot commands.\n\n"+

			"*⚠️Feel free to share information about music/art events in this group without* "+
			"*requiring prior permission*. We encourage members to contribute to the community "+
			"by posting about upcoming events, performances, or any other "+
			"artistic events they find noteworthy. Sharing such information helps create a "+
			"vibrant and engaging environment that celebrates and supports the arts.👽✌️",
		username, groupName)

	msg := tgbotapi.NewMessage(chatID, msgText)
	msg.ReplyToMessageID = m.MessageID
	msg.ParseMode = "Markdown"

	sentMsg, err := h.bot.Send(msg)
	if err != nil {
		h.logger.Error("Error sending message", err)
		return
	}

	go h.DeleteMessage(chatID, sentMsg.MessageID, DeleteMsgDelayFiveMin)

	go h.DeleteMessage(chatID, m.MessageID, DeleteMsgDelayFiveMin)
}
