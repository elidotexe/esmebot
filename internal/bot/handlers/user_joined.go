package handlers

import (
	"fmt"
	"time"

	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) OnUserJoined(m *tgbotapi.Message) {
	for _, user := range m.NewChatMembers {
		userIsAdmin := c.IsAdmin(m, h.bot, h.logger)
		username := c.GetUsername(&user)

		h.logger.Infof("New user joined: %v", username)
		newUserChatID = m.Chat.ID
		newUserID = user.ID

		if user.IsBot {
			isBotAllowed, err := h.botIsAllowedToJoin(
				m,
				user,
				userIsAdmin,
				username,
				newUserChatID)
			if err != nil {
				h.logger.Errorf("Error checking if user is bot: %v", err)
				return
			}

			if isBotAllowed || !isBotAllowed {
				return
			}
		}

		verifyUserMsgText := getVerifyUserMsgText(username, m.Chat.Title)

		msg := tgbotapi.NewMessage(newUserChatID, verifyUserMsgText)
		msg.ReplyMarkup = verifyKeyboard

		sentMsg, err := h.bot.Send(msg)
		if err != nil {
			h.logger.Errorf("Error sending message: %v", err)
			return
		}

		time.AfterFunc(DeleteMsgDelayFiveMin, func() {
			if !isUserHuman {
				h.DeleteMessage(
					newUserChatID,
					sentMsg.MessageID,
					DeleteMsgDelayZeroMin)

				banChatMemberCfg := tgbotapi.BanChatMemberConfig{
					ChatMemberConfig: tgbotapi.ChatMemberConfig{
						ChatID: newUserChatID,
						UserID: newUserID,
					},
					RevokeMessages: true,
				}

				h.bot.Request(banChatMemberCfg)
			}
		})
	}

	go h.DeleteMessageFromUnverifiedUser(m)
}

func getVerifyUserMsgText(username, chatTitle string) string {
	return fmt.Sprintf("Hi %s, welcome to the %s! (You have %sec)\n"+
		"\n"+
		"Please, press the button below within the specified time frame, otherwise you "+
		"will be kicked. Thank you!", username, chatTitle, DeleteMsgDelayFiveMin.String())
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

		go h.DeleteMessage(
			chatID,
			sentMsg.MessageID,
			DeleteMsgDelayFiveMin)

		return false, nil
	}

	return true, nil
}
