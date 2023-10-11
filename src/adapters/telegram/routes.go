package telegram

func (t *telegramClient) SetRoutes() {
	t.bot.Handle("/start", t.UserStarted)
	t.bot.Handle("/follow", t.FollowChannel)
}
