package kafka_clients

import "io"

type (
	KafkaConsumer interface {
		io.ReadCloser
	}
	KafkaProducer interface {
		io.WriteCloser
	}
	Encoder interface {
		Encode(data string) []byte
		EncodeBytes(data []byte) []byte
	}
	Decoder interface {
		Decode(data []byte) (string, error)
	}
)
