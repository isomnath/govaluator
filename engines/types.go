package engines

import (
	"github.com/isomnath/govaluator/comparators"
	"github.com/isomnath/govaluator/engines/exitonstrategies"
	"github.com/isomnath/govaluator/models"
)

type Engines interface {
	Execute(map[string]interface{}, models.Rule) models.Result
}

type comparisonManager interface {
	GetComparator(string) comparators.Comparators
}

type utils interface {
	Evaluate(tmpResult map[string]interface{}, expression string) (interface{}, error)
	ProcessRemark(data map[string]interface{}, rule models.Rule, resultValue interface{}) (string, error)
}

type strategyManger interface {
	GetStrategy(strategyType string) exitonstrategies.ExitOnStrategy
}
