package comparators

import (
	"github.com/isomnath/govaluator/models"
)

func renderResult(criterion models.Criterion, resultFlag bool) interface{} {
	if resultFlag {
		if criterion.TruthyValue != nil {
			return criterion.TruthyValue
		}
		return resultFlag
	}
	if criterion.FalseyValue != nil {
		return criterion.FalseyValue
	}
	return resultFlag
}
