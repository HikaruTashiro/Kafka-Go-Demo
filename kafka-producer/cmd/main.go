package main

import (
	"bufio"
	"kafka/pkg"
	"log"
	"os"
)

func main() {
	producer, err := kafka_clients.NewProducerTest("golang-test")
	if err != nil {
		log.Println("error: ", err)
		return
	}
	defer producer.Close()
	input := bufio.NewReader(os.Stdin)

	for {
		var kafkaData []byte
		log.Println("provide any string ye cunt: ")
		kafkaData, _, err := input.ReadLine()
		if err != nil {
			log.Println("[error]: could not read string")
			continue
		}
		encoder := kafka_clients.NewKafkaEncoder()
		data := encoder.EncodeBytes(kafkaData)
		n, err := producer.Write([]byte(data))
		if err != nil {
			log.Println("unable to write to Kafka: ", err)
			continue
		}
		log.Printf("%d bytes where written for %s\n", n, kafkaData)
		input.Reset(os.Stdin)
	}
}
