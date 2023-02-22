package booleans

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/utilities"
)

type BooleanValueTestSuite struct {
	suite.Suite
	bValue            *BooleanValue
	operatorUtilities operatorUtilities
}

func (suite *BooleanValueTestSuite) SetupSuite() {
	suite.operatorUtilities = utilities.InitializeOperatorUtilities()
	suite.bValue = InitializeValueOperators(suite.operatorUtilities)
}

func (suite *BooleanValueTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.Equals, constants.NotEquals, constants.AnyOf, constants.NoneOf}

	actual := suite.bValue.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *BooleanValueTestSuite) TestEqualsShouldReturnTrue() {
	executor := suite.bValue.GetExecutor(constants.Equals)
	suite.True(suite.bValue.ExecuteOperator(executor, true, true))
}

func (suite *BooleanValueTestSuite) TestEqualsShouldReturnFalse() {
	executor := suite.bValue.GetExecutor(constants.Equals)
	suite.False(suite.bValue.ExecuteOperator(executor, true, false))
}

func (suite *BooleanValueTestSuite) TestEqualsShouldReturnFalseWhenVal2IsNotBoolean() {
	executor := suite.bValue.GetExecutor(constants.Equals)
	suite.False(suite.bValue.ExecuteOperator(executor, true, "string"))
}

func (suite *BooleanValueTestSuite) TestNotEqualsShouldReturnTrue() {
	executor := suite.bValue.GetExecutor(constants.NotEquals)
	suite.True(suite.bValue.ExecuteOperator(executor, true, false))
}

func (suite *BooleanValueTestSuite) TestNotEqualsShouldReturnFalse() {
	executor := suite.bValue.GetExecutor(constants.NotEquals)
	suite.False(suite.bValue.ExecuteOperator(executor, true, true))
}

func (suite *BooleanValueTestSuite) TestAnyOfShouldReturnTrue() {
	executor := suite.bValue.GetExecutor(constants.AnyOf)
	suite.True(suite.bValue.ExecuteOperator(executor, true, []bool{true, false}))
}

func (suite *BooleanValueTestSuite) TestAnyOfShouldReturnFalse() {
	executor := suite.bValue.GetExecutor(constants.AnyOf)
	suite.False(suite.bValue.ExecuteOperator(executor, true, []bool{false, false}))
}

func (suite *BooleanValueTestSuite) TestAnyOfShouldReturnFalseWhenVal2IsBooleanSlice() {
	executor := suite.bValue.GetExecutor(constants.AnyOf)
	suite.False(suite.bValue.ExecuteOperator(executor, true, []string{"stringOne", "stringTwo"}))
}

func (suite *BooleanValueTestSuite) TestNoneOfShouldReturnTrue() {
	executor := suite.bValue.GetExecutor(constants.NoneOf)
	suite.True(suite.bValue.ExecuteOperator(executor, true, []bool{false, false}))
}

func (suite *BooleanValueTestSuite) TestNoneOfShouldReturnFalse() {
	executor := suite.bValue.GetExecutor(constants.NoneOf)
	suite.False(suite.bValue.ExecuteOperator(executor, true, []bool{true, false}))
}

func TestBooleanValueTestSuite(t *testing.T) {
	suite.Run(t, new(BooleanValueTestSuite))
}
