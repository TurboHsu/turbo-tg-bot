package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tgbot "gopkg.in/tucnak/telebot.v2"
)

//Starts the bot
func startBot() {
	bot, err := tgbot.NewBot(tgbot.Settings{
		Token:  config.BotToken,
		Poller: &tgbot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}
	log.Println("Bot started")

	//Handle messages

	bot.Handle("/start", func(m *tgbot.Message) {
		bot.Send(m.Sender, "Hello, "+m.Sender.FirstName+`!
		This is a bot from TurboHsu.
		In order to learn more, type /info. 
		For usage, type /help.`)
		if !config.Silent {
			log.Printf("Dealed with [%s]'s start [%d].", m.Sender.FirstName, m.Chat.ID)
		}
	})

	bot.Handle("/help", func(m *tgbot.Message) {
		exactParameter := strings.Split(m.Text, " ")
		if len(exactParameter) == 1 {
			bot.Send(m.Sender, `TurboHsu's Bot Help 
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
				var langs string
				for i := 0; i < len(config.Glot); i++ {
					langs += config.Glot[i].Name + " "
				}
				msg = `/run <language> <input> -- Reply to a piece of code to run.
				Available languages: ` + langs
			default:
				msg = "Introduction not found."
			}
			bot.Send(m.Sender, msg)
		}
		if !config.Silent {
			log.Printf("Dealed with [%s]'s help [%d].", m.Sender.FirstName, m.Chat.ID)
		}
	})

	bot.Handle("/ping", func(m *tgbot.Message) {
		bot.Reply(m, "Pong!")
		if !config.Silent {
			log.Printf("Dealed with [%s]'s ping [%d].", m.Sender.Username, m.Chat.ID)
		}
	})

	bot.Handle("/search", func(m *tgbot.Message) {
		if m.IsReply() {
			exactParameter := strings.Split(m.Text, " ")
			var db, respText string
			if len(exactParameter) == 1 {
				db = "saucenao"
			} else {
				db = exactParameter[1]
			}
			switch db {
			case "saucenao":
				//Download photo
				bot.Download(&tgbot.File{FileID: m.ReplyTo.Photo.FileID}, fmt.Sprintf("./cache%s.jpg", m.ReplyTo.Photo.FileID))
				saucenaoJson := saucenaoSearch(m.ReplyTo.Photo.FileID)
				if config.Debug {
					bot.Send(m.Chat, saucenaoJson)
				}
				var result saucenaoJSONStruct
				var extURLs string
				json.Unmarshal([]byte(saucenaoJson), &result)
				for i := 0; i < len(result.Results[0].Data.ExtUrls); i++ {
					extURLs += fmt.Sprintf("%s\n", result.Results[0].Data.ExtUrls[i])
				}
				respText = fmt.Sprintf(`[%s]
					 Accuracy: %s%% 
					 Creator: %s 
					 Material: %s 
					 External URLs: 
					  %s 
					 Source: 
					  %s 
					 Account Limit: %v/%v %v/%v`, result.Results[0].Header.IndexName, result.Results[0].Header.Similarity, result.Results[0].Data.Creator, result.Results[0].Data.Material, extURLs, result.Results[0].Data.Source, result.Header.ShortRemaining, result.Header.ShortLimit, result.Header.LongRemaining, result.Header.LongLimit)
				//bot.Reply(m, &tgbot.Photo{File: tgbot.FromURL(result.Results[0].Header.Thumbnail)})
				//Telegram will show a thumbnail from URL, so there's no need to send it
				os.Remove(fmt.Sprintf("./cache%s.jpg", m.ReplyTo.Photo.FileID))
			default:
				respText = "Unknown database. Type \"/help search\" for more information."
			}
			bot.Reply(m, respText)
			if !config.Silent {
				log.Printf("Dealed with [%s]'s photo[%s] search[%d].", m.Sender.Username, m.ReplyTo.Photo.FileID, m.Chat.ID)
			}
		} else {
			bot.Reply(m, "Please reply to a photo to search.")
		}

	})

	bot.Handle("/info", func(m *tgbot.Message) {
		bot.Reply(m, `TurboHsu's Personal Bot.
		Github repo:[https://github.com/TurboHsu/turbo-tg-bot]
		Enjoy!`)
		if !config.Silent {
			log.Printf("Dealed with [%s]'s info [%d].", m.Sender.FirstName, m.Chat.ID)
		}
	})

	bot.Handle("/run", func(m *tgbot.Message) {
		if m.IsReply() {
			exactParameter := strings.Split(m.Text, " ")
			var respText string
			switch len(exactParameter) {
			case 1:
				respText = "Please specify a language."
			case 2:
				lang := exactParameter[1]
				respText = glotRun(m.ReplyTo.Text, lang, "")
			case 3:
				lang := exactParameter[1]
				input := exactParameter[2]
				respText = glotRun(m.ReplyTo.Text, lang, input)
			default:
				respText = "Too many parameters."
			}
			bot.Reply(m, respText)
			if !config.Silent {
				log.Printf("Dealed with [%s]'s run [%d].", m.Sender.FirstName, m.Chat.ID)
			}

		} else {
			bot.Reply(m, "Please reply to a piece of code to run.")
		}
	})

	bot.Start()
}
