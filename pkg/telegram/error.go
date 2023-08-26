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
	msg := tgbotapi.NewMessage(chatID, "Произошла неизвестная ошибка")
	
	switch err {
	case errUnauthorized:
		msg.Text = "Ты не авторизирован! Используй команду /start "
		b.bot.Send(msg)
	case errInvalidURL:
		msg.Text = "Это невалидная ссылка! "
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = "Не удалось сохранить ссылку. Попробуй позже."
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}