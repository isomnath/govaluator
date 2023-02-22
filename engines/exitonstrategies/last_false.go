package exitonstrategies

import (
	"github.com/isomnath/govaluator/models"
)

type lastFalse struct{}

func (strategy *lastFalse) GetResultValue(results []models.TransientResult) interface{} {
	filteredResults := filter(false, results)
	if len(filteredResults) > 0 {
		return filteredResults[len(filteredResults)-1].Result
	}
	return nil
}

func initializeLastFalse() *lastFalse {
	return &lastFalse{}
}
