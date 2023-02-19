package picsearch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/TurboHsu/turbo-tg-bot/utils/log"
)

func GenerateHelp() string {
	return `/search <database> -- Reply to an image to search for it in a specific database.
	Available databases:
	saucenao

	Searches the SauceNAO API by default. `
}

func SearchHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	replied := ctx.EffectiveMessage.ReplyToMessage
	photos := ctx.EffectiveMessage.ReplyToMessage.Photo
	var err error = nil
	exactParameter := ctx.Args()
	if replied != nil && photos != nil {
		//Parameters: /search <count> <database>
		//database defaults to saucenao (now only have one so i won't intemplate it)
		//count defaults to 1
		//If count is not a number, it will be treated as database
		if len(exactParameter) >= 1 && len(exactParameter) <= 3 {
			var count = 1
			//Parse count
			if len(exactParameter) >= 2 {
				cnt, err := strconv.Atoi(exactParameter[1])
				if err != nil {
					_, _ = ctx.EffectiveMessage.Reply(
						bot,
						"<count> is invalid. Setting it to default value.",
						&gotgbot.SendMessageOpts{},
					)
				} else {
					count = cnt
				}
			}
			err = handlePhoto(bot, ctx, photos[len(photos)-1], count)
			return err
		} else if len(exactParameter) == 2 && exactParameter[1] == "limit" {
			err = handleLimit(bot, ctx)
		} else {
			_, err = ctx.EffectiveMessage.Reply(bot,
				"Too many parameters.",
				&gotgbot.SendMessageOpts{},
			)
		}
	} else {
		if len(exactParameter) == 2 && exactParameter[1] == "limit" {
			err = handleLimit(bot, ctx)
		} else {
			_, err = ctx.EffectiveMessage.Reply(bot,
				"In order to search for an image, please reply to an image ;)",
				&gotgbot.SendMessageOpts{})
		}
	}
	if err != nil {
		log.HandleError(err)
	}
	return nil
}

func handlePhoto(bot *gotgbot.Bot, ctx *ext.Context, photo gotgbot.PhotoSize, count int) error {
	// Download the picture
	file, err := bot.GetFile(photo.FileId, &gotgbot.GetFileOpts{})
	if err != nil {
		return err
	}

	client := http.Client{}
	imageRes, err := client.Get(file.GetURL(bot))
	if imageRes.StatusCode != 200 {
		return fmt.Errorf(
			"failed to download image %s: HTTP request failed with status %d, error is %s",
			photo.FileId, imageRes.StatusCode, err.Error(),
		)
	}

	// Search it
	cachePath := fmt.Sprintf("./cache/%s.jpg", file.FileId)
	if _, err = os.Stat("./cache"); os.IsNotExist(err) {
		os.Mkdir("./cache", os.ModePerm)
	}
	imageFile, err := os.OpenFile(cachePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	log.HandleError(err)
	defer imageFile.Close()
	io.Copy(imageFile, imageRes.Body)
	res := searchSauseNAO(cachePath, count)
	os.Remove(cachePath)

	// Respond something
	for i := 0; i < len(res.Results); i++ {
		result := &res.Results[i]
		caption := fmt.Sprintf("[%s] %s %s\n", result.Header.Similarity,
			result.Data.Title,
			result.Data.ExtUrls[0])
		_, err = bot.SendPhoto(ctx.EffectiveChat.Id, result.Header.Thumbnail, &gotgbot.SendPhotoOpts{
			Caption:          caption,
			ReplyToMessageId: ctx.EffectiveMessage.ReplyToMessage.MessageId,
		})
		if err != nil {
			log.HandleError(err)
			_, _ = ctx.EffectiveChat.SendMessage(bot,
				"<i>An error occurred when sending the photo</i>",
				&gotgbot.SendMessageOpts{
					ParseMode:        "html",
					ReplyToMessageId: ctx.EffectiveMessage.ReplyToMessage.MessageId,
				})
		}
	}

	return nil
}

func handleLimit(bot *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(bot,
		fmt.Sprintf("Account Limit:\nIn short term: %d/%d\nIn long term: %d/%s",
			saucenaoUser.ShortRemaining, saucenaoUser.ShortLimit,
			saucenaoUser.LongRemaining, saucenaoUser.LongLimit,
		),
		&gotgbot.SendMessageOpts{},
	)

	return err
}
