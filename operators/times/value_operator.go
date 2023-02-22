package times

import (
	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type TimeValue struct {
	utils       operatorUtilities
	operatorMap map[string]descriptors.Executor
}

func (tvo *TimeValue) GetOperators() []string {
	var operators []string
	for k := range tvo.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (tvo *TimeValue) GetExecutor(operatorName string) descriptors.Executor {
	return tvo.operatorMap[operatorName]
}

func (tvo *TimeValue) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// before To be read as val1 is before val2
func (tvo *TimeValue) before(val1 interface{}, val2 interface{}) bool {
	tVal1, err := tvo.utils.CastToTimeValue(val1)
	tVal2, err := tvo.utils.CastToTimeValue(val2)
	if err != nil {
		return false
	}
	return tVal1.Before(*tVal2)
}

// before To be read as val1 is after val2
func (tvo *TimeValue) after(val1 interface{}, val2 interface{}) bool {
	tVal1, err := tvo.utils.CastToTimeValue(val1)
	tVal2, err := tvo.utils.CastToTimeValue(val2)
	if err != nil {
		return false
	}
	return tVal1.After(*tVal2)
}

// during To be read as val1 falls within the range of val2
func (tvo *TimeValue) during(val1 interface{}, val2 interface{}) bool {
	tVal1, err := tvo.utils.CastToTimeValue(val1)
	tSlice2, err := tvo.utils.CastToTimeSlice(val2)
	if err != nil {
		return false
	}
	if len(tSlice2) < 2 || len(tSlice2) > 2 {
		return false
	}
	return tVal1.After(*tSlice2[0]) && tVal1.Before(*tSlice2[1])
}

func InitializeValueOperators(operatorUtilities operatorUtilities) *TimeValue {
	op := TimeValue{
		utils: operatorUtilities,
	}

	f := make(map[string]descriptors.Executor, 0)

	f[constants.Before] = op.before
	f[constants.After] = op.after
	f[constants.During] = op.during

	op.operatorMap = f
	return &op
}
