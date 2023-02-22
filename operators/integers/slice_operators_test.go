package integers

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type IntegerSliceTestSuite struct {
	suite.Suite
	iSlice            *IntegerSlice
	operatorUtilities operatorUtilities
}

func (suite *IntegerSliceTestSuite) SetupSuite() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.iSlice = InitializeSliceOperators(suite.operatorUtilities)
}

func (suite *IntegerSliceTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.UnorderedEquals, constants.UnorderedNotEquals,
		constants.OrderedEquals, constants.OrderedNotEquals, constants.SupersetOf, constants.NotSupersetOf,
		constants.SubsetOf, constants.NotSubsetOf, constants.Intersection, constants.NotIntersection,
		constants.ReversedAnyOf, constants.ReversedNoneOf, constants.SizeEquals, constants.SizeNotEquals,
		constants.SizeLessThan, constants.SizeLessThanOrEqualTo, constants.SizeGreaterThan, constants.SizeGreaterThanOrEqualTo,
		constants.AnyLessThan, constants.AnyLessThanOrEqualTo, constants.NoneLessThan, constants.NoneLessThanOrEqualTo,
		constants.AllLessThan, constants.AllLessThanOrEqualTo, constants.AnyGreaterThan, constants.AnyGreaterThanOrEqualTo,
		constants.NoneGreaterThan, constants.NoneGreaterThanOrEqualTo, constants.AllGreaterThan, constants.AllGreaterThanOrEqualTo,
		constants.ReversedBetween, constants.ReversedNotBetween}

	actual := suite.iSlice.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *IntegerSliceTestSuite) TestUnOrderedEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1, 4}, false},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.UnorderedEquals)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestUnOrderedNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, false},
		{[]int{1, 2, 3}, []int{3, 2, 1}, false},
		{[]int{1, 2, 3}, []int{3, 2, 1, 4}, true},
		{[]int{1, 2, 3}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.UnorderedNotEquals)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestOrderedEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, false},
		{[]int{1, 2, 3}, []int{3, 2, 1, 4}, false},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.OrderedEquals)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestOrderedNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, false},
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1, 4}, true},
		{[]int{1, 2, 3}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.OrderedNotEquals)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestSuperSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, []int{1, 2}, false},
		{[]int{1, 2, 3}, []int{1, 2}, true},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.SupersetOf)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestNotSuperSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, []int{1, 2}, true},
		{[]int{1, 2, 3}, []int{1, 2}, false},
		{[]int{1, 2, 3}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.NotSupersetOf)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestSubSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, []int{1, 2}, true},
		{[]int{1, 2, 3}, []int{1, 2}, false},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.SubsetOf)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestNotSubSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, []int{1, 2}, false},
		{[]int{1, 2, 3}, []int{1, 2}, true},
		{[]int{1, 2, 3}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.NotSubsetOf)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestIntersection() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, []int{1, 2}, true},
		{[]int{1, 2, 3}, []int{4, 5, 6}, false},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.Intersection)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestNotIntersection() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, []int{1, 2}, false},
		{[]int{1, 2, 3}, []int{4, 5, 6}, true},
		{[]int{1, 2, 3}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.NotIntersection)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestReversedAnyOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1, 2}, 2, true},
		{[]int{1, 2, 3}, 4, false},
		{[]int{1, 2, 3}, "12313", false},
	}

	executor := suite.iSlice.GetExecutor(constants.ReversedAnyOf)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestReversedNoneOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1, 2}, 2, false},
		{[]int{1, 2, 3}, 4, true},
		{[]int{1, 2, 3}, "12313", true},
	}

	executor := suite.iSlice.GetExecutor(constants.ReversedNoneOf)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestSizeEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1, 2}, 2, true},
		{[]int{1, 2}, 2, true},
		{[]int{1, 2, 3}, 2, false},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.SizeEquals)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestSizeNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1, 2}, 2, false},
		{[]int{1, 2}, 2, false},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.SizeNotEquals)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestSizeLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, 2, true},
		{[]int{1, 2, 3}, 2, false},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.SizeLessThan)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestSizeLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, 2, true},
		{[]int{1, 2}, 2, true},
		{[]int{1, 2, 3}, 2, false},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.SizeLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestSizeGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, 2, false},
		{[]int{1, 2}, 2, false},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.SizeGreaterThan)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestSizeGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{1}, 2, false},
		{[]int{1, 2}, 2, true},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.SizeGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestAnyLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, true},
		{[]int{10, 12}, 11, true},
		{[]int{10, 12}, 10, false},
		{[]int{10, 12}, 9, false},
		{[]int{10, 12}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.AnyLessThan)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestNoneLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, false},
		{[]int{10, 12}, 11, false},
		{[]int{10, 12}, 10, true},
		{[]int{10, 12}, 9, true},
		{[]int{10, 12}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.NoneLessThan)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestAnyLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, true},
		{[]int{10, 12}, 11, true},
		{[]int{10, 12}, 10, true},
		{[]int{10, 12}, 9, false},
		{[]int{10, 12}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.AnyLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestNoneLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, false},
		{[]int{10, 12}, 11, false},
		{[]int{10, 12}, 10, false},
		{[]int{10, 12}, 9, true},
		{[]int{10, 12}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.NoneLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestAllLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, true},
		{[]int{10, 12}, 11, false},
		{[]int{10, 12}, 10, false},
		{[]int{10, 12}, 9, false},
		{[]int{10, 12}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.AllLessThan)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestAllLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, true},
		{[]int{10, 12}, 11, false},
		{[]int{10, 12}, 12, true},
		{[]int{10, 12}, 9, false},
		{[]int{10, 12}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.AllLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestAnyGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, false},
		{[]int{10, 12}, 11, true},
		{[]int{10, 12}, 10, true},
		{[]int{10, 12}, 9, true},
		{[]int{10, 12}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.AnyGreaterThan)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestNoneGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, true},
		{[]int{10, 12}, 11, false},
		{[]int{10, 12}, 10, false},
		{[]int{10, 12}, 9, false},
		{[]int{10, 12}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.NoneGreaterThan)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestAnyGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, false},
		{[]int{10, 12}, 11, true},
		{[]int{10, 10}, 10, true},
		{[]int{10, 12}, 9, true},
		{[]int{10, 12}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.AnyGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestNoneGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, true},
		{[]int{10, 12}, 11, false},
		{[]int{10, 10}, 10, false},
		{[]int{10, 12}, 9, false},
		{[]int{10, 12}, "invalid value", true},
	}

	executor := suite.iSlice.GetExecutor(constants.NoneGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestAllGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, false},
		{[]int{10, 12}, 9, true},
		{[]int{10, 10}, 10, false},
		{[]int{10, 12}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.AllGreaterThan)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestAllGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{10, 12}, 13, false},
		{[]int{10, 12}, 9, true},
		{[]int{10, 10}, 10, true},
		{[]int{10, 12}, "invalid value", false},
	}

	executor := suite.iSlice.GetExecutor(constants.AllGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestReverseBetween() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{6, 8}, 7, true},
		{[]int{8, 10}, 8, false},
		{[]int{8, 10}, 10, false},
		{[]int{11, 12}, 8, false},
		{[]int{11, 12, 13}, 8, false},
		{[]int{11}, 8, false},
		{10, "test", false},
	}

	executor := suite.iSlice.GetExecutor(constants.ReversedBetween)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerSliceTestSuite) TestReverseNotBetween() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]int{6, 8}, 7, false},
		{[]int{8, 10}, 8, true},
		{[]int{8, 10}, 10, true},
		{[]int{11, 12}, 8, true},
		{[]int{11, 12, 13}, 8, true},
		{[]int{11}, 8, true},
		{10, "test", true},
	}

	executor := suite.iSlice.GetExecutor(constants.ReversedNotBetween)

	for _, testData := range testDataSet {
		result := suite.iSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func TestIntegerSliceTestSuite(t *testing.T) {
	suite.Run(t, new(IntegerSliceTestSuite))
}
