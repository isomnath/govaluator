package exitonstrategies

import (
	"github.com/isomnath/govaluator/models"
)

type firstTrue struct{}

func (strategy *firstTrue) GetResultValue(results []models.TransientResult) interface{} {
	filteredResults := filter(true, results)
	if len(filteredResults) > 0 {
		return filteredResults[0].Result
	}
	return nil
}

func initializeFirstTrue() *firstTrue {
	return &firstTrue{}
}
