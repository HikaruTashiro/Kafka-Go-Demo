package kafka_clients

import (
	"bytes"
	"encoding/binary"
	"log"
)

type KafkaEncoder struct {
	builder *bytes.Buffer
}

func NewKafkaEncoder() Encoder {
	return &KafkaEncoder{
		builder: bytes.NewBuffer(make([]byte, 0)),
	}
}

func (k *KafkaEncoder) Encode(data string) []byte {
	k.builder.Reset()
	lenght := len(data)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(lenght))
	k.builder.Write(bs)
	k.builder.WriteString(data)
	return k.builder.Bytes()
}

func (k *KafkaEncoder) EncodeBytes(data []byte) []byte {
	k.builder.Reset()
	lenght := uint32(len(data))
	log.Printf("[Encoding]: lenght = %d\n", lenght)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, lenght)
	k.builder.Write(bs)
	k.builder.Write(data)
	return k.builder.Bytes()
}
