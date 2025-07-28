package handlers

import (
	"chekify/bot/constants"
	"chekify/bot/utils"
	tele "gopkg.in/telebot.v4"
)

func SendCheck(bot *tele.Bot) {
	bot.Handle(tele.OnCallback, func(ctx tele.Context) error {
		data := ctx.Callback().Data
		command := utils.ParseCallback(data)

		if command == constants.SendCheck {
			_ = ctx.Respond()

			return ctx.Send("Окей, надішли мені фото або скріншот чека")
		}

		return nil
	})
}
