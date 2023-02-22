package operators

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type OperationManagerTestSuite struct {
	suite.Suite
	operatorUtilities operatorUtilities
	operationManager  *OperationManager
}

func (suite *OperationManagerTestSuite) SetupTest() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.operationManager = InitializeOperators(suite.operatorUtilities)
}

func (suite *OperationManagerTestSuite) TestGetOperator() {
	timeNow := time.Now()
	unidentifiedInterface := make(map[string]string)

	type args struct {
		val interface{}
	}
	tests := []struct {
		name        string
		args        args
		expectation Operators
	}{
		{
			name:        "Nil Value",
			args:        args{val: nil},
			expectation: suite.operationManager.operators[constants.Nil],
		},
		{
			name:        "String Value",
			args:        args{val: "ABC"},
			expectation: suite.operationManager.operators[constants.StringValue],
		},
		{
			name:        "String Slice",
			args:        args{val: []string{"ABC"}},
			expectation: suite.operationManager.operators[constants.StringSlice],
		},
		{
			name:        "Boolean Value",
			args:        args{val: true},
			expectation: suite.operationManager.operators[constants.BooleanValue],
		},
		{
			name:        "Boolean Slice",
			args:        args{val: []bool{true, false}},
			expectation: suite.operationManager.operators[constants.BooleanSlice],
		},
		{
			name:        "Float32 Value",
			args:        args{val: float32(123)},
			expectation: suite.operationManager.operators[constants.Float32Value],
		},
		{
			name:        "Float32 Slice",
			args:        args{val: []float32{123, 3453}},
			expectation: suite.operationManager.operators[constants.Float32Slice],
		},
		{
			name:        "Float64 Value",
			args:        args{val: float64(123)},
			expectation: suite.operationManager.operators[constants.Float64Value],
		},
		{
			name:        "Float64 Slice",
			args:        args{val: []float64{123, 3453}},
			expectation: suite.operationManager.operators[constants.Float64Slice],
		},
		{
			name:        "Integer Value",
			args:        args{val: 123},
			expectation: suite.operationManager.operators[constants.IntegerValue],
		},
		{
			name:        "Integer Slice",
			args:        args{val: []int{123, 3453}},
			expectation: suite.operationManager.operators[constants.IntegerSlice],
		},
		{
			name:        "Integer8 Value",
			args:        args{val: int8(123)},
			expectation: suite.operationManager.operators[constants.Integer8Value],
		},
		{
			name:        "Integer8 Slice",
			args:        args{val: []int8{123, 111}},
			expectation: suite.operationManager.operators[constants.Integer8Slice],
		},
		{
			name:        "Integer16 Value",
			args:        args{val: int16(123)},
			expectation: suite.operationManager.operators[constants.Integer16Value],
		},
		{
			name:        "Integer16 Slice",
			args:        args{val: []int16{123, 111}},
			expectation: suite.operationManager.operators[constants.Integer16Slice],
		},
		{
			name:        "Integer32 Value",
			args:        args{val: int32(123)},
			expectation: suite.operationManager.operators[constants.Integer32Value],
		},
		{
			name:        "Integer32 Slice",
			args:        args{val: []int32{123, 111}},
			expectation: suite.operationManager.operators[constants.Integer32Slice],
		},
		{
			name:        "Integer64 Value",
			args:        args{val: int64(123)},
			expectation: suite.operationManager.operators[constants.Integer64Value],
		},
		{
			name:        "Integer64 Slice",
			args:        args{val: []int64{123, 3453}},
			expectation: suite.operationManager.operators[constants.Integer64Slice],
		},
		{
			name:        "Time Value",
			args:        args{val: timeNow},
			expectation: suite.operationManager.operators[constants.TimeValue],
		},
		{
			name:        "Pointer To Time Value",
			args:        args{val: &timeNow},
			expectation: suite.operationManager.operators[constants.TimeValue],
		},
		{
			name:        "Unidentified type resolving to default type",
			args:        args{val: unidentifiedInterface},
			expectation: suite.operationManager.operators[constants.Default],
		},
	}
	for _, tt := range tests {
		operator := suite.operationManager.GetOperator(tt.args.val)
		suite.Equal(tt.expectation, operator)
	}
}

func (suite *OperationManagerTestSuite) TestGetAllDataTypes() {
	expected := []string{constants.StringValue, constants.StringSlice,
		constants.BooleanValue, constants.BooleanSlice,
		constants.IntegerValue, constants.IntegerSlice, constants.Integer8Value, constants.Integer8Slice, constants.Integer16Value,
		constants.Integer16Slice, constants.Integer32Value, constants.Integer32Slice, constants.Integer64Value, constants.Integer64Slice,
		constants.Float32Value, constants.Float32Slice, constants.Float64Value, constants.Float64Slice,
		constants.TimeValue}
	actual := suite.operationManager.GetAllDataTypes()

	sort.Strings(expected)
	sort.Strings(actual)
	suite.Equal(expected, actual)
}

func (suite *OperationManagerTestSuite) TestGetOperators() {
	expected := []string{constants.Before, constants.After, constants.During}
	actual := suite.operationManager.GetOperators(constants.TimeValue)

	sort.Strings(expected)
	sort.Strings(actual)
	suite.Equal(expected, actual)
}

func TestOperationManagerTestSuite(t *testing.T) {
	suite.Run(t, new(OperationManagerTestSuite))
}
