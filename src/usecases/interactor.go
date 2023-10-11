package usecases

import (
	"TwitchNotifierForTelegram/src/adapters/events/hub"
	"TwitchNotifierForTelegram/src/domain"
	"github.com/google/uuid"
)

type TelegramBot interface {
	SendMessage(chatID int64, text string) error
}

type UserRepo interface {
	RegisterUser(userID int64, username string, displayName string) error
	UpdateUser(userID int64, username string, displayName string) error
	GetUser(userID int64) (domain.User, error)
}

type ChannelRepo interface {
	CreateChannel(channel domain.Channel) error
	GetChannel(channelName string) (domain.Channel, error)
	ListAllChannels() ([]domain.Channel, error)
}

type FollowedChannelRepo interface {
	FollowChannel(userID int64, channelID uuid.UUID) error
	UnfollowChannel(userID int64, channelID uuid.UUID) error
	GetFollowedChannels(userID int64) ([]uuid.UUID, error)
	IsFollowing(userID int64, channelID uuid.UUID) bool
	GetFollowers(channelID uuid.UUID) ([]int64, error)
}

type TitleHistoryRepo interface {
	TitleChanged(channelID uuid.UUID, title string) error
	GetChannelTitle(channelID uuid.UUID) (string, error)
}

type CategoryHistoryRepo interface {
	CategoryChanged(channelID uuid.UUID, title string) error
	GetChannelCategory(channelID uuid.UUID) (string, error)
}

type StatusHistoryRepo interface {
	StatusChanged(channelID uuid.UUID, isLive bool) error
	GetChannelStatus(channelID uuid.UUID) (bool, error)
}

type StreamHistoryRepo interface {
	StreamStarted(channelName string) error
	StreamEnded(channelName string) error
}

type TwitchAPI interface {
	GetChannel(channelName string) (domain.Channel, error)
}

type interactor struct {
	userRepo            UserRepo
	channelRepo         ChannelRepo
	followRepo          FollowedChannelRepo
	titleHistoryRepo    TitleHistoryRepo
	categoryHistoryRepo CategoryHistoryRepo
	statusHistoryRepo   StatusHistoryRepo

	twitchAPI   TwitchAPI
	telegramBot TelegramBot
	hub         *hub.Hub
}

func NewInteractor(uR UserRepo, cR ChannelRepo, fR FollowedChannelRepo,
	sHR StatusHistoryRepo, thR TitleHistoryRepo, chR CategoryHistoryRepo,
	tA TwitchAPI, tB TelegramBot, h *hub.Hub) Usecases {
	return interactor{
		userRepo:    uR,
		channelRepo: cR,
		followRepo:  fR,

		statusHistoryRepo:   sHR,
		titleHistoryRepo:    thR,
		categoryHistoryRepo: chR,

		twitchAPI:   tA,
		telegramBot: tB,
		hub:         h,
	}
}
