package telegram

import (
	"context"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

const (
	commandStart = "start"
	startReplyTemplate = "Привет! Чтобы сохранять ссылки в своем Pocket аккаунте, для начала тебе необходимо дать мне на это доступ. Для этого переходи по ссылке: \n%s"
	replyAlreadyAutorized = "Ты уже авторизирован. Присылай ссылку, а я ее сохраню."
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnkownCommand(message)
	}
	
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	// log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ссылка успешно сохранена!")
	
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		msg.Text = "Это невалидная ссылка!"
		_, err = b.bot.Send(msg)
		return err
	}
	
	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		msg.Text = "Ты не авторизирован!"
		_, err = b.bot.Send(msg)
		return err
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL: message.Text,
	}); err != nil {
		msg.Text = "Не удалось сохранить ссылку. Попробуй позже."
		_, err = b.bot.Send(msg)
		return err
	}

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAutorizationProcess(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, replyAlreadyAutorized)
	_, err = b.bot.Send(msg)
	return err
	
}

func (b *Bot) handleUnkownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
		
	_, err := b.bot.Send(msg)
	return err
}