package comparators

import (
	"github.com/isomnath/govaluator/models"
	"github.com/isomnath/govaluator/operators"
)

type operationManager interface {
	GetOperator(interface{}) operators.Operators
}

type Comparators interface {
	ExecuteCriterion(map[string]interface{}, models.Criterion) (bool, interface{})
}
