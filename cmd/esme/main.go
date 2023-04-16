package main

import (
	"log"
	"os"

	b "github.com/elidotexe/esme/internal/bot"
	"github.com/elidotexe/esme/internal/logger"
	"github.com/joho/godotenv"
	zaplog "go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	logger, err := logger.NewLogger(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	botToken := os.Getenv("BOT_TOKEN")
	bot, err := b.NewBot(botToken, logger)
	if err != nil {
		logger.Fatal("failed to start bot", zaplog.Error(err))
	}

	err = bot.Start()
	if err != nil {
		logger.Fatal("failed to start bot", zaplog.Error(err))
	}
}
