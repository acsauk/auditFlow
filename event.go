package main

import "time"

type Event struct {
	ID        string            `json:"id"`
	Source    string            `json:"source"`
	Action    string            `json:"action"`
	Actor     string            `json:"actor"`
	Timestamp time.Time         `json:"timestamp"`
	Metadata  map[string]string `json:"metadata"`
}
