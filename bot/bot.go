package bot

import (
	"time"

	// Modules start
	basicCommand "github.com/TurboHsu/turbo-tg-bot/bot/modules/basic_command"
	glotrunner "github.com/TurboHsu/turbo-tg-bot/bot/modules/glot_runner"
	picsearch "github.com/TurboHsu/turbo-tg-bot/bot/modules/pic_search"

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
	log.HandleError(err)

	/*Some modules may need to access bot API*/
	picsearch.BotFetcher(bot)

	/* Modules start */
	bot.Handle("/start", basicCommand.StartHandler)
	bot.Handle("/help", basicCommand.HelpHandler)
	bot.Handle("/ping", basicCommand.PingHandler)
	bot.Handle("/info", basicCommand.InfoHandler)
	bot.Handle("/run", glotrunner.RunHandler)
	bot.Handle("/search", picsearch.SearchHandler)
	//bot.Handle(tgbot.OnQuery, whattoeat.EatHandler)
	/* Modules end */

	log.HandleInfo("Congrats! Bot started successfully.")

	bot.Start()
}
