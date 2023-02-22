package comparators

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/operators"
	"github.com/isomnath/govaluator/utilities"
)

type ComparisonManagerTestSuite struct {
	suite.Suite
	operationManager  operationManager
	comparisonManager *ComparisonManager
}

func (suite *ComparisonManagerTestSuite) SetupTest() {
	suite.operationManager = operators.InitializeOperators(utilities.InitializeOperatorUtilities())
	suite.comparisonManager = InitializeComparators(suite.operationManager)
}

func (suite *ComparisonManagerTestSuite) TestGetComparator() {
	type args struct {
		comparatorType string
	}

	tests := []struct {
		name        string
		args        args
		expectation Comparators
	}{
		{
			name:        "Static Comparator",
			args:        args{comparatorType: constants.StaticComparator},
			expectation: suite.comparisonManager.comparators[constants.StaticComparator],
		},
		{
			name:        "Dynamic Comparator",
			args:        args{comparatorType: constants.DynamicComparator},
			expectation: suite.comparisonManager.comparators[constants.DynamicComparator],
		},
	}

	for _, tt := range tests {
		operator := suite.comparisonManager.GetComparator(tt.args.comparatorType)
		suite.Equal(tt.expectation, operator)
	}
}

func (suite *ComparisonManagerTestSuite) TestGetComparators() {
	expected := []string{constants.StaticComparator, constants.DynamicComparator}
	actual := suite.comparisonManager.GetComparators()
	sort.Strings(expected)
	sort.Strings(actual)
	suite.Equal(expected, actual)
}

func TestComparatorsManagerTestSuite(t *testing.T) {
	suite.Run(t, new(ComparisonManagerTestSuite))
}
