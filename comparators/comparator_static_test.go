package comparators

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/models"
	"github.com/isomnath/govaluator/operators"
	"github.com/isomnath/govaluator/utilities"
)

type StaticComparatorTestSuite struct {
	suite.Suite
	operationManager operationManager
	comparator       Comparators
}

func (suite *StaticComparatorTestSuite) SetupTest() {
	suite.operationManager = operators.InitializeOperators(utilities.InitializeOperatorUtilities())
	suite.comparator = initializeStaticComparator(suite.operationManager)
}

func (suite *StaticComparatorTestSuite) TestExecuteCriterionReturnFalseWhenExecutorNotFound() {
	criterion := models.Criterion{
		ID:         "123",
		Alias:      "A",
		Comparator: constants.StaticComparator,
		Static: &models.Static{
			FieldOne: "parent.child_1.child_2",
			Operator: constants.Equals,
			Value:    []int{1231},
		},
	}

	data := map[string]interface{}{
		"parent.child_1.child_2": []int{1231},
	}

	flag, result := suite.comparator.ExecuteCriterion(data, criterion)
	suite.Equal(false, flag)
	suite.Equal(false, result)
}

func (suite *StaticComparatorTestSuite) TestExecuteCriterionReturnResultOnSuccessfulExecution() {
	criterion := models.Criterion{
		ID:         "123",
		Alias:      "A",
		Comparator: constants.StaticComparator,
		Static: &models.Static{
			FieldOne: "parent.child_1.child_2",
			Operator: constants.OrderedEquals,
			Value:    []int{1231},
		},
	}

	data := map[string]interface{}{
		"parent.child_1.child_2": []int{1231},
	}

	flag, result := suite.comparator.ExecuteCriterion(data, criterion)
	suite.Equal(true, flag)
	suite.Equal(true, result)
}

func TestStaticComparatorTestSuite(t *testing.T) {
	suite.Run(t, new(StaticComparatorTestSuite))
}
