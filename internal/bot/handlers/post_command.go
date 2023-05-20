package handlers

import (
	"strings"

	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) OnPostCommand(m *tgbotapi.Message) {
	username := c.GetUsername(m.From)
	userIsAdmin := c.IsAdmin(m, h.bot, h.logger)

	if !userIsAdmin && m.From.ID != 1087968824 { // annonimous group user ID
		h.logger.Infof("User %s is not an admin", username)
		return
	}

	h.logger.Infof("Received /post command from %s", username)

	args := strings.SplitN(m.Text, " ", 2)
	if len(args) < 2 {
		h.logger.Errorf("You must provide a message to post")
		return
	}

	postMsg := args[1]
	msg := tgbotapi.NewMessage(RavenexusID, postMsg)

	_, err := h.bot.Send(msg)
	if err != nil {
		h.logger.Error("Error sending message", err)
		return
	}

	go h.DeleteMessage(RavenexusID, m.MessageID, DeleteMsgDelayZeroMin)
}
