package usecases

import (
	"TwitchNotifierForTelegram/src/domain"
	"TwitchNotifierForTelegram/src/events"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

func (i interactor) FollowChannel(_ events.Event, data events.EventData) {
	userID := data["userID"].(int64)
	channelName := data["channelName"].(string)

	channel, err := i.channelRepo.GetChannel(strings.ToLower(channelName))

	if err != nil {
		log.Warn().Err(err).Str("channelName", channelName).Msg("Channel not in records, checking twitch for channel")

		twitchChannel, err := i.twitchAPI.GetChannel(channelName)
		if err != nil {
			i.SendMessage(userID, fmt.Sprintf("The channel <b>%s</b> was not found", channelName))
			log.Warn().Err(err).Str("channelName", channelName).Msg("Channel not found on twitch")
			return
		}

		log.Info().Str("channel_name", twitchChannel.DispName()).Msg("Channel found on twitch, adding to records")
		channel = domain.Channel{
			Name:        twitchChannel.Name,
			DisplayName: twitchChannel.DisplayName,
		}
		channel.Init()
		err = i.channelRepo.CreateChannel(channel)

		if err != nil {
			i.SendMessage(userID, fmt.Sprintf("An error happened while following <b>%s</b>", channelName))
			log.Warn().Err(err).Str("channelName", channelName).Msg("Unable to follow channel")
			return
		}
	}

	if i.followRepo.IsFollowing(userID, channel.ID) {
		i.SendMessage(userID, fmt.Sprintf("You are already following <b>%s</b>", channel.DispName()))
		return
	}

	err = i.followRepo.FollowChannel(userID, channel.ID)
	if err != nil {
		i.SendMessage(userID, fmt.Sprintf("An error happened while following <b>%s</b>", channelName))
		return
	}

	i.SendMessage(userID, fmt.Sprintf("You are now following <b>%s</b>", channel.DispName()))

	_ = channel

	return

}
