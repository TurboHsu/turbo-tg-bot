package glotrunner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TurboHsu/turbo-tg-bot/utils/config"
	"github.com/TurboHsu/turbo-tg-bot/utils/log"
)

func run(code string, lang string, input string) string {
	for i := 0; i < len(glotLangs); i++ {
		//Find the language
		if glotLangs[i].Name == lang {
			//Run the code
			var reqData request
			var filename string

			//Deal with filename.
			//In Java's case, the filename should be "Main.java".
			if glotLangs[i].File == "" {
				filename = "main." + glotLangs[i].Ext
			} else {
				filename = glotLangs[i].File + "." + glotLangs[i].Ext
			}
			reqData.Files = append(reqData.Files, requestFile{Name: filename, Content: code})

			//Generate the request body.
			reqBody, err := json.Marshal(reqData)
			log.HandleError(err)
			req, err := http.NewRequest("POST", "https://glot.io/api/run/"+glotLangs[i].Name+"/latest", bytes.NewBuffer(reqBody))
			log.HandleError(err)
			req.Header.Add("Authorization", "Token "+config.Config.APIKeys.GlotAPIKey)
			req.Header.Add("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.HandleError(err)
				return "Error while requesting runner API: " + err.Error()
			}
			defer resp.Body.Close()
			result, err := io.ReadAll(resp.Body)
			log.HandleError(err)

			//Convert the result to string.
			var respData response
			err = json.Unmarshal(result, &respData)
			log.HandleError(err)
			ret := fmt.Sprintf("--------------\nStdout:\n--------------\n%s\n--------------\nStderr:\n--------------\n%s\n--------------\nError:\n--------------\n%s\n--------------",
				respData.Stdout, respData.Stderr, respData.Error)
			return ret
		}
	}
	return "Language not found."
}
