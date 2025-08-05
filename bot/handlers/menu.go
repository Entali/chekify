package handlers

import (
	"chekify/bot/constants"
	tele "gopkg.in/telebot.v4"
)

func createMenu() tele.ReplyMarkup {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}

	addManuallyBtn := menu.Data("📝 Додати вручну", constants.CallbackAddManual)
	sendCheckBtn := menu.Data("📸 Надіслати чек", constants.CallbackSendCheck)
	postponedTransactionsBtn := menu.Data("⏳ Відкладені транзакції",
		constants.CallbackPostponed,
	)

	menu.Inline(
		menu.Row(addManuallyBtn),
		menu.Row(sendCheckBtn),
		menu.Row(postponedTransactionsBtn),
	)

	return menu
}

func Menu(bot *tele.Bot) {
	bot.Handle(constants.CommandMenu, func(ctx tele.Context) error {
		menu := createMenu()

		return ctx.Send("Що хочеш зробити?", &tele.SendOptions{
			ReplyMarkup: &menu,
		})
	})
}
