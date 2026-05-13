package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	brokerAddress := "localhost:9092"
	topic := "audit-events"

	kafkaWriter := newKafkaWriter(brokerAddress, topic)
	defer func() {
		if err := kafkaWriter.Close(); err != nil {
			log.Printf("error closing kafka writer: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	kafkaReader := newKafkaReader(brokerAddress, topic, "audit-consumer-group")
	defer func() {
		if err := kafkaReader.Close(); err != nil {
			log.Printf("error closing kafka reader: %v", err)
		}
	}()

	go consumeEvents(ctx, kafkaReader)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", handlePing)
	mux.HandleFunc("POST /events", handleCreateEvent(kafkaWriter))

	log.Println("Listening on port 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
