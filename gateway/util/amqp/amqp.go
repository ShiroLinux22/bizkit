package amqp

import (
	"fmt"
	"os"

	"github.com/chakernet/ryuko/gateway/util"
	"github.com/streadway/amqp"
)

var (
	log = util.Logger {
		Name: "amqp",
	}
)

func Connect() (*amqp.Connection) {
	uri := os.Getenv("AMQP_URI")

	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatal("Failed to connect to amqp: ", err)
	}

	log.Info("Connected to RabbitMQ v%s.%s", fmt.Sprint(conn.Major), fmt.Sprint(conn.Minor))
	return conn
}

func Channel(conn *amqp.Connection) (*amqp.Channel) {
	c, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel: ", err)
	}

	return c
}