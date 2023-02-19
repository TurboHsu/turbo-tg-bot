package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	basicCommand "github.com/TurboHsu/turbo-tg-bot/bot/modules/basic_command"
	glotrunner "github.com/TurboHsu/turbo-tg-bot/bot/modules/glot_runner"
	picsearch "github.com/TurboHsu/turbo-tg-bot/bot/modules/pic_search"
	whattoeat "github.com/TurboHsu/turbo-tg-bot/bot/modules/what_to_eat"
	"net/http"
	"time"

	// Modules end
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/TurboHsu/turbo-tg-bot/utils/config"
	"github.com/TurboHsu/turbo-tg-bot/utils/log"
)

func InitBot() {
	bot, err := gotgbot.NewBot(config.Config.APIKeys.BotToken, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: gotgbot.DefaultTimeout,
			APIURL:  gotgbot.DefaultAPIURL,
		},
	})
	log.HandleError(err)

	/*Init Module*/
	log.HandleError(whattoeat.Init())

	/*Create updater and dispatcher*/
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.HandleError(err)
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})

	dispatcher := updater.Dispatcher

	/* Modules start */
	dispatcher.AddHandler(handlers.NewCommand("start", basicCommand.StartHandler))
	dispatcher.AddHandler(handlers.NewCommand("help", basicCommand.HelpHandler))
	dispatcher.AddHandler(handlers.NewCommand("ping", basicCommand.PingHandler))
	dispatcher.AddHandler(handlers.NewCommand("info", basicCommand.InfoHandler))
	dispatcher.AddHandler(handlers.NewCommand("run", glotrunner.RunHandler))
	dispatcher.AddHandler(handlers.NewCommand("search", picsearch.SearchHandler))
	dispatcher.AddHandler(handlers.NewCommand("eat", whattoeat.CommandHandler))
	/* Modules end */

	/* Query */
	dispatcher.AddHandler(handlers.NewInlineQuery(basicCommand.QueryFilter, basicCommand.QueryResponse))
	/* Messaging */
	dispatcher.AddHandler(handlers.NewMessage(basicCommand.MessageFilter, basicCommand.MessageResponse))

	/* Start pulling */
	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: config.Config.Common.DropPendingUpdate,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})

	if err != nil {
		log.HandleError(err)
		return
	}

	log.HandleInfo("Congrats! Bot started successfully.")

	updater.Idle()
}
