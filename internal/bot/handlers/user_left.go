package handlers

import (
	"fmt"
	"math/rand"
	"time"

	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) OnUserLeft(m *tgbotapi.Message) {
	leftUser := m.LeftChatMember
	chatID := m.Chat.ID
	username := c.GetUsername(leftUser)

	if leftUser.IsBot {
		return
	}

	h.logger.Infof("User %s left the chat", username)
	goodbyeMsgText := getGoodbyeMsgText(username)

	msg := tgbotapi.NewMessage(chatID, goodbyeMsgText)
	msg.ReplyToMessageID = m.MessageID

	sentMsg, err := h.bot.Send(msg)
	if err != nil {
		h.logger.Errorf("Error sending goodbye message: %s", err)
		return
	}

	go h.DeleteMessage(chatID, sentMsg.MessageID, DeleteMsgDelayThirty)
}

func getGoodbyeMsgText(username string) string {
	var messages = []string{
		"I'll miss you %s, but don't worry, I'll find someone else to make fun of 😉",
		"Aww, leaving so soon, %s? Well, it was nice while it lasted. Bye! 👋",
		"Goodbye, %s. And don't forget to send me a postcard from wherever it is you're going! 📬",
		"You're leaving? But we were just getting to the good part! Fine, bye then, %s. 👐",
		"Thanks for chatting, but it's time for you to go now. See ya, %s! 👽",
		"Bye, and don't forget to come back and see me! I'll be here waiting, %s. ✌️ ",
		"Bye, don't forget to come back and entertain me with your presence, %s! 🤣",
		"I hope you're not leaving because I laughed too much at my own jokes... Fine, bye then, %s! 🤭",
		"Leaving already? But I was just about to show you my collection of events! Bye then! %s 👀",
		"Thanks for the chat, but I have to go now. I have a meeting with my imaginary friend. Bye! %s! 😹",
		"Don't forget to grab a souvenir on your way out. How about a mug with my face on it? Bye, %s! 🧝",
	}

	rand.Seed(time.Now().UnixNano())
	randMsg := messages[rand.Intn(len(messages))]

	return fmt.Sprintf(randMsg, username)
}
