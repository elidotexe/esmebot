package handlers

import (
	"fmt"
	"time"

	c "github.com/elidotexe/esme/internal/bot/common"
	k "github.com/elidotexe/esme/internal/bot/keyboard"
	"github.com/elidotexe/esme/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var isUserHuman bool

func (h *Handlers) OnUserJoined(m *tgbotapi.Message) {
	for _, user := range m.NewChatMembers {
		isUserHuman = false

		chatID := m.Chat.ID
		userID := user.ID

		userIsAdmin := c.IsAdmin(m, h.bot, h.logger)
		username := c.GetUsername(&user)

		h.logger.Infof("New user joined: %v", username)
		if user.IsBot {
			isBotAllowed, err := h.botIsAllowedToJoin(
				m,
				user,
				userIsAdmin,
				username,
				userID)
			if err != nil {
				h.logger.Errorf("Error checking if user is bot: %v", err)
				return
			}

			if isBotAllowed || !isBotAllowed {
				return
			}
		}

		captchaMsgText := getCaptchaMsgText(username, m.Chat.Title)

		msg := tgbotapi.NewMessage(chatID, captchaMsgText)
		msg.ReplyMarkup = k.CreateInlineKeyboard("I am a Human", VerifyUserButton)

		sentMsg, err := h.bot.Send(msg)
		if err != nil {
			h.logger.Errorf("Error sending message: %v", err)
			return
		}

		h.storage.Add(chatID, userID, storage.Info{
			CaptchaMessage: sentMsg,
			IsHuman:        isUserHuman,
		})

		go h.deleteCaptchaMessage(chatID, userID, sentMsg.MessageID)
	}
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

		go h.DeleteMessage(chatID, sentMsg.MessageID, DeleteMsgDelayFiveMin)

		return false, nil
	}

	return true, nil
}

func getCaptchaMsgText(username, chatTitle string) string {
	return fmt.Sprintf("Hi %s, welcome to the %s! (You have %s)\n"+
		"\n"+
		"Please, press the button below within the specified time frame, otherwise you "+
		"will be kicked. Thank you!", username, chatTitle, DeleteMsgDelayFiveMin.String())
}

func (h *Handlers) deleteCaptchaMessage(chatID int64, userID int64, msgID int) {
	time.AfterFunc(DeleteMsgDelayFiveMin, func() {
		if !isUserHuman {
			h.DeleteMessage(chatID, msgID, DeleteMsgDelayZeroMin)

			banChatMemberCfg := tgbotapi.BanChatMemberConfig{
				ChatMemberConfig: tgbotapi.ChatMemberConfig{
					ChatID: chatID,
					UserID: userID,
				},
				RevokeMessages: true,
			}

			h.bot.Request(banChatMemberCfg)

			h.storage.Remove(chatID, userID)
		}
	})
}
