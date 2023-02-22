package models

import (
	"fmt"

	"github.com/isomnath/govaluator/constants"
)

type Criterion struct {
	ID          string      `json:"id"`
	Alias       string      `json:"alias"`
	Comparator  string      `json:"comparator"`
	Static      *Static     `json:"static,omitempty"`
	Dynamic     *Dynamic    `json:"dynamic,omitempty"`
	TruthyValue interface{} `json:"truthy_value,omitempty"`
	FalseyValue interface{} `json:"falsey_value,omitempty"`
}

func (c *Criterion) Scan(data interface{}) error {
	err := scan(data, c)
	if err != nil {
		return fmt.Errorf("failed to deserialize criterion: %s", err.Error())
	}

	switch c.Comparator {
	case constants.StaticComparator:
		err = c.Static.Scan(data)
		if err != nil {
			return fmt.Errorf("failed to deserialize criterion: %s", err.Error())
		}
		c.Dynamic = nil
		break
	case constants.DynamicComparator:
		err = c.Dynamic.Scan(data)
		if err != nil {
			return fmt.Errorf("failed to deserialize criterion: %s", err.Error())
		}
		c.Static = nil
		break
	default:
		return fmt.Errorf("failed to deserialize criterion: unsupported comparator: %s", c.Comparator)
	}
	return nil
}
