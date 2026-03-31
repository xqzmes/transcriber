package telegramHandlers

import (
	"context"
	"log"
	"transcriber/internal/downloader"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	startMessage := "Привет!\nЯ бот-транскрибатор. Запиши или перешли мне аудиосообщение и я превращу его в текст!"
	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   startMessage,
	})
	if err != nil {
		log.Printf("Error sending startMessage: %v\n", err)
	}

	log.Printf("Sending startMessage: %v\n", msg)
}

func VoiceMessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	file, err := b.GetFile(ctx, &bot.GetFileParams{
		FileID: update.Message.Voice.FileID,
	})
	if err != nil {
		log.Printf("Error getting file: %v\n", err)
	}

	fileLink := b.FileDownloadLink(file)

	filePath := "./tmp/ogg/tmp.ogg"

	if err := downloader.DownloadFile(filePath, fileLink); err != nil {
		log.Printf("Error downloading file: %v\n", err)
	}

	log.Printf("Downloading file: %v\n", filePath)
	log.Printf("File successfully downloaded: %v\n", update.Message.Voice.FileID)
}

func TextMessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	textMessage := "К сожалению, я не умею работать с текстовыми сообщениями =(\n" +
		"Но вы всегда можете прислать мне аудиосообщения и я превращу его в текст =)"
	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   textMessage,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}

	log.Printf("Sending message: %v\n", msg)
}

func MatchVoice(update *models.Update) bool {
	if update.Message.Voice != nil {
		return true
	}

	return false
}
