package kafka_clients

import (
	"fmt"
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
		"bootstrap.servers": "192.168.122.192:9092",
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
	e := <-p.kafkaProd.Events()
	switch m := e.(type) {
	case *kafka.Message:
		if m.TopicPartition.Error != nil {
			return 0, fmt.Errorf("Delivery Failed: %v", m.TopicPartition.Error)
		} else {
			log.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	case kafka.Error:
		return 0, m
	}
	return len(data), err
}

func (p *ProducerTest) Close() error {
	p.kafkaProd.Close()
	return nil
}
