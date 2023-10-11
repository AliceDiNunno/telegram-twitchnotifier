package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CategoryHistory struct {
	ID int64

	ChannelID uuid.UUID
	Channel   Channel

	Category string
	Date     int64
}

type categoryHistoryRepo struct {
	db *gorm.DB
}

func (c categoryHistoryRepo) CategoryChanged(channelID uuid.UUID, category string) error {
	result := c.db.Create(&CategoryHistory{
		ChannelID: channelID,
		Category:  category,
		Date:      time.Now().Unix(),
	})

	return result.Error
}

func (c categoryHistoryRepo) GetChannelCategory(channelID uuid.UUID) (string, error) {
	var latestCategoryKnown CategoryHistory

	result := c.db.Where("channel_id = ?", channelID).Order("date DESC").First(&latestCategoryKnown)

	if result.Error != nil {
		return "", result.Error
	}

	return latestCategoryKnown.Category, nil
}

func NewCategoryHistoryRepo(db *gorm.DB) categoryHistoryRepo {
	return categoryHistoryRepo{
		db: db,
	}
}
