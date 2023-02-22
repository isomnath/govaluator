package integers

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type IntegerValueTestSuite struct {
	suite.Suite
	iValue            *IntegerValue
	operatorUtilities operatorUtilities
}

func (suite *IntegerValueTestSuite) SetupSuite() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.iValue = InitializeValueOperators(suite.operatorUtilities)
}

func (suite *IntegerValueTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.Equals, constants.NotEquals,
		constants.LessThan, constants.LessThanOrEqualTo, constants.LessThanAny, constants.LessThanOrEqualToAny, constants.LessThanNone,
		constants.LessThanOrEqualToNone, constants.LessThanAll, constants.LessThanOrEqualToAll, constants.GreaterThan,
		constants.GreaterThanOrEqualTo, constants.GreaterThanAny, constants.GreaterThanOrEqualToAny, constants.GreaterThanNone,
		constants.GreaterThanOrEqualToNone, constants.GreaterThanAll, constants.GreaterThanOrEqualToAll, constants.Between,
		constants.NotBetween, constants.AnyOf, constants.NoneOf}

	actual := suite.iValue.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *IntegerValueTestSuite) TestEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{1231223, 1231223, true},
		{1231223, 1231313, false},
		{1231223, "123113", false},
	}

	executor := suite.iValue.GetExecutor(constants.Equals)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{1231223, 1231223, false},
		{1231223, 1231313, true},
		{1231223, "123113", true},
	}

	executor := suite.iValue.GetExecutor(constants.NotEquals)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, 11, true},
		{10, 10, false},
		{12, 10, false},
		{10, "9", false},
	}

	executor := suite.iValue.GetExecutor(constants.LessThan)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, 10, true},
		{10, 11, true},
		{12, 10, false},
		{10, "9", false},
	}

	executor := suite.iValue.GetExecutor(constants.LessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestLessThanAny() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{10, 11}, true},
		{10, []int{10, 10}, false},
		{12, []int{10, 9}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.LessThanAny)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestLessThanNone() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{10, 11}, false},
		{12, []int{10, 9}, true},
		{10, "test", true},
	}

	executor := suite.iValue.GetExecutor(constants.LessThanNone)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestLessThanOrEqualToAny() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{10, 11}, true},
		{10, []int{10, 10}, true},
		{12, []int{10, 9}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.LessThanOrEqualToAny)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestLessThanOrEqualToNone() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{10, 11}, false},
		{10, []int{10, 10}, false},
		{12, []int{10, 9}, true},
		{10, "test", true},
	}

	executor := suite.iValue.GetExecutor(constants.LessThanOrEqualToNone)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestLessThanAll() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{11, 12}, true},
		{10, []int{10, 11}, false},
		{10, []int{10, 10}, false},
		{12, []int{10, 9}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.LessThanAll)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestLessThanOrEqualToAll() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{11, 12}, true},
		{10, []int{10, 11}, true},
		{10, []int{10, 10}, true},
		{12, []int{11, 9}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.LessThanOrEqualToAll)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{13, 11, true},
		{10, 10, false},
		{9, 10, false},
		{10, "9", false},
	}

	executor := suite.iValue.GetExecutor(constants.GreaterThan)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{13, 11, true},
		{10, 10, true},
		{9, 10, false},
		{10, "9", false},
	}

	executor := suite.iValue.GetExecutor(constants.GreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestGreaterThanAny() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{9, 10}, true},
		{10, []int{10, 10}, false},
		{8, []int{10, 9}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.GreaterThanAny)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestGreaterThanNone() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{9, 10}, false},
		{10, []int{10, 10}, true},
		{8, []int{10, 9}, true},
		{10, "test", true},
	}

	executor := suite.iValue.GetExecutor(constants.GreaterThanNone)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestGreaterThanOrEqualToAny() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{9, 10}, true},
		{10, []int{11, 10}, true},
		{10, []int{10, 10}, true},
		{8, []int{10, 9}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.GreaterThanOrEqualToAny)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestGreaterThanOrEqualToNone() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{9, 10}, false},
		{10, []int{11, 10}, false},
		{10, []int{10, 10}, false},
		{8, []int{10, 9}, true},
		{10, "test", true},
	}

	executor := suite.iValue.GetExecutor(constants.GreaterThanOrEqualToNone)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestGreaterThanAll() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{9, 8}, true},
		{10, []int{9, 10}, false},
		{10, []int{10, 10}, false},
		{8, []int{11, 12}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.GreaterThanAll)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestGreaterThanOrEqualToAll() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10, []int{9, 8}, true},
		{10, []int{9, 10}, true},
		{10, []int{10, 10}, true},
		{8, []int{11, 12}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.GreaterThanOrEqualToAll)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestBetween() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{7, []int{6, 8}, true},
		{8, []int{8, 10}, false},
		{10, []int{8, 10}, false},
		{8, []int{11, 12}, false},
		{8, []int{11, 12, 13}, false},
		{8, []int{11}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.Between)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestNotBetween() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{7, []int{6, 8}, false},
		{8, []int{8, 10}, true},
		{10, []int{8, 10}, true},
		{8, []int{11, 12}, true},
		{10, "test", true},
	}

	executor := suite.iValue.GetExecutor(constants.NotBetween)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestAnyOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{7, []int{6, 7, 8}, true},
		{8, []int{8, 10}, true},
		{10, []int{8, 11}, false},
		{10, "test", false},
	}

	executor := suite.iValue.GetExecutor(constants.AnyOf)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *IntegerValueTestSuite) TestNoneOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{7, []int{6, 7, 8}, false},
		{8, []int{8, 10}, false},
		{10, []int{8, 11}, true},
		{10, "test", true},
	}

	executor := suite.iValue.GetExecutor(constants.NoneOf)

	for _, testData := range testDataSet {
		result := suite.iValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func TestIntegerValueTestSuite(t *testing.T) {
	suite.Run(t, new(IntegerValueTestSuite))
}
