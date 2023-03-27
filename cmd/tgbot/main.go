package main

import (
	"log"
	"os"
	"tg/iternal/usecases/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)


func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("TGBOTAPI_TOKEN")) // get token from .env file
	if err != nil {
		log.Fatal(err)
	}

	botAPI.Debug = true

	bot := telegram.NewBot(botAPI)

	if err := bot.StartBot(); err != nil {
		log.Fatal(err)
	}

}
