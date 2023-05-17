package handlers

import (
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const ChatID = -1001628672322 // Test chat

func (h *Handlers) NotificationMessage() {
	go func() {
		for {
			msg := tgbotapi.NewMessage(ChatID, getRandomMessage())
			msg.DisableWebPagePreview = true
			msg.ParseMode = "markdown"

			_, err := h.bot.Send(msg)
			if err != nil {
				h.logger.Errorf("Failed to send random message: %v", err)
				return
			}

			time.Sleep(time.Hour * 48)
		}
	}()
}

func getRandomMessage() string {
	var messages = []string{
		"✨Just a friendly reminder that you can use the following commands in the chat: \n" +
			"'/info' - Get a list of upcoming events🎉\n" +
			"'/info (followed by your town 'Machester' etc)' - Get a list of upcoming " +
			"events for your town🏡\n" +
			"'/commands' - Get a list of commands📜\n" +
			"'/rules' - Read the rules of the chat⚠️ \n\n"+  
			"📢Please take note that the list of new commands will be added soon. " +
			"Make sure to check the list of commands more frequently.",

		"⚠️If you're a psytrance event organiser and you wish to promote your event and " +
			"get additional traffic to your page, you can easily achieve this by adding " +
			"your event to the https://goabase.net\n\n" +
			"We are pulling the data from their API and all events are automatically added " +
			"here in the chat. That's because we are wizards..🧙👽",

		"📢*Feel free to share information about music/art events in this group without* " +
			"*requiring prior permission*.\n\n" +
			"We encourage members to contribute to the community " +
			"by posting about upcoming events, performances, or any other " +
			"artistic events they find noteworthy. Sharing such information helps create a " +
			"vibrant and engaging environment that celebrates and supports the arts.👽✌️",

		"🎉*Coming soon*\n" +
			"Stay tuned as I'm currently exploring Skiddle's API, and more events will be " +
			"added soon to the existing event list.",

		"🤓I have nothing better to do with my time, but code...\n" +
			"I can buld you a custom bot for your group and make it do whatever you want.\n\n" +
			"Hit me up if you're interested: @elicodesbot",
	}

	rand.Seed(time.Now().UnixNano())
	randMsg := messages[rand.Intn(len(messages))]

	return randMsg
}
