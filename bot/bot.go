package bot

import (
	tele "gopkg.in/telebot.v4"
	"log"
	"time"
)

func StartTelegramBot(token string) {
	telegramBotSettings := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	telegramBot, err := tele.NewBot(telegramBotSettings)

	if err != nil {
		log.Fatal(err)
		return
	}

	telegramBot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Вітаю. Я твій розумник для обліку витрат")
	})

	telegramBot.Start()
}
