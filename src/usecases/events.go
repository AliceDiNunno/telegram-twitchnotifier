package usecases

import (
	"TwitchNotifierForTelegram/src/events"
	"github.com/rs/zerolog/log"
)

func (i interactor) RegisterForEvents() {
	log.Info().Msg("Registering for events")

	i.hub.Subscribe(events.UserRegisterEvent, i.RegisterUser)

	i.hub.Subscribe(events.UserFollowChannelEvent, i.FollowChannel)
}
