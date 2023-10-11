package twitch

import (
	"TwitchNotifierForTelegram/src/adapters/events/hub"
	"TwitchNotifierForTelegram/src/config"
	"TwitchNotifierForTelegram/src/domain"
	"errors"
	"github.com/nicklaw5/helix/v2"
	"strings"
)

type TwitchClient struct {
	hub    *hub.Hub
	client *helix.Client
}

func (t TwitchClient) GetChannel(channelName string) (domain.Channel, error) {
	resp, err := t.client.SearchChannels(&helix.SearchChannelsParams{
		Channel:  channelName,
		First:    1,
		LiveOnly: false,
	})

	if err != nil {
		return domain.Channel{}, err
	}

	if len(resp.Data.Channels) == 0 {
		return domain.Channel{}, errors.New("channel not found")
	}

	if strings.ToLower(resp.Data.Channels[0].BroadcasterLogin) != strings.ToLower(channelName) {
		return domain.Channel{}, errors.New("channel not found")
	}

	channelStatus := domain.ChannelStatus{
		IsLive:   resp.Data.Channels[0].IsLive,
		Title:    resp.Data.Channels[0].Title,
		Category: resp.Data.Channels[0].GameName,
	}

	return domain.Channel{
		Name:        resp.Data.Channels[0].BroadcasterLogin,
		DisplayName: resp.Data.Channels[0].DisplayName,
		Status:      &channelStatus,
	}, nil
}

func NewTwitchClient(cfg *config.Twitch, hub *hub.Hub) *TwitchClient {
	if cfg == nil {
		panic("Twitch config is nil")
	}

	twitchClient, err := helix.NewClient(&helix.Options{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
	})

	if err != nil {
		panic(err)
	}

	token, err := twitchClient.RequestAppAccessToken([]string{})
	if err != nil {
		panic(err)
	}
	twitchClient.SetAppAccessToken(token.Data.AccessToken)

	return &TwitchClient{
		hub:    hub,
		client: twitchClient,
	}
}
