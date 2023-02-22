package dependencies

import (
	"github.com/isomnath/govaluator/comparators"
	"github.com/isomnath/govaluator/engines"
	"github.com/isomnath/govaluator/engines/exitonstrategies"
	"github.com/isomnath/govaluator/operators"
	"github.com/isomnath/govaluator/utilities"
)

type Dependencies struct {
	Utilities         utils
	OperatorUtilities operatorUtilities
	OperationManager  operationManager
	ComparisonManager comparisonManager
	StrategyManger    strategyManger
	EngineManager     engineManager
}

func Initialize() *Dependencies {
	u := utilities.InitializeUtilities()
	ou := utilities.InitializeOperatorUtilities()
	om := operators.InitializeOperators(ou)
	cm := comparators.InitializeComparators(om)
	sm := exitonstrategies.InitializeStrategyManager()
	em := engines.InitializeEngineManager(u, sm, cm)
	return &Dependencies{
		Utilities:         u,
		OperatorUtilities: ou,
		OperationManager:  om,
		ComparisonManager: cm,
		EngineManager:     em,
	}
}
