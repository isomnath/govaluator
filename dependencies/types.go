package dependencies

import (
	"time"

	mapset "github.com/deckarep/golang-set"

	"github.com/isomnath/govaluator/comparators"
	"github.com/isomnath/govaluator/engines"
	"github.com/isomnath/govaluator/engines/exitonstrategies"
	"github.com/isomnath/govaluator/models"
	"github.com/isomnath/govaluator/operators"
)

type operatorUtilities interface {
	CastToBooleanValue(interface{}) (bool, error)
	CastToBooleanSlice(interface{}) ([]bool, error)
	CastBooleanSliceToFrequencyDistributionMap(i interface{}) (map[bool]int64, error)
	CastBooleanSliceToSet(i interface{}) (mapset.Set, error)

	CastToFloatValue(interface{}) (float64, error)
	CastToFloatSlice(interface{}) ([]float64, error)
	CastFloatSliceToFrequencyDistributionMap(i interface{}) (map[float64]int64, error)
	CastFloatSliceToSet(i interface{}) (mapset.Set, error)

	CastToIntegerValue(interface{}) (int64, error)
	CastToIntegerSlice(interface{}) ([]int64, error)
	CastIntegerSliceToFrequencyDistributionMap(i interface{}) (map[int64]int64, error)
	CastIntegerSliceToSet(i interface{}) (mapset.Set, error)

	CastToStringValue(interface{}) (string, error)
	CastToStringSlice(interface{}) ([]string, error)
	CastStringSliceToFrequencyDistributionMap(i interface{}) (map[string]int64, error)
	CastStringSliceToSet(i interface{}) (mapset.Set, error)

	CastToTimeValue(i interface{}) (*time.Time, error)
	CastToTimeSlice(i interface{}) ([]*time.Time, error)
}

type utils interface {
	Evaluate(tmpResult map[string]interface{}, expression string) (interface{}, error)
	Flatten(data map[string]interface{}) map[string]interface{}
	ProcessRemark(data map[string]interface{}, rule models.Rule, resultValue interface{}) (string, error)
}

type operationManager interface {
	GetOperator(interface{}) operators.Operators
	GetAllDataTypes() []string
	GetOperators(dataType string) []string
}

type comparisonManager interface {
	GetComparator(string) comparators.Comparators
	GetComparators() []string
}

type engineManager interface {
	GetEngine(string) engines.Engines
	GetEngines() []string
}

type strategyManger interface {
	GetStrategy(strategyType string) exitonstrategies.ExitOnStrategy
}
