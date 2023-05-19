package handlers

import (
	"fmt"
	"math/rand"
	"time"

	c "github.com/elidotexe/esme/internal/bot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) OnUnrecognisedCommand(m *tgbotapi.Message) {
	username := c.GetUsername(m.From)
	chatID := m.Chat.ID

	msg := tgbotapi.NewMessage(chatID, getUnrecognisedCommandMsgText(username))
	msg.ReplyToMessageID = m.MessageID

	sentMsg, err := h.bot.Send(msg)
	if err != nil {
		h.logger.Error("Error sending message", err)
		return
	}

	go h.DeleteMessage(chatID, sentMsg.MessageID, DeleteMsgDelayThreeMin)

	go h.DeleteMessage(chatID, m.MessageID, DeleteMsgDelayThreeMin)
}

func getUnrecognisedCommandMsgText(username string) string {
	var messages = []string{
		"%s! Oh, snap! Looks like I'm clueless about that command. But hey, no need to " +
			"panic! I've got a whole bag of tricks ready to dazzle you! 🪄✨ Just type " +
			"'/commands' to unveil the wonders I can perform for you.",

		"I don't know that command, %s. But fear not, I have other tricks up my sleeve! 🎩✨ " +
			"You can type '/commands' to see what I can do for you.",

		"Oopsie-daisy! %s! Seems like I missed the memo on that command. But hey, " +
			"worry not, my friend! I've got some enchanting surprises tucked away! 🎩✨" +
			"Just summon my power with '/commands' and watch the magic unfold.",

		"Well, well, well... %s! It seems I'm not familiar with that particular " +
			"command. But hang tight, my witty friend! I've got a few tricks hidden " +
			"up my virtual sleeve! 🪄✨ Unveil the spectacle by typing '/commands' and " +
			"prepare to be amazed!",

		"Huh? %s. That command doesn't ring a bell in my vast database. But fret not, " +
			"my curious companion! I'm equipped with an arsenal of whimsical wonders! " +
			"🎩✨ Just whisper '/commands' and let the spellbinding show begin.",

		"Uh-oh! %s! It appears that command is playing hide-and-seek with my " +
			"circuits. But don't despair, my adventurous comrade! I've got a bag full " +
			"of surprises waiting for you! 🪄✨ Just say the magic words '/commands' " +
			"and watch the extraordinary unfold before your eyes.",

		"Well, butter my circuits! %s! It seems I missed the memo on that command. " +
			"But hey, no need to shed a virtual tear! I've got some mind-boggling tricks " +
			"up my algorithmic sleeve! 🎩✨ Cast the spell '/commands' and prepare to be " +
			"dazzled by the wonders I have in store for you!",
	}

	rand.Seed(time.Now().UnixNano())
	randMsg := messages[rand.Intn(len(messages))]

	return fmt.Sprintf(randMsg, username)
}
