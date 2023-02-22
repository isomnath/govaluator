package booleans

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type BooleanSliceTestSuite struct {
	suite.Suite
	bSlice            *BooleanSlice
	operatorUtilities operatorUtilities
}

func (suite *BooleanSliceTestSuite) SetupSuite() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.bSlice = InitializeSliceOperators(suite.operatorUtilities)
}

func (suite *BooleanSliceTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.UnorderedEquals, constants.UnorderedNotEquals,
		constants.OrderedEquals, constants.OrderedNotEquals, constants.SupersetOf, constants.NotSupersetOf,
		constants.SubsetOf, constants.NotSubsetOf, constants.Intersection, constants.NotIntersection,
		constants.ReversedAnyOf, constants.ReversedNoneOf, constants.SizeEquals, constants.SizeNotEquals,
		constants.SizeLessThan, constants.SizeLessThanOrEqualTo, constants.SizeGreaterThan, constants.SizeGreaterThanOrEqualTo}

	actual := suite.bSlice.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *BooleanSliceTestSuite) TestUnOrderedEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false, true}, []bool{true, false, true}, true},
		{[]bool{true, false, true}, []bool{true, true, false}, true},
		{[]bool{true, false, true}, []bool{true, false}, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.UnorderedEquals)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestUnOrderedNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false, true}, []bool{true, false, true}, false},
		{[]bool{true, false, true}, []bool{true, true, false}, false},
		{[]bool{true, false, true}, []bool{true, false}, true},
		{[]bool{true, false, true}, "invalid value", true},
	}

	executor := suite.bSlice.GetExecutor(constants.UnorderedNotEquals)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestOrderedEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false, true}, []bool{true, false, true}, true},
		{[]bool{true, false, true}, []bool{true, true, false}, false},
		{[]bool{true, false, true}, []bool{true, false}, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.OrderedEquals)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestOrderedNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false, true}, []bool{true, false, true}, false},
		{[]bool{true, false, true}, []bool{true, true, false}, true},
		{[]bool{true, false, true}, []bool{true, false}, true},
		{[]bool{true, false, true}, "invalid value", true},
	}

	executor := suite.bSlice.GetExecutor(constants.OrderedNotEquals)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestSuperSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true}, []bool{true, false}, false},
		{[]bool{true, false}, []bool{true}, true},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.SupersetOf)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestNotSuperSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true}, []bool{true, false}, true},
		{[]bool{true, false}, []bool{true}, false},
		{[]bool{true, false, true}, "invalid value", true},
	}

	executor := suite.bSlice.GetExecutor(constants.NotSupersetOf)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestSubSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true}, []bool{true, false}, true},
		{[]bool{true, false}, []bool{true}, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.SubsetOf)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestNotSubSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true}, []bool{true, false}, false},
		{[]bool{true, false}, []bool{true}, true},
		{[]bool{true, false, true}, "invalid value", true},
	}

	executor := suite.bSlice.GetExecutor(constants.NotSubsetOf)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestIntersection() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true}, []bool{true, false}, true},
		{[]bool{true, false}, []bool{}, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.Intersection)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestNotIntersection() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true}, []bool{true, false}, false},
		{[]bool{true, false}, []bool{}, true},
		{[]bool{true, false, true}, "invalid value", true},
	}

	executor := suite.bSlice.GetExecutor(constants.NotIntersection)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestReversedAnyOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false}, true, true},
		{[]bool{true}, false, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.ReversedAnyOf)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestReversedNoneOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false}, true, false},
		{[]bool{true}, false, true},
		{[]bool{true, false, true}, "invalid value", true},
	}

	executor := suite.bSlice.GetExecutor(constants.ReversedNoneOf)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestSizeEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false}, 2, true},
		{[]bool{true, false, true}, 2, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.SizeEquals)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestSizeNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false}, 2, false},
		{[]bool{true, false, true}, 2, true},
		{[]bool{true, false, true}, "invalid value", true},
	}

	executor := suite.bSlice.GetExecutor(constants.SizeNotEquals)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestSizeLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false}, 3, true},
		{[]bool{true, false, true}, 2, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.SizeLessThan)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestSizeLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false}, 3, true},
		{[]bool{true, false, true}, 3, true},
		{[]bool{true, false, true}, 2, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.SizeLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestSizeGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false, true}, 2, true},
		{[]bool{true, false, true}, 3, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.SizeGreaterThan)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *BooleanSliceTestSuite) TestSizeGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]bool{true, false, true}, 2, true},
		{[]bool{true, false, true}, 3, true},
		{[]bool{true, false, true}, 4, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.bSlice.GetExecutor(constants.SizeGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.bSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func TestBooleanSliceTestSuite(t *testing.T) {
	suite.Run(t, new(BooleanSliceTestSuite))
}
