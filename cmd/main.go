package main

import (
	"TwitchNotifierForTelegram/src/adapters/environment"
	"TwitchNotifierForTelegram/src/adapters/events/hub"
	"TwitchNotifierForTelegram/src/adapters/log/zerolog"
	"TwitchNotifierForTelegram/src/adapters/telegram"
	"TwitchNotifierForTelegram/src/adapters/twitch"
	"TwitchNotifierForTelegram/src/config"
	"TwitchNotifierForTelegram/src/database/postgres"
	"TwitchNotifierForTelegram/src/usecases"
	"github.com/joho/godotenv"
	logger "github.com/rs/zerolog/log"
)

func main() {
	_ = godotenv.Load()

	zerolog.NewLogger(environment.GetEnvironment())
	cfg := config.NewConfig()

	db := postgres.StartGormDatabase(cfg.Database)

	err := db.AutoMigrate(
		&postgres.User{},
		&postgres.Channel{},
		&postgres.StatusHistory{},
		&postgres.CategoryHistory{},
		&postgres.TitleHistory{},
		&postgres.Follow{})
	if err != nil {
		logger.Err(err).Msg("Unable to migrate database")
	}

	userRepo := postgres.NewUserRepo(db)
	channelRepo := postgres.NewChannelRepo(db)
	followRepo := postgres.NewFollowRepo(db)

	titleHistoryRepo := postgres.NewTitleHistoryRepo(db)
	categoryHistoryRepo := postgres.NewCategoryHistoryRepo(db)
	statusHistoryRepo := postgres.NewStatusHistoryRepo(db)

	var eventHub = hub.NewHub()

	twitchClient := twitch.NewTwitchClient(cfg.Twitch, eventHub)

	telebot := telegram.NewTelegramClient(cfg.Telegram, eventHub)

	uc := usecases.NewInteractor(userRepo, channelRepo, followRepo, statusHistoryRepo, titleHistoryRepo, categoryHistoryRepo, twitchClient, telebot, eventHub)
	uc.RegisterForEvents()
	go uc.StartWatcher()

	telebot.Start()
}
