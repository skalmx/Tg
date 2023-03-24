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

		if update.CallbackQuery != nil{
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := b.bot.Request(callback); err != nil{
				log.Print(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
	
			_, err := b.bot.Send(msg); if err != nil {
				log.Print(err)
			}
			continue
		}
		// if update.Message == nil{
		// 	continue
		// }
		
		if update.Message.IsCommand(){
			if err := b.handleCommand(update.Message); err != nil{
				log.Print(err)
			}
			continue
		}
		
		
	}
	
	return nil
}