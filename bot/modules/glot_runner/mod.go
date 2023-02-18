package glotrunner

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func GenerateHelp() string {
	msg := `/run <language> <stdin> -- Reply to a piece of code to run it.
	Avaliable languages:`
	for _, lang := range glotLangs {
		msg += " " + lang.Name
	}
	return msg
}

func RunHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	var respText string
	if ctx.EffectiveMessage.ReplyToMessage != nil {
		exactParameter := ctx.Args()
		switch len(exactParameter) {
		case 1:
			respText = "Please specify a language."
		case 2:
			lang := exactParameter[1]
			code := ctx.EffectiveMessage.ReplyToMessage.Text
			respText = run(code, lang, "")
		case 3:
			lang := exactParameter[1]
			input := exactParameter[2]
			code := ctx.EffectiveMessage.ReplyToMessage.Text
			respText = run(code, lang, input)
		default:
			respText = "Too many parameters."
		}
	} else {
		respText = "In order to run code, please reply to a piece of code. ;)"
	}
	_, _ = ctx.EffectiveMessage.Reply(bot, respText, &gotgbot.SendMessageOpts{})

	return nil
}
