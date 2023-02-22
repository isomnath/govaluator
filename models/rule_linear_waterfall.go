package models

import (
	"fmt"
)

type LinearWaterfall struct {
	ExitOn             string      `json:"exit_on"`
	Criteria           []Criterion `json:"criteria"`
	DefaultReturnValue interface{} `json:"default_return_value"`
}

func (lw *LinearWaterfall) Scan(data interface{}) error {
	err := scan(data, lw)
	if err != nil {
		return fmt.Errorf("failed to deserialize linear waterfall rule: %s", err.Error())
	}
	return nil
}
