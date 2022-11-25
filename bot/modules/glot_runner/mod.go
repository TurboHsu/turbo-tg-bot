package glotrunner

import (
	"strings"

	tgbot "gopkg.in/telebot.v3"
)

func GenerateHelp() string {
	msg := `/run <language> <stdin> -- Reply to a piece of code to run it.
	Avaliable languages:`
	for _, lang := range glotLangs {
		msg += " " + lang.Name
	}
	return msg
}

func RunHandler(c tgbot.Context) error {
	m := tgbot.Context.Message(c)
	if m.IsReply() {
		exactParameter := strings.Split(m.Text, " ")
		var respText string
		switch len(exactParameter) {
		case 1:
			respText = "Please specify a language."
		case 2:
			lang := exactParameter[1]
			respText = run(m.ReplyTo.Text, lang, "")
		case 3:
			lang := exactParameter[1]
			input := exactParameter[2]
			respText = run(m.ReplyTo.Text, lang, input)
		default:
			respText = "Too many parameters."
		}
		c.Reply(respText)
	} else {
		c.Reply("In order to run code, please reply to a piece of code. ;)")
	}

	return nil
}
