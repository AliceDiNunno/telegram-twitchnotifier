package usecases

func (i interactor) streamStarted(channelName string) bool {
	return false
}

func (i interactor) streamEnded(channelName string) bool {
	return false
}

func (i interactor) checkForstreamTitleChange(channelName string, currentTitle string) (bool, string) {
	channelID, err := i.channelRepo.GetChannel(channelName)
	if err != nil {
		return false, ""
	}

	oldTitle, _ := i.titleHistoryRepo.GetChannelTitle(channelID.ID)

	if oldTitle != currentTitle {
		return true, oldTitle
	}

	return false, ""
}

func (i interactor) checkForstreamCategoryChange(channelName string, currentCategory string) (bool, string) {
	channelID, err := i.channelRepo.GetChannel(channelName)
	if err != nil {
		return false, ""
	}

	oldCategory, _ := i.categoryHistoryRepo.GetChannelCategory(channelID.ID)

	if oldCategory != currentCategory {
		return true, oldCategory
	}

	return false, ""
}

func (i interactor) checkForstreamStatusChange(channelName string, currentStatus bool) (bool, bool) {
	channelID, err := i.channelRepo.GetChannel(channelName)
	if err != nil {
		return false, currentStatus
	}

	oldStatus, _ := i.statusHistoryRepo.GetChannelStatus(channelID.ID)

	if oldStatus != currentStatus {
		return true, oldStatus
	}

	return false, currentStatus
}

func (i interactor) streamCategoryChanged(channelName string) bool {
	return false
}
