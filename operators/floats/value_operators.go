package floats

import (
	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type FloatValue struct {
	utils       operatorUtilities
	operatorMap map[string]descriptors.Executor
}

func (fvo *FloatValue) GetOperators() []string {
	var operators []string
	for k := range fvo.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (fvo *FloatValue) GetExecutor(operatorName string) descriptors.Executor {
	return fvo.operatorMap[operatorName]
}

func (fvo *FloatValue) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// equals To be read as val1 equals val2
func (fvo *FloatValue) equals(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bVal2, err2 := fvo.utils.CastToFloatValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 == bVal2
}

// notEquals To be read as val1 not-equals val2
func (fvo *FloatValue) notEquals(val1 interface{}, val2 interface{}) bool {
	return !fvo.equals(val1, val2)
}

// lessThan To be read as val1 less-than val2
func (fvo *FloatValue) lessThan(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bVal2, err2 := fvo.utils.CastToFloatValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 < bVal2
}

// lessThanOrEqualTo To be read as val1 less-than-or-equal-to val2
func (fvo *FloatValue) lessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bVal2, err2 := fvo.utils.CastToFloatValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 <= bVal2
}

// lessThanAny To be read as val1 less-than-any of val2
func (fvo *FloatValue) lessThanAny(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice2, err2 := fvo.utils.CastToFloatSlice(val2)
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
func (fvo *FloatValue) lessThanNone(val1 interface{}, val2 interface{}) bool {
	return !fvo.lessThanAny(val1, val2)
}

// lessThanOrEqualToAny To be read as val1 less-than-or-equal-to-any of val2
func (fvo *FloatValue) lessThanOrEqualToAny(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice2, err2 := fvo.utils.CastToFloatSlice(val2)
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
func (fvo *FloatValue) lessThanOrEqualToNone(val1 interface{}, val2 interface{}) bool {
	return !fvo.lessThanOrEqualToAny(val1, val2)
}

// lessThanAll To be read as val1 less-than-all of val2
func (fvo *FloatValue) lessThanAll(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice2, err2 := fvo.utils.CastToFloatSlice(val2)
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
func (fvo *FloatValue) lessThanOrEqualToAll(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice2, err2 := fvo.utils.CastToFloatSlice(val2)
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
func (fvo *FloatValue) greaterThan(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bVal2, err2 := fvo.utils.CastToFloatValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 > bVal2
}

// greaterThanOrEqualTo To be read as val1 greater-than-or-equal-to val2
func (fvo *FloatValue) greaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bVal2, err2 := fvo.utils.CastToFloatValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 >= bVal2
}

// greaterThanAny To be read as val1 greater-than-any of val2
func (fvo *FloatValue) greaterThanAny(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice2, err2 := fvo.utils.CastToFloatSlice(val2)
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
func (fvo *FloatValue) greaterThanNone(val1 interface{}, val2 interface{}) bool {
	return !fvo.greaterThanAny(val1, val2)
}

// greaterThanOrEqualToAny To be read as val1 greater-than-or-equal-to-any of val2
func (fvo *FloatValue) greaterThanOrEqualToAny(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice2, err2 := fvo.utils.CastToFloatSlice(val2)
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
func (fvo *FloatValue) greaterThanOrEqualToNone(val1 interface{}, val2 interface{}) bool {
	return !fvo.greaterThanOrEqualToAny(val1, val2)
}

// greaterThanAll To be read as val1 greater-than-all of val2
func (fvo *FloatValue) greaterThanAll(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice2, err2 := fvo.utils.CastToFloatSlice(val2)
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
func (fvo *FloatValue) greaterThanOrEqualToAll(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice2, err2 := fvo.utils.CastToFloatSlice(val2)
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
func (fvo *FloatValue) between(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice2, err2 := fvo.utils.CastToFloatSlice(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	if len(bSlice2) < 2 || len(bSlice2) > 2 {
		return false
	}
	return bVal1 > bSlice2[0] && bVal1 < bSlice2[1]
}

// notBetween To be read as val1 is not within the range of values contained in val2(boundaries not considered)
func (fvo *FloatValue) notBetween(val1 interface{}, val2 interface{}) bool {
	return !fvo.between(val1, val2)
}

// anyOf To be read as val1 is any-of val2
func (fvo *FloatValue) anyOf(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := fvo.utils.CastToFloatValue(val1)
	bSlice, err2 := fvo.utils.CastToFloatSlice(val2)
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
func (fvo *FloatValue) noneOf(val1 interface{}, val2 interface{}) bool {
	return !fvo.anyOf(val1, val2)
}

func InitializeValueOperators(operatorUtilities operatorUtilities) *FloatValue {
	op := FloatValue{
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
