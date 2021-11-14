package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
		bot.Send(m.Sender, "Hello, "+m.Sender.FirstName+"!\n This is a bot from TurboHsu.\n In order to learn more, type /info.")
	})

	bot.Handle("/ping", func(m *tgbot.Message) {
		bot.Reply(m, "Pong!")
		if !config.Silent {
			log.Printf("Dealed with [%s]'s ping [%d].", m.Sender.Username, m.Chat.ID)
		}
	})

	bot.Handle(tgbot.OnPhoto, func(m *tgbot.Message) {
		//Download photo
		bot.Download(&tgbot.File{FileID: m.Photo.FileID}, fmt.Sprintf("./cache%s.jpg", m.Photo.FileID))
		saucenaoJson := saucenaoSearch(m.Photo.FileID)
		var result saucenaoJSONStruct
		var extURLs string
		json.Unmarshal([]byte(saucenaoJson), &result)
		for i := 0; i < len(result.Results[0].Data.ExtUrls); i++ {
			extURLs += fmt.Sprintf("%s\n", result.Results[0].Data.ExtUrls[i])
		}
		respText := fmt.Sprintf("[%s] \n Accuracy: %s%% \n Creator: %s \n Material: %s \n External URLs: \n %s \n Source: \n %s \n Account Limit: %v/%v %v/%v", result.Results[0].Header.IndexName, result.Results[0].Header.Similarity, result.Results[0].Data.Creator, result.Results[0].Data.Material, extURLs, result.Results[0].Data.Source, result.Header.ShortRemaining, result.Header.ShortLimit, result.Header.LongRemaining, result.Header.LongLimit)
		//bot.Reply(m, &tgbot.Photo{File: tgbot.FromURL(result.Results[0].Header.Thumbnail)})
		//Telegram will show a thumbnail from URL, so there's no need to send it
		bot.Reply(m, respText)
		if config.Debug {
			bot.Send(m.Chat, saucenaoJson)
		}
		os.Remove(fmt.Sprintf("./cache%s.jpg", m.Photo.FileID))
		if !config.Silent {
			log.Printf("Dealed with [%s]'s photo[%s] search[%d].", m.Sender.Username, m.Photo.FileID, m.Chat.ID)
		}
	})

	bot.Handle("/info", func(m *tgbot.Message) {
		bot.Reply(m, "TurboHsu's Personal Bot.\nGithub repo:[https://github.com/TurboHsu/turbo-tg-bot]\nEnjoy!")
	})

	bot.Start()
}
