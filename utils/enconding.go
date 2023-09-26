package utils

import (
	"encoding/json"
)

func ToJson(v interface{}) (string, error) {
    b, err := json.Marshal(v)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return "", err
    }
    return string(b), nil
}