package operators

import (
	"time"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/operators/booleans"
	"github.com/isomnath/govaluator/operators/defaults"
	"github.com/isomnath/govaluator/operators/floats"
	"github.com/isomnath/govaluator/operators/integers"
	"github.com/isomnath/govaluator/operators/nils"
	"github.com/isomnath/govaluator/operators/strings"
	"github.com/isomnath/govaluator/operators/times"
)

type OperationManager struct {
	operators map[string]Operators
}

func (m *OperationManager) GetOperator(val interface{}) Operators {
	dataTypeAlias := getDataTypeAlias(val)
	return m.operators[dataTypeAlias]
}

func (m *OperationManager) GetAllDataTypes() []string {
	var dataTypes []string
	for k := range m.operators {
		if k != constants.Default && k != constants.Nil {
			dataTypes = append(dataTypes, k)
		}
	}
	return dataTypes
}

func (m *OperationManager) GetOperators(dataType string) []string {
	return m.operators[dataType].GetOperators()
}

func getDataTypeAlias(val interface{}) string {
	switch val.(type) {
	case nil:
		return constants.Nil
	case string:
		return constants.StringValue
	case []string:
		return constants.StringSlice
	case bool:
		return constants.BooleanValue
	case []bool:
		return constants.BooleanSlice
	case int:
		return constants.IntegerValue
	case []int:
		return constants.IntegerSlice
	case int8:
		return constants.Integer8Value
	case []int8:
		return constants.Integer8Slice
	case int16:
		return constants.Integer16Value
	case []int16:
		return constants.Integer16Slice
	case int32:
		return constants.Integer32Value
	case []int32:
		return constants.Integer32Slice
	case int64:
		return constants.Integer64Value
	case []int64:
		return constants.Integer64Slice
	case float32:
		return constants.Float32Value
	case []float32:
		return constants.Float32Slice
	case float64:
		return constants.Float64Value
	case []float64:
		return constants.Float64Slice
	case time.Time:
		return constants.TimeValue
	case *time.Time:
		return constants.TimeValue
	default:
		return constants.Default
	}
}

func InitializeOperators(utilities operatorUtilities) *OperationManager {
	o := OperationManager{operators: make(map[string]Operators)}

	o.operators[constants.Nil] = nils.InitializeNilOperators()
	o.operators[constants.Default] = defaults.InitializeOperators()

	o.operators[constants.BooleanSlice] = booleans.InitializeSliceOperators(utilities)
	o.operators[constants.BooleanValue] = booleans.InitializeValueOperators(utilities)

	floatValueOperator := floats.InitializeValueOperators(utilities)
	floatSliceOperator := floats.InitializeSliceOperators(utilities)
	o.operators[constants.Float32Slice] = floatSliceOperator
	o.operators[constants.Float32Value] = floatValueOperator
	o.operators[constants.Float64Slice] = floatSliceOperator
	o.operators[constants.Float64Value] = floatValueOperator

	integerValueOperator := integers.InitializeValueOperators(utilities)
	integerSliceOperator := integers.InitializeSliceOperators(utilities)
	o.operators[constants.IntegerSlice] = integerSliceOperator
	o.operators[constants.IntegerValue] = integerValueOperator
	o.operators[constants.Integer8Slice] = integerSliceOperator
	o.operators[constants.Integer8Value] = integerValueOperator
	o.operators[constants.Integer16Slice] = integerSliceOperator
	o.operators[constants.Integer16Value] = integerValueOperator
	o.operators[constants.Integer32Slice] = integerSliceOperator
	o.operators[constants.Integer32Value] = integerValueOperator
	o.operators[constants.Integer64Slice] = integerSliceOperator
	o.operators[constants.Integer64Value] = integerValueOperator

	o.operators[constants.StringSlice] = strings.InitializeSliceOperators(utilities)
	o.operators[constants.StringValue] = strings.InitializeValueOperators(utilities)

	o.operators[constants.TimeValue] = times.InitializeValueOperators(utilities)

	return &o
}
