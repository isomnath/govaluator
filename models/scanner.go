package models

import (
	"encoding/json"
	"fmt"
)

func scan(data interface{}, destination interface{}) error {
	switch data.(type) {
	case []byte:
		err := json.Unmarshal(data.([]byte), destination)
		if err != nil {
			return fmt.Errorf("invalid json structure: %s", err.Error())
		}
	default:
		return fmt.Errorf("value is not a json")
	}
	return nil
}
