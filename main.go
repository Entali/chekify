package main

import (
	"chekify/bank"
	"chekify/bot"
	"chekify/env"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	monobankToken := os.Getenv("MONOBANK_TOKEN")
	userID := env.GetTelegramID()

	telegramBot := bot.InitTelegramBot(telegramToken)

	// Запускаємо прослуховування Monobank у фоні
	go bank.StartTransactionListener(telegramBot, monobankToken, userID)

	log.Info("Бот успішно запущений")
	telegramBot.Start()
}
