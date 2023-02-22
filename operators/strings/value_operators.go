package strings

import (
	"regexp"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type StringValue struct {
	utils       operatorUtilities
	operatorMap map[string]descriptors.Executor
}

func (svo *StringValue) GetOperators() []string {
	var operators []string
	for k := range svo.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (svo *StringValue) GetExecutor(operatorName string) descriptors.Executor {
	return svo.operatorMap[operatorName]
}

func (svo *StringValue) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// equals To be read as val1 equals val2
func (svo *StringValue) equals(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := svo.utils.CastToStringValue(val1)
	bVal2, err2 := svo.utils.CastToStringValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	return bVal1 == bVal2
}

// notEquals To be read as val1 not-equals val2
func (svo *StringValue) notEquals(val1 interface{}, val2 interface{}) bool {
	return !svo.equals(val1, val2)
}

// anyOf To be read as val1 is any-of val2
func (svo *StringValue) anyOf(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := svo.utils.CastToStringValue(val1)
	bSlice, err2 := svo.utils.CastToStringSlice(val2)
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
func (svo *StringValue) noneOf(val1 interface{}, val2 interface{}) bool {
	return !svo.anyOf(val1, val2)
}

// regexMatch To be read as val1 matches regex val2
func (svo *StringValue) regexMatch(val1 interface{}, val2 interface{}) bool {
	bVal1, err1 := svo.utils.CastToStringValue(val1)
	bVal2, err2 := svo.utils.CastToStringValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	matched, err := regexp.MatchString(bVal2, bVal1)
	if err != nil {
		return false
	}
	return matched
}

// regexNotMatch To be read as val1 does match regex val2
func (svo *StringValue) regexNotMatch(val1 interface{}, val2 interface{}) bool {
	return !svo.regexMatch(val1, val2)
}

func InitializeValueOperators(operatorUtilities operatorUtilities) *StringValue {
	op := StringValue{
		utils: operatorUtilities,
	}

	f := make(map[string]descriptors.Executor, 0)

	f[constants.Equals] = op.equals
	f[constants.NotEquals] = op.notEquals

	f[constants.AnyOf] = op.anyOf
	f[constants.NoneOf] = op.noneOf

	f[constants.RegexMatch] = op.regexMatch
	f[constants.RegexNotMatch] = op.regexNotMatch

	op.operatorMap = f
	return &op
}
