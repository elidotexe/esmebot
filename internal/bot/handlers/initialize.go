package handlers

import (
	"time"

	"github.com/elidotexe/esme/internal/logger"
	"github.com/elidotexe/esme/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const DeleteMessageTime = time.Second * 30

type Handlers struct {
	bot     *tgbotapi.BotAPI
	logger  *logger.Logger
	storage *storage.MemoryStorage
}

// Initialize initializes a new instance of the Handlers struct with the provided BotAPI
// and Logger, and returns it along with a nil error.
func Initialize(b *tgbotapi.BotAPI, logger *logger.Logger, storage *storage.MemoryStorage) (*Handlers, error) {
	return &Handlers{
		bot:     b,
		logger:  logger,
		storage: storage,
	}, nil
}

func (h *Handlers) HandleUpdate(u tgbotapi.Update) {
	switch {
	case u.Message.NewChatMembers != nil:
		// h.OnUserJoined(u.Message)
	case u.CallbackQuery.Data != "":
		h.HandleButton(u.CallbackQuery)
	}
}

// Helper functions

// DeleteMessage deletes a message by sending a request to the Telegram API
// with a DeleteMessage command. It sleeps for DeleteMessageTime before making
// the request to allow for a delay in message delivery. Returns an error if
// the request fails.
func (h *Handlers) DeleteMessage(chatID int64, messageID int) error {
	time.Sleep(DeleteMessageTime)

	delMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := h.bot.Request(delMsg)
	if err != nil {
		h.logger.Errorf("Error deleting message: %v", err)
		return err
	}

	return nil
}
