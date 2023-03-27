package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"tg/iternal/usecases/telegram"
	"tg/iternal/usecases/webapi"


	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)


func main() {

	resp, err := http.Get("https://api.thedogapi.com/v1/breeds?limit=264&page=0")
		if err != nil{
			log.Print(err)
		}
	
	body, err := io.ReadAll(resp.Body)
		if err != nil{
			log.Print(err)
		}
	fmt.Println(string(body))	
	
	var responce webapi.BreedInfo
	if err := json.Unmarshal(body, &responce); err != nil{
		log.Print(err)
	}
	for _, r := range responce{
		fmt.Println(r.Name + "")
	}

	
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("TGBOTAPI_TOKEN")) // get token from .env file
	if err != nil {
		log.Fatal(err)
	}

	botAPI.Debug = false

	bot := telegram.NewBot(botAPI)

	if err := bot.StartBot(); err != nil {
		log.Fatal(err)
	}

}
