package bot

import (
	"time"

	// Modules start
	basicCommand "github.com/TurboHsu/turbo-tg-bot/bot/modules/basic_command"
	glotrunner "github.com/TurboHsu/turbo-tg-bot/bot/modules/glot_runner"

	// Modules end
	"github.com/TurboHsu/turbo-tg-bot/utils/config"
	"github.com/TurboHsu/turbo-tg-bot/utils/log"
	tgbot "gopkg.in/telebot.v3"
)

func InitBot() {
	bot, err := tgbot.NewBot(tgbot.Settings{
		Token:  config.Config.APIKeys.BotToken,
		Poller: &tgbot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.HandleError(err)
	}

	/* Modules start */
	bot.Handle("/start", basicCommand.StartHandler)
	bot.Handle("/help", basicCommand.HelpHandler)
	bot.Handle("/ping", basicCommand.PingHandler)
	bot.Handle("/info", basicCommand.InfoHandler)
	bot.Handle("/run", glotrunner.RunHandler)
	/* Modules end */

	log.HandleInfo("Congrats! Bot started successfully.")

	bot.Start()
}
