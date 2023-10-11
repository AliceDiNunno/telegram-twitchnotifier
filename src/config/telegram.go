package config

type Telegram struct {
	BotToken string
}

func (config *Config) getTelegram() {
	if config.Telegram != nil {
		return
	}

	telegram := &Telegram{
		BotToken: config.RequireEnvString("TELEGRAM_BOT_TOKEN"),
	}

	config.Telegram = telegram
}
