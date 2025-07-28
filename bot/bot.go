package bot

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	"log"
	"time"
)

func InitTelegramBot(token string) (bot *tele.Bot) {
	telegramBotSettings := tele.Settings{
		Token:   token,
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: true,
	}
	bot, err := tele.NewBot(telegramBotSettings)

	if err != nil {
		log.Fatal(err)
		return
	}

	registerHandlers(bot)

	return bot
}

func registerHandlers(bot *tele.Bot) {
	handleStart(bot)
	handleMenu(bot)
	bot.Start()
}

func createMenu() tele.ReplyMarkup {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}

	addManuallyBtn := menu.Data("📝 Додати вручну", "add_manually", "payload")
	sendCheckBtn := menu.Data("📸 Надіслати чек", "send_check", "payload")
	postponedTransactionsBtn := menu.Data("⏳ Відкладені транзакції",
		"postponed_transactions",
		"інший_пейлоуд")

	menu.Inline(
		menu.Row(addManuallyBtn),
		menu.Row(sendCheckBtn),
		menu.Row(postponedTransactionsBtn),
	)

	return menu
}

func handleStart(bot *tele.Bot) {
	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Вітаю. Я твій розумник для обліку витрат")
	})
}

func handleMenu(bot *tele.Bot) {
	fmt.Println("Run menu")
	bot.Handle("/menu", func(ctx tele.Context) error {
		menu := createMenu()

		return ctx.Send("Що хочеш зробити?", &tele.SendOptions{
			ReplyMarkup: &menu,
		})
	})
}
