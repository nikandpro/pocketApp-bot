package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user isn't autorized")
	errUnableToSave = errors.New("unable to save")
)


func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Default)
	
	switch err {
	case errUnauthorized:
		msg.Text = b.messages.Unauthorized
		b.bot.Send(msg)
	case errInvalidURL:
		msg.Text = b.messages.InvalidURL
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = b.messages.UnablebleToSave
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}