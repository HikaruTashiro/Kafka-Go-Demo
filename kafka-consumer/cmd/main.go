package main

import (
	"encoding/json"
	"kafka/pkg"
	"log"
)

func main() {
	consumer, err := kafka_clients.NewConsumerTest("golang-test")
	if err != nil {
		log.Println("error: ", err)
		return
	}
	defer consumer.Close()

	for {
		data := make([]byte, 2048)
		n, err := consumer.Read(data)
		if err != nil {
			log.Println("something went wrong while trying to read consumer data", err)
			continue
		}
		if n == 0 {
			log.Println("nothing was read, wait for futher messages to be sent")
			continue
		}
		decoder := kafka_clients.NewKafkaDecoder()
		value, err := decoder.Decode(data)
		if err != nil {
			log.Println("something went wrong at the decoder", err)
			continue
		}
		log.Printf("value received %s, with len %d\n", value, len(value))

		based := kafka_clients.InvoicePayerInfo{}
		err = json.Unmarshal([]byte(value), &based)
		if err != nil {
			log.Println("json provided is not of the type accepted by the application:", err)
			continue
		}

		log.Printf("based is %v\n", based)
	}
}
