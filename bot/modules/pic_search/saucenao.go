package picsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/TurboHsu/turbo-tg-bot/utils/config"
	"github.com/TurboHsu/turbo-tg-bot/utils/log"
)

var saucenaoUser saucenaoUserStat

// This function searches saucenao
func searchSauseNAO(filepath string, count int) saucenaoResponse {
	//Search the pic
	photo, err := os.Open(filepath)
	log.HandleError(err)
	defer photo.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath)
	log.HandleError(err)
	_, err = io.Copy(part, photo)
	log.HandleError(err)
	writer.Close()

	httpReq, err := http.NewRequest("POST",
		fmt.Sprintf("https://saucenao.com/search.php?output_type=2&numres=%d&hide=0&db=999&api_key=%s",
			count, config.Config.APIKeys.SaucenaoAPIKey),
		body)
	log.HandleError(err)
	httpReq.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	reqDo, err := client.Do(httpReq)
	log.HandleError(err)
	response, err := io.ReadAll(reqDo.Body)
	log.HandleError(err)

	var ret saucenaoResponse
	json.Unmarshal(response, &ret)

	//Update user limitation
	saucenaoUser.LongLimit = ret.Header.LongLimit
	saucenaoUser.ShortLimit = ret.Header.ShortRemaining
	saucenaoUser.LongRemaining = ret.Header.LongRemaining
	saucenaoUser.ShortRemaining = ret.Header.ShortRemaining

	//Return the result
	return ret
}
