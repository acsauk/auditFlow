package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlePing(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}

func handleCreateEvent(w http.ResponseWriter, r *http.Request) {
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
	
	w.WriteHeader(http.StatusAccepted)
}
