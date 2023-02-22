package models

import (
	"fmt"
)

type Static struct {
	FieldOne string      `json:"field_one"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

func (s *Static) Scan(data interface{}) error {
	err := scan(data, s)
	if err != nil {
		return fmt.Errorf("failed to deserialize static criterion: %s", err.Error())
	}
	return nil
}
