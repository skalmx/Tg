package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) StartBot() error {

	log.Printf("%s is working now", b.bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.CallbackQuery != nil {
			if err := b.handleCallbacks(update.CallbackQuery); err != nil {
				log.Print(err)
			}
			continue
		}
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				log.Print(err)
			}
			continue
		}
		if update.Message != nil {
			if err := b.handleMessages(update.Message); err != nil {
				log.Print(err)
			}
		}
		if update.EditedMessage != nil { // this check is added because after user edit his msg, bot ll not work at all
			continue
		}
	}
	return nil
}
