package main

import (
	"log"
	"os"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
    
    if err := godotenv.Load(); err != nil {
        log.Fatal("No .env file found")
    }

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGBOTAPI_TOKEN")) // get token from .env file 
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
	
	updateConfig := tgbotapi.NewUpdate(0)

    
    updateConfig.Timeout = 30

    
    updates := bot.GetUpdatesChan(updateConfig)

    
    for update := range updates {
       
        if update.Message == nil {
            continue
        }

        
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
       
        

       
        if _, err := bot.Send(msg); err != nil {
            
            panic(err)
		}
}  }