package handlers

import (
	"fmt"

	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const VerifyButton = "verify"

var (
	newUserChatID int64
	newUserID     int64
	isUserHuman   bool = false
)

func (h *Handlers) HandleNewUser(m *tgbotapi.Message, query *tgbotapi.CallbackQuery) {
	switch {
	case m != nil:
		h.onUserJoined(m)
	case query != nil:
		h.handleButton(query)
	}
}

func (h *Handlers) onUserJoined(m *tgbotapi.Message) {
	userIsAdmin := c.IsAdmin(m, h.bot, h.logger)

	for _, user := range m.NewChatMembers {
		username := c.GetUsername(&user)
		newUserChatID = m.Chat.ID
		newUserID = user.ID

		h.logger.Infof("New user joined: %v", username)
		if user.IsBot {
			isBotAllowed, err := h.botIsAllowedToJoin(m, user, userIsAdmin, username, newUserChatID)
			if err != nil {
				h.logger.Errorf("Error checking if user is bot: %v", err)
				return
			}

			if isBotAllowed || !isBotAllowed {
				return
			}
		}

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("I am a Human", VerifyButton),
			),
		)

		verifyUserMsgText := h.getVerifyUserMsgText(username, m.Chat.Title)

		msg := tgbotapi.NewMessage(newUserChatID, verifyUserMsgText)
		msg.ReplyMarkup = keyboard

		_, err := h.bot.Send(msg)
		if err != nil {
			h.logger.Errorf("Error sending message: %v", err)
			return
		}
	}

	if m.From.ID == newUserID && !isUserHuman {
		delMsg := tgbotapi.NewDeleteMessage(newUserChatID, m.MessageID)
		_, err := h.bot.Request(delMsg)
		if err != nil {
			h.logger.Errorf("Error deleting message: %v", err)
			return
		}
	}
}

func (h *Handlers) handleButton(query *tgbotapi.CallbackQuery) {
	if query.Data == VerifyButton && query.From.ID == newUserID {
		delMsg := tgbotapi.NewDeleteMessage(newUserChatID, query.Message.MessageID)

		_, err := h.bot.Request(delMsg)
		if err != nil {
			h.logger.Errorf("Error deleting message: %v", err)
			return
		}

		isUserHuman = true
	}
}

func (h *Handlers) getVerifyUserMsgText(username, chatTitle string) string {
	verifyUserMsg := fmt.Sprintf("Hi %s, welcome to the %s! (You have %sec)\n"+
		"\n"+
		"Please, press the button below within the specified time frame, otherwise you "+
		"will be kicked. Thank you!", username, chatTitle, DeleteMessageTime.String())

	return verifyUserMsg
}

func (h *Handlers) botIsAllowedToJoin(
	m *tgbotapi.Message,
	user tgbotapi.User,
	userIsAdmin bool,
	username string,
	chatID int64,
) (bool, error) {
	if !userIsAdmin {
		h.logger.Infof("Banning bot: %v", username)
		banChatMemberConfig := tgbotapi.BanChatMemberConfig{
			ChatMemberConfig: tgbotapi.ChatMemberConfig{
				ChatID: chatID,
				UserID: user.ID,
			},
			RevokeMessages: false,
		}

		h.bot.Request(banChatMemberConfig)

		msg := tgbotapi.NewMessage(chatID, "Only admins can add bots to this group.")
		sentMsg, err := h.bot.Send(msg)
		if err != nil {
			h.logger.Errorf("Error sending message: %v", err)
			return false, err
		}

		go h.DeleteMessage(chatID, sentMsg.MessageID)

		return false, nil
	}

	return true, nil
}
