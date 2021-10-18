/*
	Struct(s) for handling amqp connections
    Copyright (C) 2021 jacany <jack@chaker.net>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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