package govaluator

import (
	"github.com/isomnath/govaluator/comparators"
	"github.com/isomnath/govaluator/engines"
)

type engineManager interface {
	GetEngine(string) engines.Engines
	GetEngines() []string
}

type comparisonManager interface {
	GetComparator(string) comparators.Comparators
	GetComparators() []string
}

type operationManager interface {
	GetAllDataTypes() []string
	GetOperators(dataType string) []string
}

type utilities interface {
	Flatten(data map[string]interface{}) map[string]interface{}
}
