package nils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
)

type NilValueTestSuite struct {
	suite.Suite
	nilValue *NilValue
}

func (suite *NilValueTestSuite) SetupSuite() {
	suite.nilValue = InitializeNilOperators()
}

func (suite *NilValueTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.Equals, constants.NotEquals,
		constants.LessThan, constants.LessThanOrEqualTo,
		constants.LessThanAny, constants.LessThanOrEqualToAny,
		constants.LessThanNone, constants.LessThanOrEqualToNone,
		constants.LessThanAll, constants.LessThanOrEqualToAll,
		constants.GreaterThan, constants.GreaterThanOrEqualTo,
		constants.GreaterThanAny, constants.GreaterThanOrEqualToAny,
		constants.GreaterThanNone, constants.GreaterThanOrEqualToNone,
		constants.GreaterThanAll, constants.GreaterThanOrEqualToAll,
		constants.Between, constants.NotBetween,
		constants.AnyOf, constants.NoneOf,
		constants.RegexMatch, constants.RegexNotMatch,
		constants.UnorderedEquals, constants.UnorderedNotEquals,
		constants.OrderedEquals, constants.OrderedNotEquals,
		constants.AnyLessThan, constants.AnyLessThanOrEqualTo,
		constants.NoneLessThan, constants.NoneLessThanOrEqualTo,
		constants.AllLessThan, constants.AllLessThanOrEqualTo,
		constants.AnyGreaterThan, constants.AnyGreaterThanOrEqualTo,
		constants.NoneGreaterThan, constants.NoneGreaterThanOrEqualTo,
		constants.AllGreaterThan, constants.AllGreaterThanOrEqualTo,
		constants.ReversedBetween, constants.ReversedNotBetween,
		constants.SupersetOf, constants.NotSupersetOf,
		constants.SubsetOf, constants.NotSubsetOf,
		constants.Intersection, constants.NotIntersection,
		constants.ReversedAnyOf, constants.ReversedNoneOf,
		constants.SizeEquals, constants.SizeNotEquals,
		constants.SizeLessThan, constants.SizeLessThanOrEqualTo,
		constants.SizeGreaterThan, constants.SizeGreaterThanOrEqualTo,
		constants.AnyRegexMatch, constants.NoneRegexMatch, constants.AllRegexMatch,
		constants.Before, constants.After, constants.During}

	actual := suite.nilValue.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *NilValueTestSuite) TestEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{nil, nil, true},
		{nil, 8.0, false},
	}

	executor := suite.nilValue.GetExecutor(constants.Equals)

	for _, testData := range testDataSet {
		result := suite.nilValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *NilValueTestSuite) TestNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{nil, nil, false},
		{nil, "8.0", true},
	}

	executor := suite.nilValue.GetExecutor(constants.NotEquals)

	for _, testData := range testDataSet {
		result := suite.nilValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *NilValueTestSuite) TestDefaultFalse() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{nil, nil, false},
		{nil, []string{"8.0", "10.0"}, false},
	}

	executor := suite.nilValue.GetExecutor(constants.AnyOf)

	for _, testData := range testDataSet {
		result := suite.nilValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func TestNilValueTestSuite(t *testing.T) {
	suite.Run(t, new(NilValueTestSuite))
}
