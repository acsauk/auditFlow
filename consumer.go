package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func consumeEvents(ctx context.Context, r *kafka.Reader) {
	for {
		m, err := r.ReadMessage(ctx)

		if err != nil {
			if ctx.Err() != nil {
				break
			}

			log.Println(err)
			continue
		}

		log.Printf("message in topic %v: %v", m.Topic, string(m.Value))
	}
}
