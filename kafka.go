package main

import "github.com/segmentio/kafka-go"

func newKafkaWriter(brokerAddress, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr: kafka.TCP(brokerAddress),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
	}
}
