package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
		// NOTE: When Topic is not defined here, each Message must define it instead.
		Topic:    "topic-B",
		Balancer: &kafka.LeastBytes{},
		// The topic will be created if it is missing.
		AllowAutoTopicCreation: true,
		// 支持消息压缩
		// Compression: kafka.Snappy,
		// 支持TLS
		// Transport: &kafka.Transport{
		//     TLS: &tls.Config{},
		// }
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			// 支持 Writing to multiple topics
			//  NOTE: Each Message has Topic defined, otherwise an error is returned.
			// Topic: "topic-A",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
		kafka.Message{
			Key:   []byte("Key-D"),
			Value: []byte("Thr!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
