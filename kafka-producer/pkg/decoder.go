package kafka_clients

import (
	"encoding/binary"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

type KafkaDecoder struct {
	builder strings.Builder
}

func NewKafkaDecoder() Decoder {
	return &KafkaDecoder{
		builder: strings.Builder{},
	}
}

func (k *KafkaDecoder) Decode(data []byte) (string, error) {
	k.builder.Reset()
	lenght := binary.BigEndian.Uint32(data[:4])
	// NOTE: no checking the out of boundary data
	k.builder.Write(data[4 : lenght+4])
	log.Printf("[Deconding] lenght = %d\n", lenght)
	if !utf8.Valid(data[4 : lenght+4]) {
		return "", fmt.Errorf("received invalid UTF-8 string")
	}
	return k.builder.String(), nil
}
