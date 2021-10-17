package util

import (
	"encoding/json"
	"reflect"
)

func ToJson(i interface {}) (string, error) {
	formatted, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	text := string(formatted)

	return text, nil
}

func IsZero(v reflect.Value) bool {
	return !v.IsValid() || reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}