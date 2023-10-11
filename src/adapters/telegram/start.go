package telegram

import (
	"TwitchNotifierForTelegram/src/events"
	"fmt"
	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v3"
)

func (t *telegramClient) UserStarted(c tele.Context) error {
	log.Info().Msg("User started")

	t.hub.Publish(events.UserRegisterEvent, events.EventData{
		"userID":   c.Sender().ID,
		"username": c.Sender().Username,
		"name":     fmt.Sprintf("%s %s", c.Sender().FirstName, c.Sender().LastName),
	})

	return nil
}
