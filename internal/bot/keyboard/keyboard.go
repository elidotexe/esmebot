package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// CreateInlineKeyboard creates an inline keyboard with a single button
// with the provided text and callback data. Returns the created keyboard.
// Currently being used when new users join a group.
func CreateInlineKeyboard(buttonText, callbackData string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackData),
		),
	)
}

func CreateMenuMarkup(text, query string) *tgbotapi.InlineKeyboardButton {
	return &tgbotapi.InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: &query,
	}
}
