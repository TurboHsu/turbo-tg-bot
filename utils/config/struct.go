package config

type configStruct struct {
	Common struct {
		WriteLog          bool
		Debug             bool // If debug is enabled, all infos and errors will be logged into the file
		LogPath           string
		Silent            bool // Silent mode will not print anything to the console
		ProxyURL          string
		DropPendingUpdate bool
	}

	APIKeys struct {
		BotToken       string
		SaucenaoAPIKey string
		Debug          bool
		Silent         bool
		GlotAPIKey     string
	}
}
