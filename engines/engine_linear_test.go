package engines

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/comparators"
	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/models"
	"github.com/isomnath/govaluator/operators"
	"github.com/isomnath/govaluator/utilities"
)

type LinearEngineTestSuite struct {
	suite.Suite
	utils             utils
	comparisonManager comparisonManager
	engine            Engines
}

func (suite *LinearEngineTestSuite) SetupTest() {
	suite.utils = utilities.InitializeUtilities()
	operatorUtilities := utilities.InitializeOperatorUtilities()
	operationManager := operators.InitializeOperators(operatorUtilities)
	suite.comparisonManager = comparators.InitializeComparators(operationManager)
	suite.engine = initializeLinearEngine(suite.utils, suite.comparisonManager)
}

func (suite *LinearEngineTestSuite) TestLinearEngineExecuteExpressionError() {
	data := map[string]interface{}{
		"parent.child_1": 10,
		"parent.child_2": 9,
		"parent.child_3": []int{8, 14},
	}
	rule := models.Rule{
		ID:          "1231",
		Name:        "Test - Linear Rule",
		Description: "Test - Linear Rule - Description",
		Remark:      "",
		Engine:      constants.LinearEngine,
		Linear: &models.Linear{
			Expression: "(({{.A}} || {{.B}}) && {{.C)",
			Criteria: []models.Criterion{
				{
					ID:         "1",
					Alias:      "A",
					Comparator: constants.StaticComparator,
					Static: &models.Static{
						FieldOne: "parent.child_1",
						Operator: constants.Equals,
						Value:    10,
					},
				},
				{
					ID:         "2",
					Alias:      "B",
					Comparator: constants.DynamicComparator,
					Dynamic: &models.Dynamic{
						FieldOne: "parent.child_1",
						Operator: constants.GreaterThan,
						FieldTwo: "parent.child_2",
					},
				},
				{
					ID:         "3",
					Alias:      "C",
					Comparator: constants.DynamicComparator,
					Dynamic: &models.Dynamic{
						FieldOne: "parent.child_2",
						Operator: constants.Between,
						FieldTwo: "parent.child_3",
					},
				},
			},
		},
	}

	expectedResult := models.Result{
		Success: false,
		Rule: models.RuleMetadata{
			ID:          "1231",
			Name:        "Test - Linear Rule",
			Description: "Test - Linear Rule - Description",
			Engine:      constants.LinearEngine,
		},
		Errors: []error{fmt.Errorf(constants.ErrorInvalidExpression)},
	}
	result := suite.engine.Execute(data, rule)
	suite.Equal(expectedResult, result)
}

func (suite *LinearEngineTestSuite) TestLinearEngineExecuteRemarkError() {
	data := map[string]interface{}{
		"parent.child_1": 10,
		"parent.child_2": 9,
		"parent.child_3": []int{8, 14},
	}
	rule := models.Rule{
		ID:          "1231",
		Name:        "Test - Linear Rule",
		Description: "Test - Linear Rule - Description",
		Remark:      "ID: {{.rule.id}}, Name: {{.rule.name}}, Description: {{.rule.description}}, Engine: {{.rule.engine}}, Data: {{.data}}, Result: {{.result_value",
		Engine:      constants.LinearEngine,
		Linear: &models.Linear{
			Expression: "(({{.A}} || {{.B}}) && {{.C}})",
			Criteria: []models.Criterion{
				{
					ID:         "1",
					Alias:      "A",
					Comparator: constants.StaticComparator,
					Static: &models.Static{
						FieldOne: "parent.child_1",
						Operator: constants.Equals,
						Value:    10,
					},
				},
				{
					ID:         "2",
					Alias:      "B",
					Comparator: constants.DynamicComparator,
					Dynamic: &models.Dynamic{
						FieldOne: "parent.child_1",
						Operator: constants.GreaterThan,
						FieldTwo: "parent.child_2",
					},
				},
				{
					ID:         "3",
					Alias:      "C",
					Comparator: constants.DynamicComparator,
					Dynamic: &models.Dynamic{
						FieldOne: "parent.child_2",
						Operator: constants.Between,
						FieldTwo: "parent.child_3",
					},
				},
			},
		},
	}

	expectedResult := models.Result{
		Success: false,
		Rule: models.RuleMetadata{
			ID:          "1231",
			Name:        "Test - Linear Rule",
			Description: "Test - Linear Rule - Description",
			Engine:      constants.LinearEngine,
		},
		Errors: []error{fmt.Errorf("failed to parse expression: %s with error: template: :1: unclosed action", rule.Remark)},
	}
	result := suite.engine.Execute(data, rule)
	suite.Equal(expectedResult, result)
}

func (suite *LinearEngineTestSuite) TestLinearEngineExecuteSuccess() {
	data := map[string]interface{}{
		"parent.child_1": 10,
		"parent.child_2": 9,
		"parent.child_3": []int{8, 14},
	}
	rule := models.Rule{
		ID:          "1231",
		Name:        "Test - Linear Rule",
		Description: "Test - Linear Rule - Description",
		Remark:      "",
		Engine:      constants.LinearEngine,
		Linear: &models.Linear{
			Expression: "(({{.A}} || {{.B}}) && {{.C}})",
			Criteria: []models.Criterion{
				{
					ID:         "1",
					Alias:      "A",
					Comparator: constants.StaticComparator,
					Static: &models.Static{
						FieldOne: "parent.child_1",
						Operator: constants.Equals,
						Value:    10,
					},
				},
				{
					ID:         "2",
					Alias:      "B",
					Comparator: constants.DynamicComparator,
					Dynamic: &models.Dynamic{
						FieldOne: "parent.child_1",
						Operator: constants.GreaterThan,
						FieldTwo: "parent.child_2",
					},
				},
				{
					ID:         "3",
					Alias:      "C",
					Comparator: constants.DynamicComparator,
					Dynamic: &models.Dynamic{
						FieldOne: "parent.child_2",
						Operator: constants.Between,
						FieldTwo: "parent.child_3",
					},
				},
			},
		},
	}

	expectedResult := models.Result{
		Success: true,
		Rule: models.RuleMetadata{
			ID:          "1231",
			Name:        "Test - Linear Rule",
			Description: "Test - Linear Rule - Description",
			Engine:      constants.LinearEngine,
		},
		Data: &models.ResultData{
			Remark: "",
			Value:  true,
		},
	}
	result := suite.engine.Execute(data, rule)
	suite.Equal(expectedResult, result)
}

func TestLinearEngineTestSuite(t *testing.T) {
	suite.Run(t, new(LinearEngineTestSuite))
}
