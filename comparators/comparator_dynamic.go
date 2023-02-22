package comparators

import (
	"github.com/isomnath/govaluator/models"
)

type dynamicComparator struct {
	operationManager operationManager
}

func (comp *dynamicComparator) ExecuteCriterion(data map[string]interface{}, criterion models.Criterion) (bool, interface{}) {
	c := criterion.Dynamic
	firstFieldValue := data[c.FieldOne]
	secondFieldValue := data[c.FieldTwo]
	operatorName := c.Operator
	//operator can never be nil as default operator is returned when data type is unidentified
	operator := comp.operationManager.GetOperator(firstFieldValue)
	executor := operator.GetExecutor(operatorName)
	if executor == nil {
		return false, renderResult(criterion, false)
	}
	resultFlag := operator.ExecuteOperator(executor, firstFieldValue, secondFieldValue)
	return resultFlag, renderResult(criterion, resultFlag)
}

func initializeDynamicComparator(operationManager operationManager) Comparators {
	return &dynamicComparator{
		operationManager: operationManager,
	}
}
