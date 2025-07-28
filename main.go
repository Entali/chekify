package main

import (
	"chekify/bot"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	b := bot.InitTelegramBot(os.Getenv("TELEGRAM_TOKEN"))
	b.Start()

	log.Info("Бот успішно запущений")
}
