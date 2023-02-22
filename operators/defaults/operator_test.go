package defaults

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
)

type DefaultOperatorTestSuite struct {
	suite.Suite
	defaultOps *DefaultOperator
}

func (suite *DefaultOperatorTestSuite) SetupSuite() {
	suite.defaultOps = InitializeOperators()
}

func (suite *DefaultOperatorTestSuite) TestGetOperators() {
	expectedOperators := []string{constants.Default}

	actual := suite.defaultOps.GetOperators()
	sort.Strings(expectedOperators)
	sort.Strings(actual)
	suite.Equal(expectedOperators, actual)
}

func (suite *DefaultOperatorTestSuite) TestDefaultOperator() {
	executor := suite.defaultOps.GetExecutor("Invalid Operator Function")
	suite.False(suite.defaultOps.ExecuteOperator(executor, true, false))
}

func TestDefaultOperatorTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultOperatorTestSuite))
}
