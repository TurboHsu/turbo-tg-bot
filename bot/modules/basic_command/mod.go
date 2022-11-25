package basicCommand

import (
	"fmt"
	"strings"

	glotrunner "github.com/TurboHsu/turbo-tg-bot/bot/modules/glot_runner"
	"github.com/TurboHsu/turbo-tg-bot/utils/log"
	tgbot "gopkg.in/telebot.v3"
)

func StartHandler(c tgbot.Context) error {
	m := tgbot.Context.Message(c)
	err := c.Reply("Hello, " + m.Sender.FirstName + `!
	This is a bot from TurboHsu.
	In order to learn more, type /info.
	For usage, type /help.`)
	log.HandleInfo(
		fmt.Sprintf("Dealed with [%s]'s start [%d].", m.Sender.FirstName, m.Chat.ID),
	)
	log.HandleError(err)
	return nil
}

func HelpHandler(c tgbot.Context) error {
	m := tgbot.Context.Message(c)
	exactParameter := strings.Split(m.Text, " ")
	if len(exactParameter) == 1 {
		c.Reply(`TurboHsu's Bot Help
		/search - Reply&Search for a specific image
		/info - Get information about the bot
		/help - Get help about the bot
		/run - Reply&Run a piece of code
		To get more information about the specific command, type /help <command>`)
	} else {
		var msg string
		switch exactParameter[1] {
		case "search":
			msg = `/search <database> -- Reply to an image to search for it in a specific database.
			Available databases:
			saucenao

			Searches the SauceNAO API by default. `
		case "info":
			msg = `/info -- Get information about the bot.`
		case "help":
			msg = `/help -- Get help about the bot.`
		case "run":
			msg = glotrunner.GenerateHelp()
		default:
			msg = "Introduction not found."
		}
		err := c.Reply(msg)
		log.HandleError(err)
		log.HandleInfo("Dealed with [" + m.Sender.FirstName + "]'s help [" + fmt.Sprint(m.Chat.ID) + "].")
	}
	return nil
}

func PingHandler(c tgbot.Context) error {
	m := tgbot.Context.Message(c)
	err := c.Reply("Pong!")
	log.HandleInfo(
		fmt.Sprintf("Dealed with [%s]'s ping [%d].", m.Sender.FirstName, m.Chat.ID),
	)
	log.HandleError(err)
	return nil
}

func InfoHandler(c tgbot.Context) error {
	m := tgbot.Context.Message(c)
	c.Reply(`TurboHsu's Personal Bot.
		Github repo:[https://github.com/TurboHsu/turbo-tg-bot]
		Enjoy!`)
	log.HandleInfo(fmt.Sprintf("Dealed with [%s]'s info [%d].", m.Sender.FirstName, m.Chat.ID))
	return nil
}
