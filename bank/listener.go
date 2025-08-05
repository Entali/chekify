package bank

import (
	"gopkg.in/telebot.v4"
	"log"
	"time"
)

func StartTransactionListener(bot *telebot.Bot, token string, userID int64) {
	go func() {
		var lastTimeChecked = time.Now().Add(-15 * time.Minute) // Перший запуск

		for {
			txs, err := FetchRecentTransactions(token, lastTimeChecked)
			if err != nil {
				log.Println("Monobank error:", err)
				time.Sleep(30 * time.Second)
				continue
			}

			for _, tx := range txs {
				if isGrocery(tx.MCC) {
					msg := "💸 Ви тільки що витратили гроші в магазині. Надішліть, будь ласка, чек 🧾"
					_, _ = bot.Send(&telebot.User{ID: userID}, msg)
				}
			}

			lastTimeChecked = time.Now()
			time.Sleep(60 * time.Second)
		}
	}()
}

func isGrocery(mcc int) bool {
	return mcc == 5411 || mcc == 5422 || mcc == 5499
}
