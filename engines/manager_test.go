package engines

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/comparators"
	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/engines/exitonstrategies"
	"github.com/isomnath/govaluator/operators"
	"github.com/isomnath/govaluator/utilities"
)

type EngineManagerTestSuite struct {
	suite.Suite
	utilities         utils
	strategyManger    strategyManger
	comparisonManager comparisonManager
	engineManager     *EngineManager
}

func (suite *EngineManagerTestSuite) SetupTest() {
	u := utilities.InitializeUtilities()
	ou := utilities.InitializeOperatorUtilities()
	sm := exitonstrategies.InitializeStrategyManager()
	om := operators.InitializeOperators(ou)
	cm := comparators.InitializeComparators(om)
	suite.engineManager = InitializeEngineManager(u, sm, cm)
}

func (suite *EngineManagerTestSuite) TestGetEngine() {
	type args struct {
		engineType string
	}
	tests := []struct {
		name        string
		args        args
		expectation Engines
	}{
		{
			name: constants.LinearEngine,
			args: args{
				engineType: constants.LinearEngine,
			},
			expectation: suite.engineManager.engines[constants.LinearEngine],
		},
		{
			name: constants.LinearWaterfallEngine,
			args: args{
				engineType: constants.LinearWaterfallEngine,
			},
			expectation: suite.engineManager.engines[constants.LinearWaterfallEngine],
		},
	}

	for _, tt := range tests {
		operator := suite.engineManager.GetEngine(tt.args.engineType)
		suite.Equal(tt.expectation, operator)
	}

}

func (suite *EngineManagerTestSuite) TestGetEngines() {
	expected := []string{constants.LinearEngine, constants.LinearWaterfallEngine}
	actual := suite.engineManager.GetEngines()
	sort.Strings(expected)
	sort.Strings(actual)
	suite.Equal(expected, actual)
}

func TestEngineManagerTestSuite(t *testing.T) {
	suite.Run(t, new(EngineManagerTestSuite))
}
