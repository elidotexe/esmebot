package handlers

import (
	"time"

	"github.com/elidotexe/esme/internal/logger"
	"github.com/elidotexe/esme/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Button constants
const VerifyUserButton = "verify"

// DeleteMsgDelayFiveMin is the time to wait before deleting a message
const DeleteMsgDelayZeroMin = time.Second * 0
const DeleteMsgDelayFiveMin = time.Second * 30

var (
	newUserChatID int64
	newUserID     int64
	isUserHuman   bool = false
)

type Handlers struct {
	bot     *tgbotapi.BotAPI
	logger  *logger.Logger
	storage *storage.Storage
}

// Initialize initializes a new instance of the Handlers struct with the provided BotAPI
// and Logger, and returns it along with a nil error.
func Initialize(b *tgbotapi.BotAPI, logger *logger.Logger, s *storage.Storage) (*Handlers, error) {
	return &Handlers{
		bot:     b,
		logger:  logger,
		storage: s,
	}, nil
}

func (h *Handlers) HandleButtonQuery(m *tgbotapi.Message, query *tgbotapi.CallbackQuery) {
	switch {
	case m != nil:
		h.OnUserJoined(m)
	case query.Data == VerifyUserButton:
		h.VerifyButtonQueryHandler(query)
	}
}

// DeleteMessage deletes a message by sending a request to the Telegram API
// with a DeleteMessage command. It sleeps for DeleteMessageTime before making
// the request to allow for a delay in message delivery. Returns an error if
// the request fails.
func (h *Handlers) DeleteMessage(
	chatID int64,
	messageID int,
	deleteMsgDelay time.Duration) error {
	time.Sleep(deleteMsgDelay)

	delMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := h.bot.Request(delMsg)
	if err != nil {
		h.logger.Errorf("Error deleting message: %v", err)
		return err
	}

	return nil
}
