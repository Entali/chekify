package handlers

import (
	"chekify/bot/constants"
	tele "gopkg.in/telebot.v4"
)

func createMenu() tele.ReplyMarkup {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}

	addManuallyBtn := menu.Data("📝 Додати вручну", constants.AddManually)
	sendCheckBtn := menu.Data("📸 Надіслати чек", constants.SendCheck)
	postponedTransactionsBtn := menu.Data("⏳ Відкладені транзакції",
		constants.PostponedTransactions,
	)

	menu.Inline(
		menu.Row(addManuallyBtn),
		menu.Row(sendCheckBtn),
		menu.Row(postponedTransactionsBtn),
	)

	return menu
}

func Menu(bot *tele.Bot) {
	bot.Handle("/menu", func(ctx tele.Context) error {
		menu := createMenu()

		return ctx.Send("Що хочеш зробити?", &tele.SendOptions{
			ReplyMarkup: &menu,
		})
	})
}
