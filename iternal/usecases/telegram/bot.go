package telegram

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot{
	return &Bot{bot: bot }
}

func (b *Bot) StartBot() error{
	log.Printf("%s is working now", b.bot.Self.UserName)


	updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 30

    
    updates := b.bot.GetUpdatesChan(updateConfig)
	
	for update := range updates{
		if update.Message == nil{
			continue
		}
	}
	
	return nil
}