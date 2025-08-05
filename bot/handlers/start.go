package handlers

import (
	"chekify/bot/constants"
	tele "gopkg.in/telebot.v4"
)

func Start(bot *tele.Bot) {
	bot.Handle(constants.CommandStart, func(ctx tele.Context) error {
		return ctx.Send("Вітаю. Я твій розумник для обліку витрат")
	})
}
