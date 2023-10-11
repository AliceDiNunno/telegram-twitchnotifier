package hub

import (
	"TwitchNotifierForTelegram/src/events"
	"github.com/rs/zerolog/log"
	"time"
)
import "github.com/leandro-lugaresi/hub"

type Hub struct {
	hub *hub.Hub

	later map[events.Event]time.Time
}

func (h *Hub) Publish(topic events.Event, data events.EventData) {
	log.Info().Str("topic", string(topic)).Object("event", data).Msg("Event triggered")
	h.hub.Publish(hub.Message{
		Name:   string(topic),
		Fields: hub.Fields(data),
	})
}

func (h *Hub) PublishLater(topic events.Event, data events.EventData, delay time.Duration) {
	time.AfterFunc(delay, func() {
		if h.later[topic].IsZero() {
			return
		}
		h.Publish(topic, data)
	})
	h.later[topic] = time.Now().Add(delay)
}

func (h *Hub) CancelPublishLater(topic events.Event) {
	if h.later[topic].IsZero() {
		return
	}

	h.later[topic] = time.Time{}
}

func (h *Hub) Subscribe(topic events.Event, notify events.EventListener) {
	sub := h.hub.Subscribe(10, string(topic))

	go func(s hub.Subscription) {
		for msg := range s.Receiver {
			notify(topic, events.EventData(msg.Fields))
		}
	}(sub)
}

func NewHub() *Hub {
	return &Hub{
		hub:   hub.New(),
		later: map[events.Event]time.Time{},
	}
}
