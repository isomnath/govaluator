package exitonstrategies

import (
	"github.com/isomnath/govaluator/constants"
)

type ExitOnStrategyManger struct {
	strategies map[string]ExitOnStrategy
}

func InitializeStrategyManager() *ExitOnStrategyManger {
	m := ExitOnStrategyManger{strategies: make(map[string]ExitOnStrategy)}

	m.strategies[constants.FirstTrue] = initializeFirstTrue()
	m.strategies[constants.LastTrue] = initializeLastTrue()
	m.strategies[constants.FirstFalse] = initializeFirstFalse()
	m.strategies[constants.LastFalse] = initializeLastFalse()

	return &m
}

func (m ExitOnStrategyManger) GetStrategy(strategyType string) ExitOnStrategy {
	return m.strategies[strategyType]
}
