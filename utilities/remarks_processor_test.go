package utilities

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/models"
)

type ProcessRemarkTestSuite struct {
	suite.Suite
	utils *Utilities
}

func (suite *ProcessRemarkTestSuite) SetupTest() {
	suite.utils = InitializeUtilities()
}

func (suite *ProcessRemarkTestSuite) TestProcessRemark() {
	data := map[string]interface{}{
		"parent.child_1": 10,
		"parent.child_2": 9,
		"parent.child_3": []int{8, 14},
	}

	ruleWithRemark := models.Rule{
		ID:          "1231",
		Name:        "Test - Linear Waterfall Rule",
		Description: "Test - Linear Waterfall Rule - Description",
		Remark:      "ID: {{.rule.id}}, Name: {{.rule.name}}, Description: {{.rule.description}}, Engine: {{.rule.engine}}, Data: {{.data}}, Result: {{.result_value}}",
		Engine:      constants.LinearWaterfallEngine,
	}

	ruleWithoutRemark := models.Rule{
		ID:          "1232",
		Name:        "Test - Linear Waterfall Rule",
		Description: "Test - Linear Waterfall Rule - Description",
		Remark:      "",
		Engine:      constants.LinearWaterfallEngine,
	}

	ruleWithInvalidRemark := models.Rule{
		ID:          "1233",
		Name:        "Test - Linear Waterfall Rule",
		Description: "Test - Linear Waterfall Rule - Description",
		Remark:      "ID: {{.rule.id}}, Name: {{.rule.name",
		Engine:      constants.LinearWaterfallEngine,
	}

	type args struct {
		data        map[string]interface{}
		rule        models.Rule
		resultValue interface{}
	}

	testCases := []struct {
		name           string
		args           args
		expectedRemark string
		expectedError  error
	}{
		{
			name: "Successful Processing",
			args: args{
				data:        data,
				rule:        ruleWithRemark,
				resultValue: 100,
			},
			expectedRemark: "ID: 1231, Name: Test - Linear Waterfall Rule, Description: Test - Linear Waterfall Rule - Description, Engine: LINEAR_WATERFALL, Data: map[parent:map[child_1:10 child_2:9 child_3:[8 14]]], Result: 100",
			expectedError:  nil,
		},
		{
			name: "Successful Processing - Empty Remark",
			args: args{
				data:        data,
				rule:        ruleWithoutRemark,
				resultValue: 1001,
			},
			expectedRemark: "",
			expectedError:  nil,
		},
		{
			name: "Successful Processing - Error Remark",
			args: args{
				data:        data,
				rule:        ruleWithInvalidRemark,
				resultValue: 1002,
			},
			expectedRemark: "",
			expectedError: fmt.Errorf("failed to parse expression: ID: {{.rule.id}}, Name: {{.rule.name with error: %v",
				errors.New("template: :1: unclosed action")),
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			remark, err := suite.utils.ProcessRemark(tc.args.data, tc.args.rule, tc.args.resultValue)
			suite.Equal(tc.expectedRemark, remark)
			suite.Equal(tc.expectedError, err)
		})
	}
}

func TestProcessRemarkTestSuite(t *testing.T) {
	suite.Run(t, new(ProcessRemarkTestSuite))
}
