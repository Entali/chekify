package main

import (
	"chekify/bot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot.StartTelegramBot(os.Getenv("TELEGRAM_TOKEN"))
}
