package initialize

import (
	"log"

	"github.com/segmentio/kafka-go"

	"Go/global"
)

// Init kafka Producer

var KafkaProducer * kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr: kafka.TCP("localhost:19092"),
		Topic: "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}


}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Failed to close kafka producer: %v", err)
	}
}