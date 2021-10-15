package amqp

import (
	"os"

	"github.com/chakernet/ryuko/common/util"
	"github.com/streadway/amqp"
)

var (
	log = util.Logger {
		Name: "amqp",
	}
)

func Connect() (*amqp.Connection, error) {
	uri := os.Getenv("AMQP_URI")

	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Channel(conn *amqp.Connection) (*amqp.Channel, error) {
	c, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return c, nil
}