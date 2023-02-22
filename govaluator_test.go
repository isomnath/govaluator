package govaluator

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/dependencies"
	"github.com/isomnath/govaluator/models"
)

type GovaluatorTestSuite struct {
	suite.Suite
	dep        *dependencies.Dependencies
	govaluator Govaluator
}

func (suite *GovaluatorTestSuite) SetupTest() {
	suite.dep = dependencies.Initialize()
	suite.govaluator = Initialize()
}

func (suite *GovaluatorTestSuite) TestGetAllDataTypes() {
	expected := []string{constants.StringValue, constants.StringSlice,
		constants.BooleanValue, constants.BooleanSlice,
		constants.IntegerValue, constants.IntegerSlice, constants.Integer8Value, constants.Integer8Slice, constants.Integer16Value,
		constants.Integer16Slice, constants.Integer32Value, constants.Integer32Slice, constants.Integer64Value, constants.Integer64Slice,
		constants.Float32Value, constants.Float32Slice, constants.Float64Value, constants.Float64Slice,
		constants.TimeValue}
	actual := suite.govaluator.GetAllDataTypes()

	sort.Strings(expected)
	sort.Strings(actual)
	suite.Equal(expected, actual)
}

func (suite *GovaluatorTestSuite) TestGetOperatorsByDataType() {
	expected := []string{constants.Before, constants.After, constants.During}
	actual := suite.govaluator.GetOperatorsByDataType(constants.TimeValue)

	sort.Strings(expected)
	sort.Strings(actual)
	suite.Equal(expected, actual)
}

func (suite *GovaluatorTestSuite) TestGetComparators() {
	expected := []string{constants.StaticComparator, constants.DynamicComparator}
	actual := suite.govaluator.GetComparators()
	sort.Strings(expected)
	sort.Strings(actual)
	suite.Equal(expected, actual)
}

func (suite *GovaluatorTestSuite) TestGetEngines() {
	expected := []string{constants.LinearEngine, constants.LinearWaterfallEngine}
	actual := suite.govaluator.GetEngines()
	sort.Strings(expected)
	sort.Strings(actual)
	suite.Equal(expected, actual)
}

func (suite *GovaluatorTestSuite) TestExecuteOne_LinearRule() {
	data := map[string]interface{}{
		"parent.child_1": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_2": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_3": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
	}
	expectedResult := models.Result{
		Success: true,
		Rule: models.RuleMetadata{
			ID:          "123",
			Name:        "Linear Rule",
			Description: "Description - Linear Rule",
			Engine:      constants.LinearEngine,
		},
		Data: &models.ResultData{
			Remark: "",
			Value:  true,
		},
	}

	result, err := suite.govaluator.ExecuteOne(data, suite.getRule("valid_linear_rule"))
	suite.Equal(expectedResult, result)
	suite.NoError(err)
}

func (suite *GovaluatorTestSuite) TestExecuteOne_LinearWaterfallRule() {
	data := map[string]interface{}{
		"parent.child_1": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_2": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_3": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
	}
	expectedResult := models.Result{
		Success: true,
		Rule: models.RuleMetadata{
			ID:          "456",
			Name:        "Linear Waterfall Rule",
			Description: "Description - Linear Waterfall Rule",
			Engine:      constants.LinearWaterfallEngine,
		},
		Data: &models.ResultData{
			Remark: "",
			Value:  80,
		},
	}

	result, err := suite.govaluator.ExecuteOne(data, suite.getRule("valid_linear_waterfall_rule"))
	suite.Equal(expectedResult, result)
	suite.NoError(err)
}

func (suite *GovaluatorTestSuite) TestExecuteMany() {
	data := map[string]interface{}{
		"parent.child_1": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_2": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
		"parent.child_3": map[string]interface{}{
			"grand_child_1": 10,
			"grand_child_2": 20,
		},
	}
	expectedResults := []models.Result{
		{
			Success: true,
			Rule: models.RuleMetadata{
				ID:          "123",
				Name:        "Linear Rule",
				Description: "Description - Linear Rule",
				Engine:      constants.LinearEngine,
			},
			Data: &models.ResultData{
				Remark: "",
				Value:  true,
			},
		},
		{
			Success: true,
			Rule: models.RuleMetadata{
				ID:          "456",
				Name:        "Linear Waterfall Rule",
				Description: "Description - Linear Waterfall Rule",
				Engine:      constants.LinearWaterfallEngine,
			},
			Data: &models.ResultData{
				Remark: "",
				Value:  80,
			},
		},
	}
	rules := []models.Rule{
		suite.getRule("valid_linear_rule"),
		suite.getRule("valid_linear_waterfall_rule"),
	}

	results, err := suite.govaluator.ExecuteMany(data, rules)
	suite.Equal(expectedResults, results)
	suite.NoError(err)
}

func (suite *GovaluatorTestSuite) getRule(ruleType string) models.Rule {
	rules := map[string]models.Rule{
		"valid_linear_rule": {
			ID:          "123",
			Name:        "Linear Rule",
			Description: "Description - Linear Rule",
			Remark:      "",
			Engine:      constants.LinearEngine,
			Linear: &models.Linear{
				Expression: "(({{.A}} && {{.B}}) || {{.C}})",
				Criteria: []models.Criterion{
					{
						ID:         "1",
						Alias:      "A",
						Comparator: constants.StaticComparator,
						Static: &models.Static{
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.Equals,
							Value:    10,
						},
					},
					{
						ID:         "2",
						Alias:      "B",
						Comparator: constants.DynamicComparator,
						Dynamic: &models.Dynamic{
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.LessThan,
							FieldTwo: "parent.child_1.grand_child_2",
						},
					},
					{
						ID:         "3",
						Alias:      "C",
						Comparator: constants.StaticComparator,
						Static: &models.Static{
							FieldOne: "parent.child_1.grand_child_3",
							Operator: constants.GreaterThan,
							Value:    20,
						},
					},
				},
			},
		},
		"valid_linear_waterfall_rule": {
			ID:          "456",
			Name:        "Linear Waterfall Rule",
			Description: "Description - Linear Waterfall Rule",
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
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.Between,
							Value:    []int{0, 50},
						},
						TruthyValue: 80,
						FalseyValue: 0,
					},
					{
						ID:         "2",
						Alias:      "B",
						Comparator: constants.StaticComparator,
						Static: &models.Static{
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.Between,
							Value:    []int{50, 70},
						},
						TruthyValue: 90,
						FalseyValue: 0,
					},
					{
						ID:         "3",
						Alias:      "C",
						Comparator: constants.StaticComparator,
						Static: &models.Static{
							FieldOne: "parent.child_1.grand_child_1",
							Operator: constants.Between,
							Value:    []int{70, 100},
						},
						TruthyValue: 100,
						FalseyValue: 0,
					},
				},
			},
		},
	}
	return rules[ruleType]
}

func TestGovaluatorTestSuite(t *testing.T) {
	suite.Run(t, new(GovaluatorTestSuite))
}
