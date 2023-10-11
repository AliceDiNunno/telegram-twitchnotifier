package config

type Twitch struct {
	ClientID     string
	ClientSecret string
}

func (config *Config) getTwitch() {
	if config.Twitch != nil {
		return
	}

	twitch := &Twitch{
		ClientID:     config.RequireEnvString("TWITCH_CLIENT_ID"),
		ClientSecret: config.RequireEnvString("TWITCH_CLIENT_SECRET"),
	}

	config.Twitch = twitch
}
