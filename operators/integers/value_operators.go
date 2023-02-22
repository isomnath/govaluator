package integers

import (
	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type IntegerValue struct {
	utils       operatorUtilities
	operatorMap map[string]descriptors.Executor
}

func (ivo *IntegerValue) GetOperators() []string {
	var operators []string
	for k := range ivo.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (ivo *IntegerValue) GetExecutor(operatorName string) descriptors.Executor {
	return ivo.operatorMap[operatorName]
}

func (ivo *IntegerValue) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// equals To be read as val1 equals val2
func (ivo *IntegerValue) equals(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bVal2, err2 := ivo.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 == bVal2
}

// notEquals To be read as val1 not-equals val2
func (ivo *IntegerValue) notEquals(val1 interface{}, val2 interface{}) bool {
	return !ivo.equals(val1, val2)
}

// lessThan To be read as val1 less-than val2
func (ivo *IntegerValue) lessThan(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bVal2, err2 := ivo.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 < bVal2
}

// lessThanOrEqualTo To be read as val1 less-than-or-equal-to val2
func (ivo *IntegerValue) lessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bVal2, err2 := ivo.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 <= bVal2
}

// lessThanAny To be read as val1 less-than-any of val2
func (ivo *IntegerValue) lessThanAny(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice2, err2 := ivo.utils.CastToIntegerSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	for _, bVal2 := range bSlice2 {
		if bVal1 < bVal2 {
			return true
		}
	}
	return false
}

// lessThanNone To be read as val1 less-than-none of val2
func (ivo *IntegerValue) lessThanNone(val1 interface{}, val2 interface{}) bool {
	return !ivo.lessThanAny(val1, val2)
}

// lessThanOrEqualToAny To be read as val1 less-than-or-equal-to-any of val2
func (ivo *IntegerValue) lessThanOrEqualToAny(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice2, err2 := ivo.utils.CastToIntegerSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	for _, bVal2 := range bSlice2 {
		if bVal1 <= bVal2 {
			return true
		}
	}
	return false
}

// lessThanOrEqualToNone To be read as val1 less-than-or-equal-to-none of val2
func (ivo *IntegerValue) lessThanOrEqualToNone(val1 interface{}, val2 interface{}) bool {
	return !ivo.lessThanOrEqualToAny(val1, val2)
}

// lessThanAll To be read as val1 less-than-all of val2
func (ivo *IntegerValue) lessThanAll(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice2, err2 := ivo.utils.CastToIntegerSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal2 := range bSlice2 {
		if bVal1 < bVal2 {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

// lessThanOrEqualToAll To be read as val1 less-than-or-equal-to-all of val2
func (ivo *IntegerValue) lessThanOrEqualToAll(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice2, err2 := ivo.utils.CastToIntegerSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal2 := range bSlice2 {
		if bVal1 <= bVal2 {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

// greaterThan To be read as val1 greater-than val2
func (ivo *IntegerValue) greaterThan(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bVal2, err2 := ivo.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 > bVal2
}

// greaterThanOrEqualTo To be read as val1 greater-than-or-equal-to val2
func (ivo *IntegerValue) greaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bVal2, err2 := ivo.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 >= bVal2
}

// greaterThanAny To be read as val1 greater-than-any of val2
func (ivo *IntegerValue) greaterThanAny(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice2, err2 := ivo.utils.CastToIntegerSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	for _, bVal2 := range bSlice2 {
		if bVal1 > bVal2 {
			return true
		}
	}
	return false
}

// greaterThanNone To be read as val1 greater-than-none of val2
func (ivo *IntegerValue) greaterThanNone(val1 interface{}, val2 interface{}) bool {
	return !ivo.greaterThanAny(val1, val2)
}

// greaterThanOrEqualToAny To be read as val1 greater-than-or-equal-to-any of val2
func (ivo *IntegerValue) greaterThanOrEqualToAny(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice2, err2 := ivo.utils.CastToIntegerSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	for _, bVal2 := range bSlice2 {
		if bVal1 >= bVal2 {
			return true
		}
	}
	return false
}

// greaterThanOrEqualToNone To be read as val1 greater-than-or-equal-to-none of val2
func (ivo *IntegerValue) greaterThanOrEqualToNone(val1 interface{}, val2 interface{}) bool {
	return !ivo.greaterThanOrEqualToAny(val1, val2)
}

// greaterThanAll To be read as val1 greater-than-all of val2
func (ivo *IntegerValue) greaterThanAll(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice2, err2 := ivo.utils.CastToIntegerSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal2 := range bSlice2 {
		if bVal1 > bVal2 {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

// greaterThanOrEqualToAll To be read as val1 greater-than-or-equal-to-all of val2
func (ivo *IntegerValue) greaterThanOrEqualToAll(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice2, err2 := ivo.utils.CastToIntegerSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal2 := range bSlice2 {
		if bVal1 >= bVal2 {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

// between To be read as val1 is within the range of values contained in val2 (boundaries not considered)
func (ivo *IntegerValue) between(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice2, err2 := ivo.utils.CastToIntegerSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	if len(bSlice2) < 2 || len(bSlice2) > 2 {
		return false
	}
	return bVal1 > bSlice2[0] && bVal1 < bSlice2[1]
}

// notBetween To be read as val1 is not within the range of values contained in val2(boundaries not considered)
func (ivo *IntegerValue) notBetween(val1 interface{}, val2 interface{}) bool {
	return !ivo.between(val1, val2)
}

// anyOf To be read as val1 is any-of val2
func (ivo *IntegerValue) anyOf(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := ivo.utils.CastToIntegerValue(val1)
	bSlice, err2 := ivo.utils.CastToIntegerSlice(val2)
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
func (ivo *IntegerValue) noneOf(val1 interface{}, val2 interface{}) bool {
	return !ivo.anyOf(val1, val2)
}

func InitializeValueOperators(operatorUtilities operatorUtilities) *IntegerValue {
	op := IntegerValue{
		utils: operatorUtilities,
	}
	operatorMap := make(map[string]descriptors.Executor, 0)
	operatorMap[constants.Equals] = op.equals
	operatorMap[constants.NotEquals] = op.notEquals

	operatorMap[constants.LessThan] = op.lessThan
	operatorMap[constants.LessThanOrEqualTo] = op.lessThanOrEqualTo
	operatorMap[constants.LessThanAny] = op.lessThanAny
	operatorMap[constants.LessThanOrEqualToAny] = op.lessThanOrEqualToAny
	operatorMap[constants.LessThanNone] = op.lessThanNone
	operatorMap[constants.LessThanOrEqualToNone] = op.lessThanOrEqualToNone
	operatorMap[constants.LessThanAll] = op.lessThanAll
	operatorMap[constants.LessThanOrEqualToAll] = op.lessThanOrEqualToAll
	operatorMap[constants.GreaterThan] = op.greaterThan
	operatorMap[constants.GreaterThanOrEqualTo] = op.greaterThanOrEqualTo
	operatorMap[constants.GreaterThanAny] = op.greaterThanAny
	operatorMap[constants.GreaterThanOrEqualToAny] = op.greaterThanOrEqualToAny
	operatorMap[constants.GreaterThanNone] = op.greaterThanNone
	operatorMap[constants.GreaterThanOrEqualToNone] = op.greaterThanOrEqualToNone
	operatorMap[constants.GreaterThanAll] = op.greaterThanAll
	operatorMap[constants.GreaterThanOrEqualToAll] = op.greaterThanOrEqualToAll
	operatorMap[constants.Between] = op.between
	operatorMap[constants.NotBetween] = op.notBetween

	operatorMap[constants.AnyOf] = op.anyOf
	operatorMap[constants.NoneOf] = op.noneOf

	op.operatorMap = operatorMap
	return &op
}
