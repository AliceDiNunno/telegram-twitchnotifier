package config

import (
	"TwitchNotifierForTelegram/src/domain/types"
)

func (config *Config) getVersion() {
	if config.Version != "" {
		return
	}

	envVar := config.GetEnvStringOrDefault("GITHUB_SHA", "unknown")
	config.Version = types.Version(envVar)
}
