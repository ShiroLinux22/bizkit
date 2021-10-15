package util

import (
	"encoding/json"
)

var (
	log = Logger {
		Name: "util",
	}
)

func ToJson(i interface {}) (string) {
	formatted, err := json.Marshal(i)
	if err != nil {
		log.Error("Failed to Convert String to JSON: ", err)
	}
	text := string(formatted)

	return text
}
