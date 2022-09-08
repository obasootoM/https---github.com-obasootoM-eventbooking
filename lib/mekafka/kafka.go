package mekafka

type Event interface {
	EventName() string
	PartionKey() string
}

type EventProducer interface {
	Produce(e Event) error
}

type Consumer interface {
	Listen(eventName ...string) (<- chan Event, <- chan error, error)
}
