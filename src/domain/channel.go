package domain

import "github.com/google/uuid"

type ChannelStatus struct {
	IsLive   bool
	Title    string
	Category string
}

type Channel struct {
	ID          uuid.UUID
	Name        string
	DisplayName string

	Status *ChannelStatus
}

func (c *Channel) DispName() string {
	if c.DisplayName != "" {
		return c.DisplayName
	}

	return c.Name
}

func (c *Channel) Init() {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
}
