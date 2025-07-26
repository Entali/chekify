package bot

import (
	tele "gopkg.in/telebot.v4"
	"log"
	"time"
)

func InitTelegramBot(token string) (bot *tele.Bot) {
	telegramBotSettings := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tele.NewBot(telegramBotSettings)

	if err != nil {
		log.Fatal(err)
		return
	}

	handleStart(bot)
	bot.Start()

	return bot
}

func handleStart(bot *tele.Bot) {
	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Вітаю. Я твій розумник для обліку витрат")
	})
}
