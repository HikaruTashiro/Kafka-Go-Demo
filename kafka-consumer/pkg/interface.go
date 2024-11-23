package kafka_clients

import "io"

type (
	KafkaConsumer interface {
		io.ReadCloser
	}
	KafkaProducer interface {
		io.WriteCloser
	}
	// NOTE: encoder/decoder are not stricly necessary since you are not obligated to follow a
	//      encoding pattern for kafka to send and receive messages
	Encoder interface {
		Encode(data string) []byte
	}
	Decoder interface {
		Decode(data []byte) (string, error)
	}
)
