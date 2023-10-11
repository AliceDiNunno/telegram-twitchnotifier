package events

import "github.com/rs/zerolog"

type Event string
type EventData map[string]interface{}

func (e2 EventData) MarshalZerologObject(e *zerolog.Event) {
	for k, v := range e2 {
		e.Interface(k, v)
	}
}

type EventListener func(event Event, data EventData)

const UserRegisterEvent Event = "user.register"
const UserFollowChannelEvent Event = "user.follow"
