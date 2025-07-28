package bot

import (
	"chekify/bot/handlers"
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
	handlers.Start(bot)
	handlers.Menu(bot)
	bot.Start()
}
