package floats

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type FloatSliceTestSuite struct {
	suite.Suite
	fSlice            *FloatSlice
	operatorUtilities operatorUtilities
}

func (suite *FloatSliceTestSuite) SetupSuite() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.fSlice = InitializeSliceOperators(suite.operatorUtilities)
}

func (suite *FloatSliceTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.UnorderedEquals, constants.UnorderedNotEquals,
		constants.OrderedEquals, constants.OrderedNotEquals, constants.SupersetOf, constants.NotSupersetOf,
		constants.SubsetOf, constants.NotSubsetOf, constants.Intersection, constants.NotIntersection,
		constants.ReversedAnyOf, constants.ReversedNoneOf, constants.SizeEquals, constants.SizeNotEquals,
		constants.SizeLessThan, constants.SizeLessThanOrEqualTo, constants.SizeGreaterThan, constants.SizeGreaterThanOrEqualTo,
		constants.AnyLessThan, constants.AnyLessThanOrEqualTo, constants.NoneLessThan, constants.NoneLessThanOrEqualTo,
		constants.AllLessThan, constants.AllLessThanOrEqualTo, constants.AnyGreaterThan, constants.AnyGreaterThanOrEqualTo,
		constants.NoneGreaterThan, constants.NoneGreaterThanOrEqualTo, constants.AllGreaterThan, constants.AllGreaterThanOrEqualTo,
		constants.ReversedBetween, constants.ReversedNotBetween}

	actual := suite.fSlice.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *FloatSliceTestSuite) TestUnOrderedEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, true},
		{[]float64{1, 2, 3}, []float64{3, 2, 1}, true},
		{[]float64{1, 2, 3}, []float64{3, 2, 1, 4}, false},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.UnorderedEquals)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestUnOrderedNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, false},
		{[]float64{1, 2, 3}, []float64{3, 2, 1}, false},
		{[]float64{1, 2, 3}, []float64{3, 2, 1, 4}, true},
		{[]float64{1, 2, 3}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.UnorderedNotEquals)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestOrderedEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, true},
		{[]float64{1, 2, 3}, []float64{3, 2, 1}, false},
		{[]float64{1, 2, 3}, []float64{3, 2, 1, 4}, false},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.OrderedEquals)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestOrderedNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, false},
		{[]float64{1, 2, 3}, []float64{3, 2, 1}, true},
		{[]float64{1, 2, 3}, []float64{3, 2, 1, 4}, true},
		{[]float64{1, 2, 3}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.OrderedNotEquals)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestSuperSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, []float64{1, 2}, false},
		{[]float64{1, 2, 3}, []float64{1, 2}, true},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.SupersetOf)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestNotSuperSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, []float64{1, 2}, true},
		{[]float64{1, 2, 3}, []float64{1, 2}, false},
		{[]float64{1, 2, 3}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.NotSupersetOf)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestSubSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, []float64{1, 2}, true},
		{[]float64{1, 2, 3}, []float64{1, 2}, false},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.SubsetOf)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestNotSubSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, []float64{1, 2}, false},
		{[]float64{1, 2, 3}, []float64{1, 2}, true},
		{[]float64{1, 2, 3}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.NotSubsetOf)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestIntersection() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, []float64{1, 2}, true},
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, false},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.Intersection)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestNotIntersection() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, []float64{1, 2}, false},
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, true},
		{[]float64{1, 2, 3}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.NotIntersection)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestReversedAnyOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1, 2}, 2.0, true},
		{[]float64{1, 2, 3}, 4.0, false},
		{[]float64{1, 2, 3}, "12313", false},
	}

	executor := suite.fSlice.GetExecutor(constants.ReversedAnyOf)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestReversedNoneOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1, 2}, 2.0, false},
		{[]float64{1, 2, 3}, 4.0, true},
		{[]float64{1, 2, 3}, "12313", true},
	}

	executor := suite.fSlice.GetExecutor(constants.ReversedNoneOf)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestSizeEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1, 2}, 2, true},
		{[]float64{1, 2}, 2, true},
		{[]float64{1, 2, 3}, 2, false},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.SizeEquals)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestSizeNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1, 2}, 2, false},
		{[]float64{1, 2}, 2, false},
		{[]float64{1, 2, 3}, 2, true},
		{[]float64{1, 2, 3}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.SizeNotEquals)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestSizeLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, 2, true},
		{[]float64{1, 2, 3}, 2, false},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.SizeLessThan)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestSizeLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, 2, true},
		{[]float64{1, 2}, 2, true},
		{[]float64{1, 2, 3}, 2, false},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.SizeLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestSizeGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, 2, false},
		{[]float64{1, 2}, 2, false},
		{[]float64{1, 2, 3}, 2, true},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.SizeGreaterThan)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestSizeGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{1}, 2, false},
		{[]float64{1, 2}, 2, true},
		{[]float64{1, 2, 3}, 2, true},
		{[]float64{1, 2, 3}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.SizeGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestAnyLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, true},
		{[]float64{10, 12}, 11.0, true},
		{[]float64{10, 12}, 10.0, false},
		{[]float64{10, 12}, 9.0, false},
		{[]float64{10, 12}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.AnyLessThan)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestNoneLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, false},
		{[]float64{10, 12}, 11.0, false},
		{[]float64{10, 12}, 10.0, true},
		{[]float64{10, 12}, 9.0, true},
		{[]float64{10, 12}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.NoneLessThan)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestAnyLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, true},
		{[]float64{10, 12}, 11.0, true},
		{[]float64{10, 12}, 10.0, true},
		{[]float64{10, 12}, 9.0, false},
		{[]float64{10, 12}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.AnyLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestNoneLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, false},
		{[]float64{10, 12}, 11.0, false},
		{[]float64{10, 12}, 10.0, false},
		{[]float64{10, 12}, 9.0, true},
		{[]float64{10, 12}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.NoneLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestAllLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, true},
		{[]float64{10, 12}, 11.0, false},
		{[]float64{10, 12}, 10.0, false},
		{[]float64{10, 12}, 9.0, false},
		{[]float64{10, 12}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.AllLessThan)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestAllLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, true},
		{[]float64{10, 12}, 11.0, false},
		{[]float64{10, 12}, 12.0, true},
		{[]float64{10, 12}, 9.0, false},
		{[]float64{10, 12}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.AllLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestAnyGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, false},
		{[]float64{10, 12}, 11.0, true},
		{[]float64{10, 12}, 10.0, true},
		{[]float64{10, 12}, 9.0, true},
		{[]float64{10, 12}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.AnyGreaterThan)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestNoneGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, true},
		{[]float64{10, 12}, 11.0, false},
		{[]float64{10, 12}, 10.0, false},
		{[]float64{10, 12}, 9.0, false},
		{[]float64{10, 12}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.NoneGreaterThan)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestAnyGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, false},
		{[]float64{10, 12}, 11.0, true},
		{[]float64{10, 10}, 10.0, true},
		{[]float64{10, 12}, 9.0, true},
		{[]float64{10, 12}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.AnyGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestNoneGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, true},
		{[]float64{10, 12}, 11.0, false},
		{[]float64{10, 10}, 10.0, false},
		{[]float64{10, 12}, 9.0, false},
		{[]float64{10, 12}, "invalid value", true},
	}

	executor := suite.fSlice.GetExecutor(constants.NoneGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestAllGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, false},
		{[]float64{10, 12}, 9.0, true},
		{[]float64{10, 10}, 10.0, false},
		{[]float64{10, 12}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.AllGreaterThan)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestAllGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{10, 12}, 13.0, false},
		{[]float64{10, 12}, 9.0, true},
		{[]float64{10, 10}, 10.0, true},
		{[]float64{10, 12}, "invalid value", false},
	}

	executor := suite.fSlice.GetExecutor(constants.AllGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestReverseBetween() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{6, 8}, 7.0, true},
		{[]float64{8, 10}, 8.0, false},
		{[]float64{8, 10}, 10.0, false},
		{[]float64{11, 12}, 8.0, false},
		{[]float64{11, 12, 13}, 8.0, false},
		{[]float64{11}, 8.0, false},
		{10, "test", false},
	}

	executor := suite.fSlice.GetExecutor(constants.ReversedBetween)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatSliceTestSuite) TestReverseNotBetween() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]float64{6, 8}, 7.0, false},
		{[]float64{8, 10}, 8.0, true},
		{[]float64{8, 10}, 10.0, true},
		{[]float64{11, 12}, 8.0, true},
		{[]float64{11, 12, 13}, 8.0, true},
		{[]float64{11}, 8.0, true},
		{10, "test", true},
	}

	executor := suite.fSlice.GetExecutor(constants.ReversedNotBetween)

	for _, testData := range testDataSet {
		result := suite.fSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func TestIntegerSliceTestSuite(t *testing.T) {
	suite.Run(t, new(FloatSliceTestSuite))
}
