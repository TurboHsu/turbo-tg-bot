package main

import (
	"flag"

	"github.com/TurboHsu/turbo-tg-bot/bot"
	"github.com/TurboHsu/turbo-tg-bot/utils/config"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.toml", "config file path")
	flag.Parse()

	config.Init(configPath)
	bot.InitBot()
}
