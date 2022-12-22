package whattoeat

import (
	tgbot "gopkg.in/telebot.v3"
)

func EatHandler(c tgbot.Context) error {
	c.Answer(&tgbot.QueryResponse{
		CacheTime: 60,
		Results:   tgbot.Results{&tgbot.ResultBase{}},
	})

	return nil
}
