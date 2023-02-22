package engines

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/comparators"
	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/engines/exitonstrategies"
	"github.com/isomnath/govaluator/models"
	"github.com/isomnath/govaluator/operators"
	"github.com/isomnath/govaluator/utilities"
)

type LinearWaterfallEngineTestSuite struct {
	suite.Suite
	utils             utils
	strategyManager   strategyManger
	comparisonManager comparisonManager
	engine            Engines
}

func (suite *LinearWaterfallEngineTestSuite) SetupTest() {
	suite.utils = utilities.InitializeUtilities()
	operatorUtilities := utilities.InitializeOperatorUtilities()
	operationManager := operators.InitializeOperators(operatorUtilities)
	suite.strategyManager = exitonstrategies.InitializeStrategyManager()
	suite.comparisonManager = comparators.InitializeComparators(operationManager)
	suite.engine = initializeLinearWaterfallEngine(suite.utils, suite.strategyManager, suite.comparisonManager)
}

func (suite *LinearWaterfallEngineTestSuite) TestLinearWaterfallEngineRemarkError() {
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
		Engine:      constants.LinearWaterfallEngine,
		LinearWaterfall: &models.LinearWaterfall{
			ExitOn: constants.FirstTrue,
			Criteria: []models.Criterion{
				{
					ID:         "1",
					Alias:      "A",
					Comparator: constants.StaticComparator,
					Static: &models.Static{
						FieldOne: "parent.child_1",
						Operator: constants.Equals,
						Value:    11,
					},
					TruthyValue: "True-ABC",
					FalseyValue: "False-ABC",
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
					TruthyValue: "True-DEF",
					FalseyValue: "False-DEF",
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
					TruthyValue: "True-GHI",
					FalseyValue: "False-GHI",
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
			Engine:      constants.LinearWaterfallEngine,
		},
		Errors: []error{fmt.Errorf("failed to parse expression: %s with error: template: :1: unclosed action", rule.Remark)},
	}
	result := suite.engine.Execute(data, rule)
	suite.Equal(expectedResult, result)
}

func (suite *LinearWaterfallEngineTestSuite) TestLinearWaterfallEngineExecuteSuccessWithDefaultReturnValue() {
	data := map[string]interface{}{
		"parent.child_1": "10",
		"parent.child_2": "9",
		"parent.child_3": []string{"8", "14"},
	}
	rule := models.Rule{
		ID:          "1231",
		Name:        "Test - Linear Rule",
		Description: "Test - Linear Rule - Description",
		Remark:      "",
		Engine:      constants.LinearWaterfallEngine,
		LinearWaterfall: &models.LinearWaterfall{
			DefaultReturnValue: "Default Value",
			ExitOn:             constants.FirstTrue,
			Criteria: []models.Criterion{
				{
					ID:         "1",
					Alias:      "A",
					Comparator: constants.StaticComparator,
					Static: &models.Static{
						FieldOne: "parent.child_1",
						Operator: constants.Equals,
						Value:    11,
					},
					TruthyValue: "True-ABC",
					FalseyValue: "False-ABC",
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
					TruthyValue: "True-DEF",
					FalseyValue: "False-DEF",
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
					TruthyValue: "True-GHI",
					FalseyValue: "False-GHI",
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
			Engine:      constants.LinearWaterfallEngine,
		},
		Data: &models.ResultData{
			Remark: "",
			Value:  "Default Value",
		},
	}
	result := suite.engine.Execute(data, rule)
	suite.Equal(expectedResult, result)
}

func (suite *LinearWaterfallEngineTestSuite) TestLinearWaterfallEngineExecuteSuccess() {
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
		Engine:      constants.LinearWaterfallEngine,
		LinearWaterfall: &models.LinearWaterfall{
			ExitOn: constants.FirstTrue,
			Criteria: []models.Criterion{
				{
					ID:         "1",
					Alias:      "A",
					Comparator: constants.StaticComparator,
					Static: &models.Static{
						FieldOne: "parent.child_1",
						Operator: constants.Equals,
						Value:    11,
					},
					TruthyValue: "True-ABC",
					FalseyValue: "False-ABC",
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
					TruthyValue: "True-DEF",
					FalseyValue: "False-DEF",
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
					TruthyValue: "True-GHI",
					FalseyValue: "False-GHI",
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
			Engine:      constants.LinearWaterfallEngine,
		},
		Data: &models.ResultData{
			Remark: "",
			Value:  "True-DEF",
		},
	}
	result := suite.engine.Execute(data, rule)
	suite.Equal(expectedResult, result)
}

func TestLinearWaterfallEngineTestSuite(t *testing.T) {
	suite.Run(t, new(LinearWaterfallEngineTestSuite))
}
