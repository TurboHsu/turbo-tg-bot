package whattoeat

import (
	tgbot "gopkg.in/telebot.v3"
)

/*
	Structure:
	User -> User Group -> Eat Data -> Random generate -> send
*/

func EatQueryResultHandler(c tgbot.Context) *tgbot.ArticleResult {
	description, text := foodGenerate(c.Sender().ID)
	return &tgbot.ArticleResult{
		Title:       "Decide what to eat!",
		Description: description,
		Text:        text,
		ThumbURL:    "https://api.tcloud.site/static/rice.jpg",
		ThumbWidth:  612,
		ThumbHeight: 455,
	}
}

func foodGenerate(uid int64) (description string, text string) {
	return "nil", "nil"
}
