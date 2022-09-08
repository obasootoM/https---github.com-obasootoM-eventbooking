package kafka

import (
	"encoding/json"
	"eventsbook/lib/mekafka"

	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

type MessageEnvelope struct {
	EventName string `json:"event-name"`
	Payload   interface{} `json:"pay-load"`
}

func NewProducer(client sarama.Client) (mekafka.EventProducer, error) {
	produce, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}
	kafkaProduce := &Producer{
		producer: produce,
	}
	return kafkaProduce, nil
}

func (p *Producer) Produce(event mekafka.Event) error {
	envelope := MessageEnvelope{
      event.EventName(),
	  event,
	}
	jsonBody, err := json.Marshal(&envelope)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: event.EventName(),
		Value: sarama.ByteEncoder(jsonBody),
	}
	_, _, err = p.producer.SendMessage(msg)

	return err
}
