package telegram

import tele "gopkg.in/telebot.v3"

func (t *telegramClient) SendMessage(chatID int64, text string) error {
	_, err := t.bot.Send(&tele.User{ID: chatID}, text, &tele.SendOptions{
		ParseMode: "HTML",
	})
	return err
}
