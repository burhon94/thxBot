package main

import (
	"fmt"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var token = "1521513385:AAGUpQxRxmV8USIsGyKPFRo6-S_xJOsCwBo"

func main() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Error set logFile: %v", err))
	}
	log.SetOutput(logFile)
	log.Printf("Log inited")

	bot, err := tgBotApi.NewBotAPI(token)
	if err != nil {
		log.Panic(fmt.Sprintf("Error: can't accept bot Token: %v", err))
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	fmt.Println("bot is activated success...")

	channel := tgBotApi.NewUpdate(0)
	channel.Timeout = 60

	updates, err := bot.GetUpdatesChan(channel)
	handlerLoop(updates, bot)
}

func handlerLoop(updates tgBotApi.UpdatesChannel, bot *tgBotApi.BotAPI) {
		for update := range updates {

			//var isCommand bool
			isBot := update.Message.From.IsBot
			if isBot {
				continue
			}
			userId := update.Message.From.ID
			textFromChat := update.Message.Text
			userName := update.Message.From.UserName
			firstName := update.Message.From.FirstName
			lastName := update.Message.From.LastName
			update.Message.From.String()

			log.Printf("[userName: %s | userFirstName: %s | userLanstName: %s] | command: %s", userName, firstName, lastName, textFromChat)

			log.Printf("tgUserID: %v", userId)
		}
}