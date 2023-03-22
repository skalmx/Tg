package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

 const (
	commandStart = "start"

// commandRndFact = "fact"
// commandRndImage = "image"
// commandBreedImage = "breed"
// commandGetBreed = "list"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUknownCommand(message)
	}

}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Its start command hello <3")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I dont know this command sry :(")
	_, err := b.bot.Send(msg)
	return err
}
