package util

import (
	"encoding/json"
)

func ToJson(i interface {}) (string, error) {
	formatted, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	text := string(formatted)

	return text, nil
}
