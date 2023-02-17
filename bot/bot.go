package bot

import (
	"time"

	// Modules start
	basicCommand "github.com/TurboHsu/turbo-tg-bot/bot/modules/basic_command"
	glotrunner "github.com/TurboHsu/turbo-tg-bot/bot/modules/glot_runner"
	picsearch "github.com/TurboHsu/turbo-tg-bot/bot/modules/pic_search"
	whattoeat "github.com/TurboHsu/turbo-tg-bot/bot/modules/what_to_eat"

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

	/*Init Module*/
	log.HandleError(whattoeat.Init())

	/*Telegram Bot API Provider*/
	picsearch.BotFetcher(bot)

	/* Modules start */
	bot.Handle("/start", basicCommand.StartHandler)
	bot.Handle("/help", basicCommand.HelpHandler)
	bot.Handle("/ping", basicCommand.PingHandler)
	bot.Handle("/info", basicCommand.InfoHandler)
	bot.Handle("/run", glotrunner.RunHandler)
	bot.Handle("/search", picsearch.SearchHandler)
	/* Modules end */

	/* Query */
	bot.Handle(tgbot.OnQuery, func(c tgbot.Context) error {
		//Gather query results
		results := make(tgbot.Results, 1)
		results[0] = whattoeat.EatQueryResultHandler(c)
		results[0].SetResultID("0")

		return c.Answer(&tgbot.QueryResponse{
			Results:   results,
			CacheTime: 60,
		})
	})

	log.HandleInfo("Congrats! Bot started successfully.")

	bot.Start()
}
