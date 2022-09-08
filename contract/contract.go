package contract

import "time"

type EventCreateEvent struct {
	Id         string      `json:"id"`
	Name       string    `json:"name"`
	LocationId string    `json:"location-id"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
}

func (e EventCreateEvent) PartionKey() string {
	return e.Id
}
func (e EventCreateEvent) EventName() string {
	return "event-created"
}
