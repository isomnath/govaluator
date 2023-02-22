package defaults

import (
	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type DefaultOperator struct {
	operatorMap map[string]descriptors.Executor
}

func (do *DefaultOperator) GetOperators() []string {
	var operators []string
	for k := range do.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (do *DefaultOperator) GetExecutor(_ string) descriptors.Executor {
	return do.operatorMap[constants.Default]
}

func (do *DefaultOperator) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

func (do *DefaultOperator) defaultFn(_ interface{}, _ interface{}) bool {
	return false
}

func InitializeOperators() *DefaultOperator {
	op := DefaultOperator{}

	f := make(map[string]descriptors.Executor, 0)
	f[constants.Default] = op.defaultFn

	op.operatorMap = f
	return &op
}
