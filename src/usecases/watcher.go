package usecases

import (
	"TwitchNotifierForTelegram/src/domain"
	"fmt"
	"time"
)

func (i interactor) titleHasChanged(channel domain.Channel, oldTitle string, newStatus domain.ChannelStatus) string {
	err := i.titleHistoryRepo.TitleChanged(channel.ID, newStatus.Title)
	if err != nil {
		println("Error while saving title change")
		return ""
	}

	if oldTitle == "" {
		return fmt.Sprintf("- Title set to <b>%s</b>", newStatus.Title)
	} else {
		return fmt.Sprintf("- Title changed from <b>%s</b> to <b>%s</b>", oldTitle, newStatus.Title)
	}
}

func (i interactor) categoryHasChanged(channel domain.Channel, oldCategory string, newCategory domain.ChannelStatus) string {
	err := i.categoryHistoryRepo.CategoryChanged(channel.ID, newCategory.Category)
	if err != nil {
		println("Error while saving category change")
		return ""
	}

	if oldCategory == "" {
		return fmt.Sprintf("- Category set to <b>%s</b>", newCategory.Category)
	} else {
		return fmt.Sprintf("- Category changed from <b>%s</b> to <b>%s</b>", oldCategory, newCategory.Category)
	}
}

func (i interactor) statusHasChanged(channel domain.Channel, newStatus domain.ChannelStatus) {
	err := i.statusHistoryRepo.StatusChanged(channel.ID, newStatus.IsLive)
	if err != nil {
		println("Error while saving new status")
		return
	}
}

func (i interactor) StartWatcher() {
	for _ = range time.Tick(time.Second) {
		channels, err := i.channelRepo.ListAllChannels()
		if err != nil {
			return
		}

		for _, registeredchannel := range channels {
			channel, err := i.twitchAPI.GetChannel(registeredchannel.Name)
			if err == nil && channel.Status != nil {
				//check for title change
				titleChanged, oldTitle := i.checkForstreamTitleChange(channel.Name, channel.Status.Title)
				categoryChanged, oldCategory := i.checkForstreamCategoryChange(channel.Name, channel.Status.Category)
				statusChanged, _ := i.checkForstreamStatusChange(channel.Name, channel.Status.IsLive)

				if !titleChanged && !categoryChanged && !statusChanged {
					continue
				}

				followers, err := i.followRepo.GetFollowers(registeredchannel.ID)
				if err != nil {
					println("unable to get followers")
				}

				updateMessage := fmt.Sprintf("%s updates: \n", channel.DisplayName)

				if statusChanged {
					i.statusHasChanged(registeredchannel, *channel.Status)
				}

				if statusChanged && channel.Status.IsLive {
					updateMessage += "now live"
				} else {
					updateMessage += "not live no more"
				}

				if titleChanged {
					updateMessage += i.titleHasChanged(registeredchannel, oldTitle, *channel.Status)
					updateMessage += "\n"
				}

				if categoryChanged {
					updateMessage += i.categoryHasChanged(registeredchannel, oldCategory, *channel.Status)
					updateMessage += "\n"
				}

				for _, follower := range followers {
					i.SendMessage(follower, updateMessage)
				}
			}
		}
	}
}
