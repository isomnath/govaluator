package booleans

import (
	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type BooleanValue struct {
	utils       operatorUtilities
	operatorMap map[string]descriptors.Executor
}

func (bvo *BooleanValue) GetOperators() []string {
	var operators []string
	for k := range bvo.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (bvo *BooleanValue) GetExecutor(operatorName string) descriptors.Executor {
	return bvo.operatorMap[operatorName]
}

func (bvo *BooleanValue) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// equals To be read as val1 equals val2
func (bvo *BooleanValue) equals(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := bvo.utils.CastToBooleanValue(val1)
	bVal2, err2 := bvo.utils.CastToBooleanValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 == bVal2
}

// notEquals To be read as val1 not-equals val2
func (bvo *BooleanValue) notEquals(val1 interface{}, val2 interface{}) bool {
	return !bvo.equals(val1, val2)
}

// anyOf To be read as val1 is any-of val2
func (bvo *BooleanValue) anyOf(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := bvo.utils.CastToBooleanValue(val1)
	bSlice, err2 := bvo.utils.CastToBooleanSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, v := range bSlice {
		if bVal1 == v {
			flag = true
			break
		}
	}
	return flag
}

// noneOf To be read as val1 is none-of val2
func (bvo *BooleanValue) noneOf(val1 interface{}, val2 interface{}) bool {
	return !bvo.anyOf(val1, val2)
}

func InitializeValueOperators(operatorUtilities operatorUtilities) *BooleanValue {
	op := BooleanValue{
		utils: operatorUtilities,
	}

	operatorMap := make(map[string]descriptors.Executor, 0)
	operatorMap[constants.Equals] = op.equals
	operatorMap[constants.NotEquals] = op.notEquals
	operatorMap[constants.AnyOf] = op.anyOf
	operatorMap[constants.NoneOf] = op.noneOf

	op.operatorMap = operatorMap
	return &op
}
