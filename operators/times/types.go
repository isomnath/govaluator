package times

import (
	"time"

	mapset "github.com/deckarep/golang-set"
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
