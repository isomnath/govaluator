package times

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type TimeValueTestSuite struct {
	suite.Suite
	tValue            *TimeValue
	operatorUtilities operatorUtilities
}

func (suite *TimeValueTestSuite) SetupSuite() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.tValue = InitializeValueOperators(suite.operatorUtilities)
}

func (suite *TimeValueTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.Before, constants.After, constants.During}

	actual := suite.tValue.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *TimeValueTestSuite) TestBefore() {
	t0 := time.Now().UTC()
	t1 := t0.Add(-90 * 24 * 60 * time.Minute)
	t2 := t0
	t3 := t0.AddDate(0, 3, 0)
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{t1, t2, true},
		{t3, t2, false},
		{"test", 123113, false},
	}

	executor := suite.tValue.GetExecutor(constants.Before)

	for _, testData := range testDataSet {
		result := suite.tValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *TimeValueTestSuite) TestAfter() {
	t0 := time.Now().UTC()
	t1 := t0.Add(-90 * 24 * 60 * time.Minute)
	t2 := t0
	t3 := t0.AddDate(0, 3, 0)
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{t2, t1, true},
		{t2, t3, false},
		{"test", 123113, false},
	}

	executor := suite.tValue.GetExecutor(constants.After)

	for _, testData := range testDataSet {
		result := suite.tValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func (suite *TimeValueTestSuite) TestDuring() {
	t0 := time.Now().UTC()
	t1 := t0.Add(-90 * 24 * 60 * time.Minute)
	t2 := t0
	t3 := t0.AddDate(0, 3, 0)
	var testDataSet = []struct {
		val1        interface{}
		val2        interface{}
		expectation bool
	}{
		{t2, []time.Time{t1, t3}, true},
		{t3, []time.Time{t1, t2}, false},
		{t2, []time.Time{}, false},
		{"test", 123113, false},
	}

	executor := suite.tValue.GetExecutor(constants.During)

	for _, testData := range testDataSet {
		result := suite.tValue.ExecuteOperator(executor, testData.val1, testData.val2)
		suite.Equal(testData.expectation, result)
	}
}

func TestTimeValueTestSuite(t *testing.T) {
	suite.Run(t, new(TimeValueTestSuite))
}
