package usecases

import (
	"TwitchNotifierForTelegram/src/events"
	logger "github.com/rs/zerolog/log"
)

func (i interactor) welcomeMessage(userID int64) {
	i.SendMessage(userID, "Welcome to Twitch Notifier for Telegram !")
}

func (i interactor) RegisterUser(_ events.Event, data events.EventData) {
	userID := data["userID"].(int64)
	userName := data["username"].(string)
	name := data["name"].(string)

	user, err := i.userRepo.GetUser(userID)
	if err != nil {
		logger.Info().Int64("userID", userID).Msg("User not found, registering user")
		err = i.userRepo.RegisterUser(userID, userName, name)
		if err != nil {
			i.SendMessage(userID, "An error happened while registering please try again")
			return
		}
	} else {
		if user.Username != userName || user.Displayname != name {
			logger.Info().Int64("userID", userID).Msg("User not up to date, updating")
			err := i.userRepo.UpdateUser(userID, userName, name)
			if err != nil {
				i.SendMessage(userID, "An error happened while updating please try again")
				return
			}
		}
	}
	logger.Info().Int64("userID", userID).Msg("Welcoming user")
	i.welcomeMessage(userID)
}
