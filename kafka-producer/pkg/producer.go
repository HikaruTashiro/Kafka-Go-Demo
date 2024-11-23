package kafka_clients

import (
	"io"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ProducerTest struct {
	io.Closer

	topic     string
	kafkaProd *kafka.Producer
	events    chan kafka.Event
}

func NewProducerTest(topic string) (KafkaProducer, error) {
	var err error
	prod := &ProducerTest{
		topic: topic,
	}

	prod.kafkaProd, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		return nil, err
	}

	log.Println("New producer created")

	return prod, nil
}

// Expects a UTF-8 byte string with the number of elements in the array in the first position
func (p *ProducerTest) Write(data []byte) (n int, err error) {
	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topic,
			Partition: kafka.PartitionAny,
		},
		Value: data,
		Headers: []kafka.Header{
			{
				Key:   "Kafka-Test-Golang",
				Value: []byte("the very first test of prducer functionality in Go"),
			},
		},
	}
	err = p.kafkaProd.Produce(&msg, p.events)
	if err != nil {
		return 0, err
	}
	return len(data), err
}

func (p *ProducerTest) Close() error {
	p.kafkaProd.Close()
	return nil
}
