package comparators

import (
	"github.com/isomnath/govaluator/constants"
)

type ComparisonManager struct {
	comparators map[string]Comparators
}

func InitializeComparators(operationManager operationManager) *ComparisonManager {
	m := ComparisonManager{comparators: make(map[string]Comparators)}

	m.comparators[constants.StaticComparator] = initializeStaticComparator(operationManager)
	m.comparators[constants.DynamicComparator] = initializeDynamicComparator(operationManager)

	return &m
}

func (m *ComparisonManager) GetComparator(comparatorType string) Comparators {
	return m.comparators[comparatorType]
}

func (m *ComparisonManager) GetComparators() []string {
	var comparators []string
	for k := range m.comparators {
		comparators = append(comparators, k)
	}
	return comparators
}
