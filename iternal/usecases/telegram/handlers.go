package telegram

import (
	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

 const (
	commandStart = "start"
	commandGetBreeds = "list"
// commandRndFact = "fact"
// commandRndImage = "image"
// commandBreedImage = "breed"
 
)



func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandGetBreeds:
		return b.handleGetBreedsCommand(message)
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
func (b *Bot) handleGetBreedsCommand(message *tgbotapi.Message) error{
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("A","A"),
			tgbotapi.NewInlineKeyboardButtonData("B","B"),
			tgbotapi.NewInlineKeyboardButtonData("C","C"),
			tgbotapi.NewInlineKeyboardButtonData("D","D"),
			tgbotapi.NewInlineKeyboardButtonData("E","E"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("F","F"),
			tgbotapi.NewInlineKeyboardButtonData("G","G"),
			tgbotapi.NewInlineKeyboardButtonData("H","H"),
			tgbotapi.NewInlineKeyboardButtonData("I","I"),
			tgbotapi.NewInlineKeyboardButtonData("J","J"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("K","K"),
			tgbotapi.NewInlineKeyboardButtonData("L","L"),
			tgbotapi.NewInlineKeyboardButtonData("M","M"),
			tgbotapi.NewInlineKeyboardButtonData("N","N"),
			tgbotapi.NewInlineKeyboardButtonData("O","O"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("P","P"),
			tgbotapi.NewInlineKeyboardButtonData("R","R"),
			tgbotapi.NewInlineKeyboardButtonData("S","S"),
			tgbotapi.NewInlineKeyboardButtonData("T","T"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("V","V"),
			tgbotapi.NewInlineKeyboardButtonData("W","W"),
			tgbotapi.NewInlineKeyboardButtonData("X","X"),
			tgbotapi.NewInlineKeyboardButtonData("Y","Y"),
		),

	)
	
	msg := tgbotapi.NewMessage(message.Chat.ID, "list command")
	msg.ReplyMarkup = numericKeyboard

	_, err := b.bot.Send(msg)
	return err

	
}

func (b *Bot) handleButtons(cb *tgbotapi.CallbackQuery) error{
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := b.bot.Request(callback); err != nil{
				return err
		}

		msg := tgbotapi.NewMessage(cb.Message.Chat.ID, cb.Data)
	
			_, err := b.bot.Send(msg); if err != nil {
				return err
		}
			return nil
}
