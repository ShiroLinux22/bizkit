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
	log.FatalOnError(err, "Failed to connect to amqp")

	log.Info("Connected to RabbitMQ v%s.%s", fmt.Sprint(conn.Major), fmt.Sprint(conn.Minor))
	return conn
}

func Channel(conn *amqp.Connection) (*amqp.Channel) {
	c, err := conn.Channel()
	log.FatalOnError(err, "Failed to open a channel")

	return c
}