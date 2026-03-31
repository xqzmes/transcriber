package telegram

import "github.com/go-telegram/bot/models"

func MatchVoice(update *models.Update) bool {
	if update.Message.Voice != nil {
		return true
	}

	return false
}
