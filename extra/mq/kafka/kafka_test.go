package kafka_test

import (
	"net"
	"strconv"
	"testing"

	"github.com/segmentio/kafka-go"
)

func TestCreateTopic(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	err = controllerConn.CreateTopics(kafka.TopicConfig{Topic: "topic-B", NumPartitions: 6, ReplicationFactor: 1})

	if err != nil {
		t.Fatal(err)
	}
}
