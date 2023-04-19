package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var verifyKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("I am a Human", VerifyUserButton),
	),
)
