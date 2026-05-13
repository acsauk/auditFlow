package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/segmentio/kafka-go"
)

func handlePing(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "pong"}); err != nil {
		http.Error(w, "error encoding body", http.StatusBadRequest)
		return
	}
}

func handleCreateEvent(writer *kafka.Writer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if !event.Valid() {
			http.Error(w, "event missing required fields (ID, Source, Action and Actor)", http.StatusBadRequest)
			return
		}

		log.Println(event)

		eventBytes, err := json.Marshal(event)
		if err != nil {
			http.Error(w, fmt.Sprintf("unable to marshal event to JSON byte: %v", err), http.StatusInternalServerError)
			return
		}

		if err := writer.WriteMessages(r.Context(), kafka.Message{
			Key:   []byte(event.ID),
			Value: eventBytes,
		}); err != nil {
			http.Error(w, fmt.Sprintf("error writing message to queue: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
