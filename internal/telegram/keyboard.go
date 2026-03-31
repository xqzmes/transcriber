package telegram

import "github.com/go-telegram/bot/models"

var currentOptions = []bool{false, false, false}

func buildKeyboard() models.ReplyMarkup {
	keyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: buttonText("Расшифровка", currentOptions[0]), CallbackData: "btn_option1"},
			},
			{
				{Text: buttonText("Краткое содержание", currentOptions[1]), CallbackData: "btn_option2"},
			},
			{
				{Text: buttonText("Основные пойнты", currentOptions[2]), CallbackData: "btn_option3"},
			},
			{
				{Text: "Select", CallbackData: "btn_select"},
			},
		},
	}

	return keyboard
}

func buttonText(text string, option bool) string {
	if option {
		return "✅ " + text
	}

	return "❌ " + text
}
