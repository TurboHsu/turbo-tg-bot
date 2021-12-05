package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func glotRun(code string, language string, input string) string {
	for i := 0; i < len(config.Glot); i++ {
		if config.Glot[i].Name == language {
			var request glotRequest
			var filename string
			request.Stdin = input

			//Summon the filename, for java needs Main.java as filename.
			if len(config.Glot[i].File) == 0 {
				filename = "main." + config.Glot[i].Ext
			} else {
				filename = config.Glot[i].File + "." + config.Glot[i].Ext
			}

			request.Files = append(request.Files, glotRequestFile{Name: filename, Content: code})

			jsonraw, _ := json.Marshal(request)
			//Send the request to glot.io
			response, _ := http.NewRequest("POST", "https://glot.io/api/run/"+config.Glot[i].Name+"/latest", bytes.NewBuffer(jsonraw))
			response.Header.Add("Authorization", "Token "+config.GlotAPIKey)
			response.Header.Add("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(response)
			if err != nil {
				return "Request error."
			}
			defer resp.Body.Close()
			result, _ := ioutil.ReadAll(resp.Body)

			//Convert the result
			var final glotResponse
			json.Unmarshal(result, &final)
			ret := `Stdout:
			` + final.Stdout + `
			Stderr: ` + final.Stderr + `
			Error: ` + final.Error

			return ret
		}
	}
	return "Language not supported."
}
