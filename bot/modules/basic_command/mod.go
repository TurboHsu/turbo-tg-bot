package basicCommand

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	glotrunner "github.com/TurboHsu/turbo-tg-bot/bot/modules/glot_runner"
	picsearch "github.com/TurboHsu/turbo-tg-bot/bot/modules/pic_search"
	whattoeat "github.com/TurboHsu/turbo-tg-bot/bot/modules/what_to_eat"
	"github.com/TurboHsu/turbo-tg-bot/utils/log"
)

func StartHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(
		bot,
		fmt.Sprintf(`Hello, %s! This is a bot from TurboHsu.
In order to learn more, type /info.
For usage, type /help.`,
			ctx.EffectiveSender.FirstName(),
		),
		&gotgbot.SendMessageOpts{ParseMode: "html"},
	)
	log.HandleInfo(
		fmt.Sprintf("Dealed with [%s]'s start [%d].", ctx.EffectiveSender.FirstName(), ctx.EffectiveChat.Id),
	)
	log.HandleError(err)
	return nil
}

func HelpHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	exactParameter := ctx.Args()
	if len(exactParameter) == 1 {
		_, _ = ctx.EffectiveMessage.Reply(
			bot,
			`TurboHsu's Bot Help
			/search - Reply&Search for a specific image
			/info - Get information about the bot
			/help - Get help about the bot
			/run - Reply&Run a piece of code
			/eat - Get food recommendations
			To get more information about the specific command, type /help <command>`,
			&gotgbot.SendMessageOpts{},
		)
	} else {
		var msg string
		switch exactParameter[1] {
		case "search":
			msg = picsearch.GenerateHelp()
		case "info":
			msg = `/info -- Get information about the bot.`
		case "help":
			msg = `/help -- Get help about the bot.`
		case "run":
			msg = glotrunner.GenerateHelp()
		case "eat":
			msg = whattoeat.GenerateHelp()
		default:
			msg = "Introduction not found."
		}
		_, err := ctx.EffectiveMessage.Reply(bot, msg, &gotgbot.SendMessageOpts{})
		log.HandleError(err)
		log.HandleInfo(fmt.Sprintf("Dealed with [%s]'s help [%d].", ctx.EffectiveSender.FirstName(), ctx.EffectiveChat.Id))
	}
	return nil
}

func PingHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(bot, "Pong!", &gotgbot.SendMessageOpts{})
	log.HandleInfo(
		fmt.Sprintf("Dealed with [%s]'s ping [%d].", ctx.EffectiveSender.FirstName(), ctx.EffectiveChat.Id),
	)
	log.HandleError(err)
	return nil
}

func InfoHandler(bot *gotgbot.Bot, c *ext.Context) error {
	_, _ = c.EffectiveMessage.Reply(bot, `TurboHsu's Personal Bot.
	Github repo:[https://github.com/TurboHsu/turbo-tg-bot]
	Enjoy!`, &gotgbot.SendMessageOpts{})
	log.HandleInfo(fmt.Sprintf("Dealed with [%s]'s info [%d].", c.EffectiveSender.FirstName(), c.EffectiveChat.Id))
	return nil
}

func QueryFilter(query *gotgbot.InlineQuery) bool {
	return true
}

func QueryResponse(bot *gotgbot.Bot, ctx *ext.Context) error {
	results := make([]gotgbot.InlineQueryResult, 1)
	results[0] = whattoeat.EatQueryResultHandler(ctx)
	_, err := ctx.InlineQuery.Answer(bot, results, &gotgbot.AnswerInlineQueryOpts{CacheTime: 60})
	return err
}
