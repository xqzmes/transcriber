package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"transcriber/internal/telegram"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("BOT_TOKEN")

	opts := []bot.Option{
		bot.WithDefaultHandler(telegram.TextMessageHandler),
		bot.WithCallbackQueryDataHandler("btn_", bot.MatchTypePrefix, telegram.VoiceProcessorHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatal(err)
	}

	b.RegisterHandlerMatchFunc(telegram.MatchVoice, telegram.VoiceMessageHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, telegram.StartHandler)

	b.Start(ctx)
}
