package tg_name

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func MainHandler(ctx context.Context, b *bot.Bot, update *models.Update) {}

func EchoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}

	log.Println(msg)
}
