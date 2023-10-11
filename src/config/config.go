package config

import (
	"TwitchNotifierForTelegram/src/domain/types"
)

// Config is a struct that holds the configuration of the service
type Config struct {
	Version  types.Version
	Database *Database
	Twitch   *Twitch
	Telegram *Telegram
}

func NewConfig() *Config {
	conf := Config{}

	conf.getVersion()
	conf.getDatabase()
	conf.getTwitch()
	conf.getTelegram()

	return &conf
}
