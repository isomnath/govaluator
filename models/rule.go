package models

import (
	"fmt"

	"github.com/isomnath/govaluator/constants"
)

type Rule struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Remark          string           `json:"remark"`
	Engine          string           `json:"engine"`
	Linear          *Linear          `json:"linear,omitempty"`
	LinearWaterfall *LinearWaterfall `json:"linear_waterfall,omitempty"`
}

func (r *Rule) Scan(data interface{}) error {
	err := scan(data, r)
	if err != nil {
		return fmt.Errorf("failed to deserialize rule: %s", err.Error())
	}

	switch r.Engine {
	case constants.LinearEngine:
		err = r.Linear.Scan(data)
		if err != nil {
			return fmt.Errorf("failed to deserialize rule: %s", err.Error())
		}
		r.LinearWaterfall = nil
		break
	case constants.LinearWaterfallEngine:
		err = r.LinearWaterfall.Scan(data)
		if err != nil {
			return fmt.Errorf("failed to deserialize rule: %s", err.Error())
		}
		r.Linear = nil
		break
	default:
		return fmt.Errorf("failed to deserialize rule: unsupported rule engine: %s", r.Engine)
	}
	return nil
}
