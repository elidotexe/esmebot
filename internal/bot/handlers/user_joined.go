package handlers

import (
	"fmt"

	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const VerifyButton = "verify"

func (h *Handlers) OnUserJoined(m *tgbotapi.Message) {
	userIsAdmin := c.IsAdmin(m, h.bot, h.logger)

	fmt.Println("We are here:")

	for _, user := range m.NewChatMembers {
		username := c.GetUsername(&user)
		chatID := m.Chat.ID

		h.logger.Infof("New user joined: %v", username)
		if user.IsBot {
			isBotAllowed, err := h.botIsAllowedToJoin(m, user, userIsAdmin, username, chatID)
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

		msg := tgbotapi.NewMessage(chatID, verifyUserMsgText)
		msg.ReplyMarkup = keyboard

		sentMsg, err := h.bot.Send(msg)
		if err != nil {
			h.logger.Errorf("Error sending message: %v", err)
			return
		}

		fmt.Println(sentMsg)

		// if u.CallbackQuery != nil {
		// 	go h.verifyNewUser(&u, sentMsg, username, &user)
		// }
	}
}

// func (h *Handlers) verifyNewUser(
// 	u *tgbotapi.Update,
// 	sentMsg tgbotapi.Message,
// 	username string,
// 	user *tgbotapi.User) {
// 	time.Sleep(8 * time.Second)
// 	fmt.Println(user)
//
// 	callbackConfig := tgbotapi.CallbackConfig{
// 		CallbackQueryID: u.CallbackQuery.ID,
// 		Text:            "You are verified!",
// 	}
//
// 	_, err := h.bot.Request(callbackConfig)
// 	if err != nil {
// 		h.logger.Errorf("Error sending callback config: %v", err)
// 		return
// 	}
// }

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

func (h *Handlers) HandleButton(query *tgbotapi.CallbackQuery) {
	fmt.Println("We are here:")
	if query.Data == VerifyButton {
		fmt.Println("Hello from VerifyButton")
	}
}
