package postgres

import (
	"TwitchNotifierForTelegram/src/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type Channel struct {
	ID uuid.UUID `gorm:"primaryKey"`

	Name        string
	DisplayName string
}

type channelRepo struct {
	db *gorm.DB
}

func (c channelRepo) GetChannel(channelName string) (domain.Channel, error) {
	var channel Channel

	result := c.db.First(&channel, "name = ?", channelName)

	if result.Error != nil {
		return domain.Channel{}, result.Error
	}

	return domain.Channel{
		ID:          channel.ID,
		Name:        strings.ToLower(channel.Name),
		DisplayName: channel.DisplayName,
	}, nil
}

func (c channelRepo) CreateChannel(channel domain.Channel) error {
	result := c.db.Create(&Channel{
		ID:          channel.ID,
		Name:        strings.ToLower(channel.Name),
		DisplayName: channel.DisplayName,
	})

	return result.Error
}

func (c channelRepo) ListAllChannels() ([]domain.Channel, error) {
	var channels []Channel

	result := c.db.Find(&channels)

	if result.Error != nil {
		return nil, result.Error
	}

	var domainChannels []domain.Channel

	for _, channel := range channels {
		domainChannels = append(domainChannels, domain.Channel{
			ID:          channel.ID,
			Name:        strings.ToLower(channel.Name),
			DisplayName: channel.DisplayName,
		})
	}

	return domainChannels, nil
}

func NewChannelRepo(db *gorm.DB) channelRepo {
	return channelRepo{
		db: db,
	}
}
