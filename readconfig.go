package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func readConfig() {
	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()
	configRaw, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(configRaw), &config)
}
