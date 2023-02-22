package exitonstrategies

import (
	"github.com/isomnath/govaluator/models"
)

type firstFalse struct{}

func (strategy *firstFalse) GetResultValue(results []models.TransientResult) interface{} {
	filteredResults := filter(false, results)
	if len(filteredResults) > 0 {
		return filteredResults[0].Result
	}
	return nil
}

func initializeFirstFalse() *firstFalse {
	return &firstFalse{}
}
