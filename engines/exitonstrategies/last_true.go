package exitonstrategies

import (
	"github.com/isomnath/govaluator/models"
)

type lastTrue struct{}

func (strategy *lastTrue) GetResultValue(results []models.TransientResult) interface{} {
	filteredResults := filter(true, results)
	if len(filteredResults) > 0 {
		return filteredResults[len(filteredResults)-1].Result
	}
	return nil
}

func initializeLastTrue() *lastTrue {
	return &lastTrue{}
}
