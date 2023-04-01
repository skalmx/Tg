package telegram

import (
	"strings"
	"tg/iternal/usecases/webapi"
	"time"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	commandGetBreeds = "list"
	commandRndFact = "randomfact"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandGetBreeds:
		return b.handleGetBreedsCommand(message)
	case commandRndFact:
		return b.handleRndFactCommand(message)
	default:
		return b.handleUknownCommand(message)
	}

}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, I am a bot who loves dogs very much. Let me tell you what I can do.\nWith /list you can see all the breeds of dogs that I know. You can also send me the breed of the dog and I will tell you what I know about it.\nYes, I forgot, I can tell you a lot of facts about dogs, for this use /randomfact!")
	var startKeyboadrd = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/list"),
			tgbotapi.NewKeyboardButton("/randomfact"),
	    ),
	)
	msg.ReplyMarkup = startKeyboadrd

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUknownCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "unkown command")
	
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleRndFactCommand(message *tgbotapi.Message) error {
	client, err := webapi.NewClient(time.Second * 5)
	if err != nil {
		return err
	}

	text, err := client.RndFact()
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleGetBreedsCommand(message *tgbotapi.Message) error {

	var lettersChoise = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("A", "A"),
			tgbotapi.NewInlineKeyboardButtonData("B", "B"),
			tgbotapi.NewInlineKeyboardButtonData("C", "C"),
			tgbotapi.NewInlineKeyboardButtonData("D", "D"),
			tgbotapi.NewInlineKeyboardButtonData("E", "E"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("F", "F"),
			tgbotapi.NewInlineKeyboardButtonData("G", "G"),
			tgbotapi.NewInlineKeyboardButtonData("H", "H"),
			tgbotapi.NewInlineKeyboardButtonData("I", "I"),
			tgbotapi.NewInlineKeyboardButtonData("J", "J"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("K", "K"),
			tgbotapi.NewInlineKeyboardButtonData("L", "L"),
			tgbotapi.NewInlineKeyboardButtonData("M", "M"),
			tgbotapi.NewInlineKeyboardButtonData("N", "N"),
			tgbotapi.NewInlineKeyboardButtonData("O", "O"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("P", "P"),
			tgbotapi.NewInlineKeyboardButtonData("R", "R"),
			tgbotapi.NewInlineKeyboardButtonData("S", "S"),
			tgbotapi.NewInlineKeyboardButtonData("T", "T"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("V", "V"),
			tgbotapi.NewInlineKeyboardButtonData("W", "W"),
			tgbotapi.NewInlineKeyboardButtonData("X", "X"),
			tgbotapi.NewInlineKeyboardButtonData("Y", "Y"),
		),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, "For a convenient search, select the letter that begins with the breed of dog you want to know about")
	msg.ReplyMarkup = lettersChoise

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessages(message *tgbotapi.Message) error {
	client, err := webapi.NewClient(time.Second * 5)
	if err != nil {
		return err
	}
	answer, url, _ := client.BreedInfo(message.Text)
	
	msg := tgbotapi.NewMessage(message.Chat.ID, answer)
	msg.Entities = append(msg.Entities, tgbotapi.MessageEntity{
		Offset: 0,
		Length: len(strings.Split(answer, "\n")[0]), // get a len of a first row to bold it
		Type:   "bold",
	})
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}
	if url == ""{
		return nil // checking if the response contains a url , if not user sent invalid message and already recieve message about it
	}
	photo := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileURL(url)) //way to send a photo
	_, err = b.bot.Send(photo)

	return err
}
func (b *Bot) handleCallbacks(cb *tgbotapi.CallbackQuery) error {
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := b.bot.Request(callback); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(cb.Message.Chat.ID, "Dont forget to copy name of the breed and send it to me!")
	msg.Entities = append(msg.Entities, tgbotapi.MessageEntity{
		Offset: 0,
		Length: len(msg.Text),
		Type:   "bold",
	})
	_, err := b.bot.Send(msg)
	if err != nil{
		return err
	}

	client, err := webapi.NewClient(time.Second * 5)
	if err != nil {
		return err
	}
	breeds, err := client.FindBreed(cb.Data[0])
	if err != nil {
		return err
	}

	text := strings.Join(breeds, "\n")
	msg = tgbotapi.NewMessage(cb.Message.Chat.ID, text)

	offset := 0
	for _, value := range breeds {
		msg.Entities = append(msg.Entities, tgbotapi.MessageEntity{
			Offset: offset,
			Length: len(value),
			Type:   "code",
		})
		offset += len(value) + 1 //add +1 cause of /n
	}

	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
