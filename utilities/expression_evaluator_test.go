package utilities

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
)

type ExpressionEvaluatorTestSuite struct {
	suite.Suite
	utils *Utilities
}

func (suite *ExpressionEvaluatorTestSuite) SetupTest() {
	suite.utils = InitializeUtilities()
}

func (suite *ExpressionEvaluatorTestSuite) TestEvaluate_Error() {
	type args struct {
		data       map[string]interface{}
		expression string
	}
	var testCases = []struct {
		name          string
		args          args
		expectedError error
	}{
		{
			name: "Case 1",
			args: args{
				data:       map[string]interface{}{"A": false, "B": true, "C": true},
				expression: "(({{.A}} and {{.B}}) or {{.C}})",
			},

			expectedError: fmt.Errorf(constants.ErrorInitializingExpressionParser),
		},
		{
			name: "Case 2",
			args: args{
				data:       map[string]interface{}{"A": false, "B": true, "C": true},
				expression: "({{.A}} + {{.B}})",
			},
			expectedError: fmt.Errorf(constants.ErrorEvaluatingExpression),
		},
		{
			name: "Case 3",
			args: args{
				data:       map[string]interface{}{"A": false, "B": true, "C": true},
				expression: "{{.A}} && {{.B",
			},
			expectedError: fmt.Errorf(constants.ErrorInvalidExpression),
		},
		{
			name: "Case 4",
			args: args{
				data:       map[string]interface{}{"A": false, "B": true, "C": true},
				expression: "",
			},
			expectedError: fmt.Errorf(constants.ErrorEmptyExpression),
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			result, err := suite.utils.Evaluate(tc.args.data, tc.args.expression)
			suite.Nil(result)
			suite.Equal(tc.expectedError, err)
		})

	}
}

func (suite *ExpressionEvaluatorTestSuite) TestEvaluate_Success() {
	type args struct {
		tmpResult  map[string]interface{}
		expression string
	}
	tests := []struct {
		name        string
		args        args
		expectation interface{}
	}{
		{
			name: "Boolean Result",
			args: args{
				tmpResult: map[string]interface{}{
					"A": true,
					"B": false,
					"C": true,
				},
				expression: "(({{.A}} && {{.B}}) || {{.C}})",
			},
			expectation: true,
		},
		{
			name: "Integer Summation Result",
			args: args{
				tmpResult: map[string]interface{}{
					"A": 11,
					"B": 12,
					"C": 13,
				},
				expression: "(({{.A}} + {{.B}}) * {{.C}})",
			},
			expectation: float64(299),
		},
		{
			name: "String Concat Result",
			args: args{
				tmpResult: map[string]interface{}{
					"A": "ABC",
					"B": "DEF",
					"C": "GHI",
				},
				expression: "'{{.A}}' + '->' + '{{.B}}' + '->' +  '{{.C}}'",
			},
			expectation: "ABC->DEF->GHI",
		},
		{
			name: "Conditional Result - PASS",
			args: args{
				tmpResult: map[string]interface{}{
					"A": true,
					"B": false,
					"C": true,
				},
				expression: "(({{.A}} && {{.B}}) || {{.C}}) ? 'PASS' : 'FAIL'",
			},
			expectation: "PASS",
		},
		{
			name: "Conditional Result - FAIL",
			args: args{
				tmpResult: map[string]interface{}{
					"A": true,
					"B": false,
					"C": true,
				},
				expression: "({{.A}} && {{.B}} && {{.C}}) ? 'PASS' : 'FAIL'",
			},
			expectation: "FAIL",
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			result, err := suite.utils.Evaluate(tt.args.tmpResult, tt.args.expression)
			suite.Equal(tt.expectation, result)
			suite.Nil(err)
		})
	}
}

func TestExpressionEvaluatorTestSuite(t *testing.T) {
	suite.Run(t, new(ExpressionEvaluatorTestSuite))
}
