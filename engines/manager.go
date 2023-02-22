package engines

import (
	"github.com/isomnath/govaluator/constants"
)

type EngineManager struct {
	engines map[string]Engines
}

func InitializeEngineManager(utilities utils, strategyManager strategyManger, comparisonManager comparisonManager) *EngineManager {
	m := EngineManager{engines: make(map[string]Engines)}

	m.engines[constants.LinearEngine] = initializeLinearEngine(utilities, comparisonManager)
	m.engines[constants.LinearWaterfallEngine] = initializeLinearWaterfallEngine(utilities, strategyManager, comparisonManager)

	return &m
}

func (m EngineManager) GetEngine(engineType string) Engines {
	return m.engines[engineType]
}

func (m EngineManager) GetEngines() []string {
	var engines []string
	for k := range m.engines {
		engines = append(engines, k)
	}
	return engines
}
