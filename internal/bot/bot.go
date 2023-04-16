package bot

import (
	"github.com/elidotexe/esme/internal/bot/handlers"
	"github.com/elidotexe/esme/internal/logger"
	"github.com/elidotexe/esme/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	zaplog "go.uber.org/zap"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	logger   *logger.Logger
	updates  tgbotapi.UpdatesChannel
	storage  storage.Storage
	handlers *handlers.Handlers
}

// NewBot creates a new instance of the bot, initializes it with the provided token,
// logger and message handlers, and returns the bot along with an error (if any).
func NewBot(token string, logger *logger.Logger) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	tgbotapi.NewUpdate(0)

	logger.Info("Authorized on Telegram", zaplog.String("bot", bot.Self.UserName))

	bot.Debug = true

	updates := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 60,
	})

	s := storage.NewStorage()

	h, err := handlers.Initialize(bot, logger, s)
	if err != nil {
		logger.Error("Error initializing handlers", zaplog.Error(err))
		return nil, err
	}

	b := &Bot{
		bot:      bot,
		logger:   logger,
		updates:  updates,
		storage:  s,
		handlers: h,
	}

	return b, nil
}

// Start starts the bot and listens for incoming messages. It uses the bot's
// handlers to route the incoming messages to their appropriate functions. If the
// message is of an unknown type, it logs an error message.
func (b *Bot) Start() error {
	b.logger.Info("Bot has been successfully started...")

	for u := range b.updates {
		switch {
		case u.Message == nil || u.CallbackQuery == nil:
			continue
		case u.Message.Command() == "info":
			b.handlers.OnInfoCommand(u.Message)
		case u.Message != nil:
			b.handlers.HandleUpdate(u)
		case u.CallbackQuery != nil:
			b.handlers.HandleUpdate(u)
		default:
			b.logger.Infof("Unknown message type: %T", u.Message)
		}
	}

	return nil
}
