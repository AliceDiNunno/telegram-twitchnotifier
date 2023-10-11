package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type StatusHistory struct {
	ID int64

	ChannelID uuid.UUID
	Channel   Channel

	IsLive bool
	Date   int64
}

type statusHistoryRepo struct {
	db *gorm.DB
}

func (s statusHistoryRepo) StatusChanged(channelID uuid.UUID, isLive bool) error {
	result := s.db.Create(&StatusHistory{
		ChannelID: channelID,
		IsLive:    isLive,
		Date:      time.Now().Unix(),
	})

	return result.Error
}

func (s statusHistoryRepo) GetChannelStatus(channelID uuid.UUID) (bool, error) {
	var latestStatusKnown StatusHistory

	result := s.db.Where("channel_id = ?", channelID).Order("date DESC").First(&latestStatusKnown)

	if result.Error != nil {
		return false, result.Error
	}

	return latestStatusKnown.IsLive, nil
}

func NewStatusHistoryRepo(db *gorm.DB) statusHistoryRepo {
	return statusHistoryRepo{
		db: db,
	}
}
