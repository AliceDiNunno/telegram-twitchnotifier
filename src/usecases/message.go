package usecases

func (i interactor) SendMessage(userID int64, message string) {
	i.telegramBot.SendMessage(userID, message)
}
