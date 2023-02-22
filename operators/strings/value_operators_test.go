package strings

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type StringValueTestSuite struct {
	suite.Suite
	sValue            *StringValue
	operatorUtilities operatorUtilities
}

func (suite *StringValueTestSuite) SetupSuite() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.sValue = InitializeValueOperators(suite.operatorUtilities)
}

func (suite *StringValueTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.Equals, constants.NotEquals, constants.AnyOf, constants.NoneOf,
		constants.RegexMatch, constants.RegexNotMatch}

	actual := suite.sValue.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *StringValueTestSuite) TestEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{"test", "test", true},
		{"test-one", "test-two", false},
		{"test", 123113, false},
	}

	executor := suite.sValue.GetExecutor(constants.Equals)

	for _, testData := range testDataSet {
		result := suite.sValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringValueTestSuite) TestNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{"test", "test", false},
		{"test-one", "test-two", true},
		{"test", 123113, true},
	}

	executor := suite.sValue.GetExecutor(constants.NotEquals)

	for _, testData := range testDataSet {
		result := suite.sValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringValueTestSuite) TestAnyOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{"test", []string{"test-one", "test-two"}, false},
		{"test-one", []string{"test-one", "test-two"}, true},
		{"test", 123113, false},
	}

	executor := suite.sValue.GetExecutor(constants.AnyOf)

	for _, testData := range testDataSet {
		result := suite.sValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringValueTestSuite) TestNoneOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{"test", []string{"test-one", "test-two"}, true},
		{"test-one", []string{"test-one", "test-two"}, false},
		{"test", 123113, true},
	}

	executor := suite.sValue.GetExecutor(constants.NoneOf)

	for _, testData := range testDataSet {
		result := suite.sValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringValueTestSuite) TestRegexMatch() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{"test", "14", false},
		{"test-one", "te", true},
		{"test", 123113, false},
		{"test", "*", false},
	}

	executor := suite.sValue.GetExecutor(constants.RegexMatch)

	for _, testData := range testDataSet {
		result := suite.sValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringValueTestSuite) TestNotRegexMatch() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{"test", "14", true},
		{"test-one", "te", false},
		{"test", 123113, true},
	}

	executor := suite.sValue.GetExecutor(constants.RegexNotMatch)

	for _, testData := range testDataSet {
		result := suite.sValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func TestStringValueTestSuite(t *testing.T) {
	suite.Run(t, new(StringValueTestSuite))
}
