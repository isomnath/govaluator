package comparators

import (
	"github.com/isomnath/govaluator/models"
)

type staticComparator struct {
	operationManager operationManager
}

func (comp *staticComparator) ExecuteCriterion(data map[string]interface{}, criterion models.Criterion) (bool, interface{}) {
	c := criterion.Static
	fieldValue := data[c.FieldOne]
	staticComparisonValue := c.Value
	operatorName := c.Operator
	//operator can never be nil as default operator is returned when data type is unidentified
	operator := comp.operationManager.GetOperator(fieldValue)
	executor := operator.GetExecutor(operatorName)
	if executor == nil {
		return false, renderResult(criterion, false)
	}
	resultFlag := operator.ExecuteOperator(executor, fieldValue, staticComparisonValue)
	return resultFlag, renderResult(criterion, resultFlag)
}

func initializeStaticComparator(operationManager operationManager) Comparators {
	return &staticComparator{
		operationManager: operationManager,
	}
}
