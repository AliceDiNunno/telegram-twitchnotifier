package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TitleHistory struct {
	ID int64

	ChannelID uuid.UUID
	Channel   Channel

	Title string
	Date  int64
}

type titleHistoryRepo struct {
	db *gorm.DB
}

func (t titleHistoryRepo) TitleChanged(channelID uuid.UUID, title string) error {
	result := t.db.Create(&TitleHistory{
		ChannelID: channelID,
		Title:     title,
		Date:      time.Now().Unix(),
	})

	return result.Error
}

func (t titleHistoryRepo) GetChannelTitle(channelID uuid.UUID) (string, error) {
	var latestTitleKnown TitleHistory

	result := t.db.Where("channel_id = ?", channelID).Order("date DESC").First(&latestTitleKnown)

	if result.Error != nil {
		return "", result.Error
	}

	return latestTitleKnown.Title, nil
}

func NewTitleHistoryRepo(db *gorm.DB) titleHistoryRepo {
	return titleHistoryRepo{
		db: db,
	}
}
