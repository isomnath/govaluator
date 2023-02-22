package exitonstrategies

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/constants"
)

type StrategyMangerTestSuite struct {
	suite.Suite
	strategyManager *ExitOnStrategyManger
}

func (suite *StrategyMangerTestSuite) SetupTest() {
	suite.strategyManager = InitializeStrategyManager()
}

func (suite *StrategyMangerTestSuite) TestGetStrategy() {
	type args struct {
		strategyName string
	}
	tests := []struct {
		name        string
		args        args
		expectation ExitOnStrategy
	}{
		{
			name:        constants.FirstFalse,
			args:        args{strategyName: constants.FirstFalse},
			expectation: suite.strategyManager.strategies[constants.FirstFalse],
		},
		{
			name:        constants.FirstTrue,
			args:        args{strategyName: constants.FirstTrue},
			expectation: suite.strategyManager.strategies[constants.FirstTrue],
		},
		{
			name:        constants.LastFalse,
			args:        args{strategyName: constants.LastFalse},
			expectation: suite.strategyManager.strategies[constants.LastFalse],
		},
		{
			name:        constants.LastTrue,
			args:        args{strategyName: constants.LastTrue},
			expectation: suite.strategyManager.strategies[constants.LastTrue],
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if got := suite.strategyManager.GetStrategy(tt.args.strategyName); !reflect.DeepEqual(got, tt.expectation) {
				suite.T().Errorf("filter() = %v, expectation %v", got, tt.expectation)
			}
		})
	}
}

func TestStrategyMangerTestSuite(t *testing.T) {
	suite.Run(t, new(StrategyMangerTestSuite))
}
