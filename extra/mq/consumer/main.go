package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		// Consumer Groups, 不指定就是普通的一个Consumer
		GroupID: "consumer-group-id",
		// 可以指定Partition消费消息
		// Partition:      0,
		GroupBalancers: []kafka.GroupBalancer{kafka.RoundRobinGroupBalancer{}},
		Topic:          "topic-B",
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
	})

	for {
		// 读取消息后会自动提交
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		// 当然我们也可以手动处理消息的处理状态, 这在处理长耗时任务的时候非常有用
		// m, err = r.FetchMessage(context.Background())
		// if err != nil {
		// 	break
		// }
		// if err := r.CommitMessages(context.Background(), m); err != nil {
		// 	log.Fatal("failed to commit messages:", err)
		// }
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
