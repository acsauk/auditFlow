package main

import "github.com/segmentio/kafka-go"

const (
	KB = 1024
	MB = 1024 * KB
)

func newKafkaWriter(brokerAddress, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr: kafka.TCP(brokerAddress),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func newKafkaReader(brokerAddress, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupID,
		MinBytes: 10 * KB,
		MaxBytes: 10 * MB,
	})
}
