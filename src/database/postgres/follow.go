package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Follow struct {
	ID uuid.UUID

	UserID int64
	User   User

	ChannelID uuid.UUID
	Channel   Channel
}

type followRepo struct {
	db *gorm.DB
}

func (f followRepo) IsFollowing(userID int64, channelID uuid.UUID) bool {
	var follow Follow

	result := f.db.First(&follow, "user_id = ? AND channel_id = ?", userID, channelID)

	return result.Error == nil
}

func (f followRepo) FollowChannel(userID int64, channelID uuid.UUID) error {
	//add channel to followed channels
	result := f.db.Create(&Follow{
		ID:        uuid.New(),
		UserID:    userID,
		ChannelID: channelID,
	})

	return result.Error
}

func (f followRepo) UnfollowChannel(userID int64, channelID uuid.UUID) error {
	result := f.db.Delete(&Follow{}, "user_id = ? AND channel_id = ?", userID, channelID)

	return result.Error
}

func (f followRepo) GetFollowedChannels(userID int64) ([]uuid.UUID, error) {
	var follows []Follow

	result := f.db.Find(&follows, "user_id = ?", userID)

	if result.Error != nil {
		return nil, result.Error
	}

	var channelIDs []uuid.UUID

	for _, follow := range follows {
		channelIDs = append(channelIDs, follow.ChannelID)
	}

	return channelIDs, nil
}

func (f followRepo) GetFollowers(channelID uuid.UUID) ([]int64, error) {
	var follows []Follow

	result := f.db.Find(&follows, "channel_id = ?", channelID)

	if result.Error != nil {
		return nil, result.Error
	}

	var userIDs []int64

	for _, follow := range follows {
		userIDs = append(userIDs, follow.UserID)
	}

	return userIDs, nil
}

func NewFollowRepo(db *gorm.DB) followRepo {
	return followRepo{
		db: db,
	}
}
