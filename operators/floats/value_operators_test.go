package floats

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type FloatValueTestSuite struct {
	suite.Suite
	fValue            *FloatValue
	operatorUtilities operatorUtilities
}

func (suite *FloatValueTestSuite) SetupSuite() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.fValue = InitializeValueOperators(suite.operatorUtilities)
}

func (suite *FloatValueTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.Equals, constants.NotEquals,
		constants.LessThan, constants.LessThanOrEqualTo, constants.LessThanAny, constants.LessThanOrEqualToAny, constants.LessThanNone,
		constants.LessThanOrEqualToNone, constants.LessThanAll, constants.LessThanOrEqualToAll, constants.GreaterThan,
		constants.GreaterThanOrEqualTo, constants.GreaterThanAny, constants.GreaterThanOrEqualToAny, constants.GreaterThanNone,
		constants.GreaterThanOrEqualToNone, constants.GreaterThanAll, constants.GreaterThanOrEqualToAll, constants.Between,
		constants.NotBetween, constants.AnyOf, constants.NoneOf}

	actual := suite.fValue.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *FloatValueTestSuite) TestEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{1231223.0, 1231223.0, true},
		{1231223.0, 1231313.0, false},
		{1231223.0, "123113", false},
	}

	executor := suite.fValue.GetExecutor(constants.Equals)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestNotEquals() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{1231223.0, 1231223.0, false},
		{1231223.0, 1231313.0, true},
		{1231223.0, "123113", true},
	}

	executor := suite.fValue.GetExecutor(constants.NotEquals)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestLessThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, 11.0, true},
		{10.0, 10.0, false},
		{12.0, 10.0, false},
		{10.0, "9", false},
	}

	executor := suite.fValue.GetExecutor(constants.LessThan)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestLessThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, 10.0, true},
		{10.0, 11.0, true},
		{12.0, 10.0, false},
		{10.0, "9", false},
	}

	executor := suite.fValue.GetExecutor(constants.LessThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestLessThanAny() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{10.0, 11.0}, true},
		{10.0, []float64{10.0, 10.0}, false},
		{12.0, []float64{10.0, 9.0}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.LessThanAny)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestLessThanNone() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{10.0, 11.0}, false},
		{12.0, []float64{10.0, 9.0}, true},
		{10.0, "test", true},
	}

	executor := suite.fValue.GetExecutor(constants.LessThanNone)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestLessThanOrEqualToAny() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{10, 11}, true},
		{10.0, []float64{10, 10}, true},
		{12.0, []float64{10, 9}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.LessThanOrEqualToAny)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestLessThanOrEqualToNone() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{10, 11}, false},
		{10.0, []float64{10, 10}, false},
		{12.0, []float64{10, 9}, true},
		{10.0, "test", true},
	}

	executor := suite.fValue.GetExecutor(constants.LessThanOrEqualToNone)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestLessThanAll() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{11, 12}, true},
		{10.0, []float64{10, 11}, false},
		{10.0, []float64{10, 10}, false},
		{12.0, []float64{10, 9}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.LessThanAll)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestLessThanOrEqualToAll() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{11, 12}, true},
		{10.0, []float64{10, 11}, true},
		{10.0, []float64{10, 10}, true},
		{12.0, []float64{11, 9}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.LessThanOrEqualToAll)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestGreaterThan() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{13.0, 11.0, true},
		{10.0, 10.0, false},
		{9.0, 10.0, false},
		{10.0, "9", false},
	}

	executor := suite.fValue.GetExecutor(constants.GreaterThan)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestGreaterThanOrEqualTo() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{13.0, 11.0, true},
		{10.0, 10.0, true},
		{9.0, 10.0, false},
		{10.0, "9", false},
	}

	executor := suite.fValue.GetExecutor(constants.GreaterThanOrEqualTo)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestGreaterThanAny() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{9, 10}, true},
		{10.0, []float64{10, 10}, false},
		{8.0, []float64{10, 9}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.GreaterThanAny)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestGreaterThanNone() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{9, 10}, false},
		{10.0, []float64{10, 10}, true},
		{8.0, []float64{10, 9}, true},
		{10.0, "test", true},
	}

	executor := suite.fValue.GetExecutor(constants.GreaterThanNone)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestGreaterThanOrEqualToAny() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{9, 10}, true},
		{10.0, []float64{11, 10}, true},
		{10.0, []float64{10, 10}, true},
		{8.0, []float64{10, 9}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.GreaterThanOrEqualToAny)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestGreaterThanOrEqualToNone() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{9, 10}, false},
		{10.0, []float64{11, 10}, false},
		{10.0, []float64{10, 10}, false},
		{8.0, []float64{10, 9}, true},
		{10.0, "test", true},
	}

	executor := suite.fValue.GetExecutor(constants.GreaterThanOrEqualToNone)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestGreaterThanAll() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{9, 8}, true},
		{10.0, []float64{9, 10}, false},
		{10.0, []float64{10, 10}, false},
		{8.0, []float64{11, 12}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.GreaterThanAll)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestGreaterThanOrEqualToAll() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{10.0, []float64{9, 8}, true},
		{10.0, []float64{9, 10}, true},
		{10.0, []float64{10, 10}, true},
		{8.0, []float64{11, 12}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.GreaterThanOrEqualToAll)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestBetween() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{7.0, []float64{6, 8}, true},
		{8.0, []float64{8, 10}, false},
		{10.0, []float64{8, 10}, false},
		{8.0, []float64{11, 12}, false},
		{8.0, []float64{11, 12, 13}, false},
		{8.0, []float64{11}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.Between)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestNotBetween() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{7.0, []float64{6, 8}, false},
		{8.0, []float64{8, 10}, true},
		{10.0, []float64{8, 10}, true},
		{8.0, []float64{11, 12}, true},
		{10.0, "test", true},
	}

	executor := suite.fValue.GetExecutor(constants.NotBetween)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestAnyOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{7.0, []float64{6, 7, 8}, true},
		{8.0, []float64{8, 10}, true},
		{10.0, []float64{8, 11}, false},
		{10.0, "test", false},
	}

	executor := suite.fValue.GetExecutor(constants.AnyOf)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *FloatValueTestSuite) TestNoneOf() {
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{7.0, []float64{6, 7, 8}, false},
		{8.0, []float64{8, 10}, false},
		{10.0, []float64{8, 11}, true},
		{10.0, "test", true},
	}

	executor := suite.fValue.GetExecutor(constants.NoneOf)

	for _, testData := range testDataSet {
		result := suite.fValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func TestFloatValueTestSuite(t *testing.T) {
	suite.Run(t, new(FloatValueTestSuite))
}
