package handlers

import (

	"github.com/elidotexe/esme/internal/logger"
	"github.com/elidotexe/esme/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Button constants
const VerifyUserButton = "verify"

type Handlers struct {
	bot     *tgbotapi.BotAPI
	logger  *logger.Logger
	storage storage.Storage
}

// Initialize initializes a new instance of the Handlers struct with the provided BotAPI
// and Logger, and returns it along with a nil error.
func Initialize(b *tgbotapi.BotAPI, logger *logger.Logger, s storage.Storage) (*Handlers, error) {
	return &Handlers{
		bot:     b,
		logger:  logger,
		storage: s,
	}, nil
}

func (h *Handlers) ButtonQueryHandler(query *tgbotapi.CallbackQuery) {
	if query.Data == VerifyUserButton {
		h.VerifyButtonQueryHandler(query)
	}
}

func (h *Handlers) VerifyButtonQueryHandler(query *tgbotapi.CallbackQuery) {
	_, ok := h.storage.Exist(query.Message.Chat.ID, query.From.ID)
	if !ok {
		return
	}

	isUserHuman = true

	go h.DeleteMessage(
		query.Message.Chat.ID,
		query.Message.MessageID,
		DeleteMsgDelayZeroMin)
}
