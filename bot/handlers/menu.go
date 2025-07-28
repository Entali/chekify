package handlers

import (
	tele "gopkg.in/telebot.v4"
)

func createMenu() tele.ReplyMarkup {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}

	addManuallyBtn := menu.Data("📝 Додати вручну", "add_manually", "payload")
	sendCheckBtn := menu.Data("📸 Надіслати чек", "send_check", "payload")
	postponedTransactionsBtn := menu.Data("⏳ Відкладені транзакції",
		"postponed_transactions",
		"payload")

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
