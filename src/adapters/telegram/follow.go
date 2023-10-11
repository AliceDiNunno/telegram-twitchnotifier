package telegram

import (
	"TwitchNotifierForTelegram/src/events"
	tele "gopkg.in/telebot.v3"
)

func (t *telegramClient) FollowChannel(c tele.Context) error {
	if len(c.Args()) == 0 {
		_, err := t.bot.Send(c.Sender(), "Please, provide a channel name")
		return err
	}

	t.hub.Publish(events.UserFollowChannelEvent, map[string]interface{}{
		"userID":      c.Sender().ID,
		"channelName": c.Args()[0],
	})
	return nil
}
