package listener

import (
	"encoding/json"
	"eventsbook/contract"
	"eventsbook/lib/kafka"
	"eventsbook/lib/mekafka"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/mitchellh/mapstructure"
)

type KafkaListener struct {
	consumer  sarama.Consumer
	partition []int32
}

func NewKafkerListener(client sarama.Client, partition []int32) (mekafka.Consumer, error) {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}
	listner := &KafkaListener{
		consumer:  consumer,
		partition: partition,
	}
	return listner, nil
}

func (k *KafkaListener) Listen(eventName ...string) (<-chan mekafka.Event, <-chan error, error) {
	var err error
	topic := "event"
	result := make(chan mekafka.Event)
	errors := make(chan error)
	partition := k.partition
	if len(partition) == 0 {
		partition, err = k.consumer.Partitions(topic)
		if err != nil {
			return nil, nil, err
		}
	}
	log.Printf("topic %s has partition %v", topic, partition)
	for _, partitions := range partition {
		con, err := k.consumer.ConsumePartition(topic, partitions, 0)
		if err != nil {
			return nil, nil, err
		}
		go func() {
			for msg := range con.Messages() {
				body := kafka.MessageEnvelope{}
				err := json.Unmarshal(msg.Value, &body)
				if err != nil {
					errors <- fmt.Errorf("cannot %s unmarshal from json", err)
					continue
				}
				var event mekafka.Event
				switch body.EventName {
				case "event-created":
					event = contract.EventCreateEvent{}
				default:
					errors <- fmt.Errorf("unknown event type %s", body.EventName)
				}

				mgs := mapstructure.DecoderConfig{
					Result:  event,
					TagName: "json",
				}
				_, err = mapstructure.NewDecoder(&mgs)
				if err != nil {
					errors <- fmt.Errorf("could not map event %s, %s", body.EventName, err)
				}
				result <- event
			}
		}()
	}
	return result, errors, nil
}
