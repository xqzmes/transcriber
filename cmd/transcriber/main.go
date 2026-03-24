package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"transcriber/internal/telegramHandlers"

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
		bot.WithDefaultHandler(telegramHandlers.EchoHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatal(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, telegramHandlers.StartHandler)

	b.Start(ctx)
}
