package strings

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type StringSliceTestSuite struct {
	suite.Suite
	sSlice            *StringSlice
	operatorUtilities operatorUtilities
}

func (suite *StringSliceTestSuite) SetupSuite() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.sSlice = InitializeSliceOperators(suite.operatorUtilities)
}

func (suite *StringSliceTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.UnorderedEquals, constants.UnorderedNotEquals, constants.OrderedEquals, constants.OrderedNotEquals,
		constants.SupersetOf, constants.NotSupersetOf, constants.SubsetOf, constants.NotSubsetOf,
		constants.Intersection, constants.NotIntersection, constants.ReversedAnyOf, constants.ReversedNoneOf,
		constants.SizeEquals, constants.SizeNotEquals, constants.SizeLessThan, constants.SizeLessThanOrEqualTo,
		constants.SizeGreaterThan, constants.SizeGreaterThanOrEqualTo,
		constants.AnyRegexMatch, constants.NoneRegexMatch, constants.AllRegexMatch}

	actual := suite.sSlice.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *StringSliceTestSuite) TestUnOrderedEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def", "ghi"}, []string{"abc", "ghi", "def"}, true},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "ghi", "def", "u87jhgsd"}, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.UnorderedEquals)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestUnOrderedNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def", "ghi"}, []string{"abc", "ghi", "def"}, false},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "ghi", "def", "u87jhgsd"}, true},
		{[]bool{true, false, true}, "invalid value", true},
	}

	executor := suite.sSlice.GetExecutor(constants.UnorderedNotEquals)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestOrderedEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def", "ghi"}, []string{"abc", "def", "ghi"}, true},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "ghi", "def"}, false},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "def", "ghi", "u87jhgsd"}, false},
		{[]bool{true, false, true}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.OrderedEquals)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestOrderedNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def", "ghi"}, []string{"abc", "def", "ghi"}, false},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "ghi", "def"}, true},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "def", "ghi", "u87jhgsd"}, true},
		{[]bool{true, false, true}, "invalid value", true},
	}

	executor := suite.sSlice.GetExecutor(constants.OrderedNotEquals)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestSuperSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc"}, []string{"abc", "def"}, false},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "def"}, true},
		{[]string{"abc", "def", "ghi"}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.SupersetOf)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestNotSuperSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc"}, []string{"abc", "def"}, true},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "def"}, false},
		{[]string{"abc", "def", "ghi"}, "invalid value", true},
	}

	executor := suite.sSlice.GetExecutor(constants.NotSupersetOf)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestSubSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc"}, []string{"abc", "def"}, true},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "def"}, false},
		{[]string{"abc", "def", "ghi"}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.SubsetOf)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestNotSubSetOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc"}, []string{"abc", "def"}, false},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "def"}, true},
		{[]string{"abc", "def", "ghi"}, "invalid value", true},
	}

	executor := suite.sSlice.GetExecutor(constants.NotSubsetOf)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestIntersection() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc"}, []string{"abc", "def"}, true},
		{[]string{"abc", "def", "ghi"}, []string{"asjhgd", "ajhsdgad"}, false},
		{[]string{"abc", "def", "ghi"}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.Intersection)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestNotIntersection() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc"}, []string{"abc", "def"}, false},
		{[]string{"abc", "def", "ghi"}, []string{"asjhgd", "ajhsdgad"}, true},
		{[]string{"abc", "def", "ghi"}, "invalid value", true},
	}

	executor := suite.sSlice.GetExecutor(constants.NotIntersection)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestReversedAnyOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def"}, "def", true},
		{[]string{"abc", "def", "ghi"}, "asjhgd", false},
		{[]string{"abc", "def", "ghi"}, 12313, false},
	}

	executor := suite.sSlice.GetExecutor(constants.ReversedAnyOf)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestReversedNoneOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def"}, "def", false},
		{[]string{"abc", "def", "ghi"}, "asjhgd", true},
		{[]string{"abc", "def", "ghi"}, 12313, true},
	}

	executor := suite.sSlice.GetExecutor(constants.ReversedNoneOf)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestSizeEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def"}, 2, true},
		{[]string{"abc", "def", "ghi"}, 2, false},
		{[]string{"abc", "def", "ghi"}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.SizeEquals)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestSizeNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def"}, 2, false},
		{[]string{"abc", "def", "ghi"}, 2, true},
		{[]string{"abc", "def", "ghi"}, "invalid value", true},
	}

	executor := suite.sSlice.GetExecutor(constants.SizeNotEquals)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestSizeLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc"}, 2, true},
		{[]string{"abc", "def", "ghi"}, 2, false},
		{[]string{"abc", "def", "ghi"}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.SizeLessThan)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestSizeLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc"}, 2, true},
		{[]string{"abc", "def"}, 2, true},
		{[]string{"abc", "def", "ghi"}, 2, false},
		{[]string{"abc", "def", "ghi"}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.SizeLessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestSizeGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def"}, 1, true},
		{[]string{"abc", "def"}, 3, false},
		{[]string{"abc", "def", "ghi"}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.SizeGreaterThan)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestSizeGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc"}, 1, true},
		{[]string{"abc", "def"}, 1, true},
		{[]string{"abc", "def"}, 3, false},
		{[]string{"abc", "def", "ghi"}, "invalid value", false},
	}

	executor := suite.sSlice.GetExecutor(constants.SizeGreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestAnyRegexMatch() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def", "ghi"}, "a", true},
		{[]string{"abc", "def", "ghi"}, "t", false},
		{[]string{"abc", "def", "ghi"}, "*", false},
		{[]string{"abc", "def", "ghi"}, 21231, false},
	}

	executor := suite.sSlice.GetExecutor(constants.AnyRegexMatch)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestNotRegexMatch() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "def", "ghi"}, "a", false},
		{[]string{"abc", "def", "ghi"}, "t", true},
		{[]string{"abc", "def", "ghi"}, "*", true},
		{[]string{"abc", "def", "ghi"}, 124312, true},
	}

	executor := suite.sSlice.GetExecutor(constants.NoneRegexMatch)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *StringSliceTestSuite) TestAllRegexMatch() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{[]string{"abc", "adef", "aghi"}, "a", true},
		{[]string{"abc", "def", "ghi"}, "a", false},
		{[]string{"abc", "def", "ghi"}, "*", false},
		{[]string{"abc", "def", "ghi"}, 21231, false},
	}

	executor := suite.sSlice.GetExecutor(constants.AllRegexMatch)

	for _, testData := range testDataSet {
		result := suite.sSlice.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func TestStringSliceTestSuite(t *testing.T) {
	suite.Run(t, new(StringSliceTestSuite))
}
