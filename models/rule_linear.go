package models

import (
	"fmt"
)

type Linear struct {
	Expression string      `json:"expression"`
	Criteria   []Criterion `json:"criteria"`
}

func (l *Linear) Scan(data interface{}) error {
	err := scan(data, l)
	if err != nil {
		return fmt.Errorf("failed to deserialize linear rule: %s", err.Error())
	}
	return nil
}
