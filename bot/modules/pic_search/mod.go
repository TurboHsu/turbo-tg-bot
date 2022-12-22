package picsearch

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TurboHsu/turbo-tg-bot/utils/log"
	tgbot "gopkg.in/telebot.v3"
)

var bot *tgbot.Bot

func GenerateHelp() string {
	return `/search <database> -- Reply to an image to search for it in a specific database.
	Available databases:
	saucenao

	Searches the SauceNAO API by default. `
}

func BotFetcher(b *tgbot.Bot) {
	bot = b
}

func SearchHandler(c tgbot.Context) error {
	m := tgbot.Context.Message(c)
	if m.IsReply() && m.ReplyTo.Photo != nil {
		//Parameters: /search <count> <database>
		//database defaults to saucenao (now only have one so i won't intemplate it)
		//count defaults to 1
		//If count is not a number, it will be treated as database
		exactParameter := strings.Split(m.Text, " ")
		if len(exactParameter) >= 1 && len(exactParameter) <= 3 {
			var count int = 1
			//Parse count
			if len(exactParameter) >= 2 {
				cnt, err := strconv.Atoi(exactParameter[1])
				if err != nil {
					c.Reply("<count> is invalid. Setting it to default value.")
				} else {
					count = cnt
				}
			}
			//Download the file
			err := bot.Download(&m.ReplyTo.Photo.File, "./cache/"+m.ReplyTo.Photo.UniqueID+".jpeg")
			log.HandleError(err)

			//Search it
			res := searchSauseNAO("./cache/"+m.ReplyTo.Photo.UniqueID+".jpeg", count)

			//Delete it
			err = os.Remove("./cache/" + m.ReplyTo.Photo.UniqueID + ".jpeg")
			log.HandleError(err)

			//Response something
			for i := 0; i < len(res.Results); i++ {
				c.Reply(&tgbot.Photo{File: tgbot.FromURL(res.Results[i].Header.Thumbnail),
					Caption: fmt.Sprintf("[%s] %s %s\n", res.Results[i].Header.Similarity,
						res.Results[i].Data.Title,
						res.Results[i].Data.ExtUrls[0])})
			}

		} else if len(exactParameter) == 2 && exactParameter[1] == "limit" {
			c.Reply(fmt.Sprintf("Account Limit:\nIn short term: %d/%d\nIn long term: %d/%s",
				saucenaoUser.ShortRemaining, saucenaoUser.ShortLimit,
				saucenaoUser.LongRemaining, saucenaoUser.LongLimit,
			))

		} else {
			c.Reply("Too many parameters.")
		}
	} else {
		exactParameter := strings.Split(m.Text, " ")
		if len(exactParameter) == 2 && exactParameter[1] == "limit" {

		} else {
			c.Reply("In order to search for an image, please reply to an image. ;)")
		}
	}
	return nil
}
