package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	kafkaWriter := newKafkaWriter("localhost:9092", "audit-events")
	defer func() {
		if err := kafkaWriter.Close(); err != nil {
			log.Printf("error closing kafka writer: %v", err)
		}
	}()

	mux.HandleFunc("GET /ping", handlePing)
	mux.HandleFunc("POST /events", handleCreateEvent(kafkaWriter))

	log.Println("Listening on port 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
