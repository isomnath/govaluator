package comparators

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/models"
	"github.com/isomnath/govaluator/operators"
	"github.com/isomnath/govaluator/utilities"
)

type DynamicComparatorTestSuite struct {
	suite.Suite
	operationManager operationManager
	comparator       Comparators
}

func (suite *DynamicComparatorTestSuite) SetupTest() {
	suite.operationManager = operators.InitializeOperators(utilities.InitializeOperatorUtilities())
	suite.comparator = initializeDynamicComparator(suite.operationManager)
}

func (suite *DynamicComparatorTestSuite) TestExecuteCriterionReturnFalseWhenExecutorNotFound() {
	criterion := models.Criterion{
		ID:         "123",
		Alias:      "A",
		Comparator: constants.DynamicComparator,
		Dynamic: &models.Dynamic{
			FieldOne: "parent.child_1.child_2",
			Operator: constants.Equals,
			FieldTwo: "parent.child_1.child_3",
		},
	}

	data := map[string]interface{}{
		"parent.child_1.child_2": []int{1231},
		"parent.child_1.child_3": []int{1231},
	}

	flag, result := suite.comparator.ExecuteCriterion(data, criterion)
	suite.Equal(false, flag)
	suite.Equal(false, result)
}

func (suite *DynamicComparatorTestSuite) TestExecuteCriterionReturnResultOnSuccessfulExecution() {
	criterion := models.Criterion{
		ID:         "123",
		Alias:      "A",
		Comparator: constants.DynamicComparator,
		Dynamic: &models.Dynamic{
			FieldOne: "parent.child_1.child_2",
			Operator: constants.OrderedEquals,
			FieldTwo: "parent.child_1.child_3",
		},
	}

	data := map[string]interface{}{
		"parent.child_1.child_2": []int{1231},
		"parent.child_1.child_3": []int{1231},
	}

	flag, result := suite.comparator.ExecuteCriterion(data, criterion)
	suite.Equal(true, flag)
	suite.Equal(true, result)
}

func TestDynamicComparatorTestSuite(t *testing.T) {
	suite.Run(t, new(DynamicComparatorTestSuite))
}
