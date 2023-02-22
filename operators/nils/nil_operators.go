package nils

import (
	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type NilValue struct {
	operatorMap map[string]descriptors.Executor
}

func (n *NilValue) GetOperators() []string {
	var operators []string
	for k := range n.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (n *NilValue) GetExecutor(operatorName string) descriptors.Executor {
	return n.operatorMap[operatorName]
}

func (n *NilValue) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// equals To be read as val1 equals val2
func (n *NilValue) equals(_ interface{}, val2 interface{}) bool {
	return val2 == nil
}

// notEquals To be read as val1 not-equals val2
func (n *NilValue) notEquals(val1 interface{}, val2 interface{}) bool {
	return !n.equals(val1, val2)
}

func (n *NilValue) defaultFalse(_ interface{}, _ interface{}) bool {
	return false
}

func InitializeNilOperators() *NilValue {
	n := NilValue{}

	operatorMap := make(map[string]descriptors.Executor, 0)
	operatorMap[constants.Equals] = n.equals
	operatorMap[constants.NotEquals] = n.notEquals

	operatorMap[constants.LessThan] = n.defaultFalse
	operatorMap[constants.LessThanOrEqualTo] = n.defaultFalse
	operatorMap[constants.LessThanAny] = n.defaultFalse
	operatorMap[constants.LessThanOrEqualToAny] = n.defaultFalse
	operatorMap[constants.LessThanNone] = n.defaultFalse
	operatorMap[constants.LessThanOrEqualToNone] = n.defaultFalse
	operatorMap[constants.LessThanAll] = n.defaultFalse
	operatorMap[constants.LessThanOrEqualToAll] = n.defaultFalse
	operatorMap[constants.GreaterThan] = n.defaultFalse
	operatorMap[constants.GreaterThanOrEqualTo] = n.defaultFalse
	operatorMap[constants.GreaterThanAny] = n.defaultFalse
	operatorMap[constants.GreaterThanOrEqualToAny] = n.defaultFalse
	operatorMap[constants.GreaterThanNone] = n.defaultFalse
	operatorMap[constants.GreaterThanOrEqualToNone] = n.defaultFalse
	operatorMap[constants.GreaterThanAll] = n.defaultFalse
	operatorMap[constants.GreaterThanOrEqualToAll] = n.defaultFalse
	operatorMap[constants.Between] = n.defaultFalse
	operatorMap[constants.NotBetween] = n.defaultFalse

	operatorMap[constants.AnyOf] = n.defaultFalse
	operatorMap[constants.NoneOf] = n.defaultFalse

	operatorMap[constants.RegexMatch] = n.defaultFalse
	operatorMap[constants.RegexNotMatch] = n.defaultFalse

	operatorMap[constants.UnorderedEquals] = n.defaultFalse
	operatorMap[constants.UnorderedNotEquals] = n.defaultFalse
	operatorMap[constants.OrderedEquals] = n.defaultFalse
	operatorMap[constants.OrderedNotEquals] = n.defaultFalse

	operatorMap[constants.AnyLessThan] = n.defaultFalse
	operatorMap[constants.AnyLessThanOrEqualTo] = n.defaultFalse
	operatorMap[constants.NoneLessThan] = n.defaultFalse
	operatorMap[constants.NoneLessThanOrEqualTo] = n.defaultFalse
	operatorMap[constants.AllLessThan] = n.defaultFalse
	operatorMap[constants.AllLessThanOrEqualTo] = n.defaultFalse
	operatorMap[constants.AnyGreaterThan] = n.defaultFalse
	operatorMap[constants.AnyGreaterThanOrEqualTo] = n.defaultFalse
	operatorMap[constants.NoneGreaterThan] = n.defaultFalse
	operatorMap[constants.NoneGreaterThanOrEqualTo] = n.defaultFalse
	operatorMap[constants.AllGreaterThan] = n.defaultFalse
	operatorMap[constants.AllGreaterThanOrEqualTo] = n.defaultFalse
	operatorMap[constants.ReversedBetween] = n.defaultFalse
	operatorMap[constants.ReversedNotBetween] = n.defaultFalse

	operatorMap[constants.SupersetOf] = n.defaultFalse
	operatorMap[constants.NotSupersetOf] = n.defaultFalse
	operatorMap[constants.SubsetOf] = n.defaultFalse
	operatorMap[constants.NotSubsetOf] = n.defaultFalse
	operatorMap[constants.Intersection] = n.defaultFalse
	operatorMap[constants.NotIntersection] = n.defaultFalse
	operatorMap[constants.ReversedAnyOf] = n.defaultFalse
	operatorMap[constants.ReversedNoneOf] = n.defaultFalse

	operatorMap[constants.SizeEquals] = n.defaultFalse
	operatorMap[constants.SizeNotEquals] = n.defaultFalse
	operatorMap[constants.SizeLessThan] = n.defaultFalse
	operatorMap[constants.SizeLessThanOrEqualTo] = n.defaultFalse
	operatorMap[constants.SizeGreaterThan] = n.defaultFalse
	operatorMap[constants.SizeGreaterThanOrEqualTo] = n.defaultFalse

	operatorMap[constants.AnyRegexMatch] = n.defaultFalse
	operatorMap[constants.NoneRegexMatch] = n.defaultFalse
	operatorMap[constants.AllRegexMatch] = n.defaultFalse

	operatorMap[constants.Before] = n.defaultFalse
	operatorMap[constants.After] = n.defaultFalse
	operatorMap[constants.During] = n.defaultFalse

	n.operatorMap = operatorMap
	return &n
}
