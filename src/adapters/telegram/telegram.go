package telegram

import (
	"TwitchNotifierForTelegram/src/adapters/events/hub"
	"TwitchNotifierForTelegram/src/config"
	tele "gopkg.in/telebot.v3"
	"time"
)

type telegramClient struct {
	bot *tele.Bot

	hub *hub.Hub
}

func NewTelegramClient(cfg *config.Telegram, hub *hub.Hub) *telegramClient {
	if cfg == nil {
		panic("Telegram config is nil")
	}

	pref := tele.Settings{
		Token:  cfg.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	tele := &telegramClient{
		bot: b,
		hub: hub,
	}

	tele.SetRoutes()

	return tele
}

func (t *telegramClient) Start() {
	t.bot.Start()
}
