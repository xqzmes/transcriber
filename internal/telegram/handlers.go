package telegram

import (
	"context"
	"log"
	"transcriber/internal/converter"
	"transcriber/internal/downloader"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	startMessage := "Привет!\nЯ бот-транскрибатор. Запиши или перешли мне аудиосообщение и я превращу его в текст!"
	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   startMessage,
	}); err != nil {
		log.Printf("Error sending startMessage: %v\n", err)
	}

	log.Print("Start message successfully send")
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

	if err = downloader.DownloadFile(filePath, fileLink); err != nil {
		log.Printf("Error downloading file: %v\n", err)
	}

	log.Printf("Downloading file: %v\n", filePath)
	log.Printf("File successfully downloaded: %v\n", update.Message.Voice.FileID)

	if _, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Голосовое сообщение успешно скачано!\nКакую информацию вы хотите из него получить?",
		ReplyMarkup: buildKeyboard(),
	}); err != nil {
		log.Printf("Error sending message: %v\n", err)
	}

	log.Print("Message with keyboard successfully send")
}

func VoiceProcessorHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if _, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	}); err != nil {
		log.Printf("Error answering callback query: %v\n", err)
	}

	switch update.CallbackQuery.Data {
	case "btn_option1":
		currentOptions[0] = !currentOptions[0]
	case "btn_option2":
		currentOptions[1] = !currentOptions[1]
	case "btn_option3":
		currentOptions[2] = !currentOptions[2]
	default:
		if _, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.Message.ID,
		}); err != nil {
			log.Printf("Error deleting message: %v\n", err)
		}

		if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "Обрабатываю аудиосообщение в соответствии с вашим выбором!",
		}); err != nil {
			log.Printf("Error sending message: %v\n", err)
		}

		if err := converter.ConvertAudio("./tmp/ogg/tmp.ogg", "./tmp/mp3/tmp.mp3"); err != nil {
			log.Printf("Error converting audio: %v\n", err)
		}

	}

	if _, err := b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		ReplyMarkup: buildKeyboard(),
	}); err != nil {
		log.Printf("Error editing message: %v\n", err)
	}

	log.Print("User choices successfully edited")
}

func TextMessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	textMessage := "К сожалению, я не умею работать с текстовыми сообщениями =(\n" +
		"Но вы всегда можете прислать мне аудиосообщения и я превращу его в текст =)"
	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   textMessage,
	}); err != nil {
		log.Printf("Error sending message: %v", err)
	}

	log.Print("Text message successfully send")
}
