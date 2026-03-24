package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"transcriber/internal/tgBot"

	"github.com/go-telegram/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	params := []bot.Option{
		bot.WithDefaultHandler(tgBot.EchoHandler),
	}

	token := os.Getenv("BOT_TOKEN")

	b, err := bot.New(token, params...)
	if err != nil {
		log.Fatal(err)
	}

	b.Start(ctx)
}
