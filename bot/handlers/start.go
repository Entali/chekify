package handlers

import tele "gopkg.in/telebot.v4"

func Start(bot *tele.Bot) {
	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Вітаю. Я твій розумник для обліку витрат")
	})
}
