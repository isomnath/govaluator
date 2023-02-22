package models

import (
	"fmt"
)

type Dynamic struct {
	FieldOne string `json:"field_one"`
	Operator string `json:"operator"`
	FieldTwo string `json:"field_two"`
}

func (s *Dynamic) Scan(data interface{}) error {
	err := scan(data, s)
	if err != nil {
		return fmt.Errorf("failed to deserialize dynamic criterion: %s", err.Error())
	}
	return nil
}
