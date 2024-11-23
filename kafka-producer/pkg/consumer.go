package kafka_clients

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ConsumerTest struct {
	consumer *kafka.Consumer
}

func NewConsumerTest(topic string) (KafkaConsumer, error) {
	var err error
	if topic == "" {
		return nil, fmt.Errorf("no topic was provided")
	}

	c := &ConsumerTest{}
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"group.id":           "ConsumerTest",
		"session.timeout.ms": 6000,
		// Start reading from the first message of each assigned
		// partition if there are no previously committed offsets
		// for this group.
		"auto.offset.reset": "latest",
		// Whether or not we store offsets automatically.
		"enable.auto.offset.store": false,
	}

	c.consumer, err = kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	log.Println("Consumer succefully created")
	err = c.consumer.Subscribe(topic, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("succefully subscribed to %s topic\n", topic)

	return c, nil
}

// when this return 0 with nil assume that no message was sent and nothing of wrong occured
func (c *ConsumerTest) Read(data []byte) (n int, err error) {
	ev := c.consumer.Poll(1000)
	if ev == nil {
		return 0, nil
	}

	switch e := ev.(type) {
	case *kafka.Message:
		log.Printf("Message received was %s", string(e.Value))
		n = copy(data, e.Value)
	case kafka.Error:
		log.Println("error was received whist reading from kafka consumer")
		return int(e.Code()), e
	}
	return
}

func (c *ConsumerTest) Close() error {
	return c.consumer.Close()
}
