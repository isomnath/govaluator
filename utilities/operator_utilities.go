package utilities

import (
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set"

	"github.com/isomnath/govaluator/constants"
)

type OperatorUtilities struct{}

func (u *OperatorUtilities) CastToBooleanValue(i interface{}) (bool, error) {
	switch i.(type) {
	case bool:
		return i.(bool), nil
	default:
		return false, createError("bool")
	}
}

func (u *OperatorUtilities) CastToBooleanSlice(i interface{}) ([]bool, error) {
	switch i.(type) {
	case []bool:
		return i.([]bool), nil
	default:
		return nil, createError("[]bool")
	}
}

func (u *OperatorUtilities) CastBooleanSliceToFrequencyDistributionMap(i interface{}) (map[bool]int64, error) {
	switch i.(type) {
	case []bool:
		return u.castBooleanSliceToFrequencyDistributionMap(i.([]bool)), nil
	default:
		return nil, createError("[]bool")
	}
}

func (u *OperatorUtilities) CastBooleanSliceToSet(i interface{}) (mapset.Set, error) {
	switch i.(type) {
	case []bool:
		set := mapset.NewSet()
		for _, el := range i.([]bool) {
			set.Add(el)
		}
		return set, nil
	default:
		return nil, createError("[]bool")
	}
}

func (u *OperatorUtilities) CastToFloatValue(i interface{}) (float64, error) {
	switch i.(type) {
	case float32:
		return float64(i.(float32)), nil
	case float64:
		return i.(float64), nil
	default:
		return 0.0, createError("float")
	}
}

func (u *OperatorUtilities) CastToFloatSlice(i interface{}) ([]float64, error) {
	switch i.(type) {
	case []float32:
		return u.castFloat32SliceToFloat64Slice(i.([]float32)), nil
	case []float64:
		return i.([]float64), nil
	default:
		return nil, createError("[]float")
	}
}

func (u *OperatorUtilities) CastFloatSliceToFrequencyDistributionMap(i interface{}) (map[float64]int64, error) {
	switch i.(type) {
	case []float32:
		return u.castFloat32SliceToFrequencyDistributionMap(i.([]float32)), nil
	case []float64:
		return u.castFloat64SliceToFrequencyDistributionMap(i.([]float64)), nil
	default:
		return nil, createError("[]float")
	}
}

func (u *OperatorUtilities) CastFloatSliceToSet(i interface{}) (mapset.Set, error) {
	switch i.(type) {
	case []float32:
		set := mapset.NewSet()
		for _, el := range i.([]float32) {
			set.Add(float64(el))
		}
		return set, nil
	case []float64:
		set := mapset.NewSet()
		for _, el := range i.([]float64) {
			set.Add(el)
		}
		return set, nil
	default:
		return nil, createError("[]float")
	}
}

func (u *OperatorUtilities) CastToIntegerValue(i interface{}) (int64, error) {
	switch i.(type) {
	case int:
		return int64(i.(int)), nil
	case int8:
		return int64(i.(int8)), nil
	case int16:
		return int64(i.(int16)), nil
	case int32:
		return int64(i.(int32)), nil
	case int64:
		return i.(int64), nil
	default:
		return 0, createError("int")
	}
}

func (u *OperatorUtilities) CastToIntegerSlice(i interface{}) ([]int64, error) {
	switch i.(type) {
	case []int:
		return u.castIntSliceToInt64Slice(i.([]int)), nil
	case []int8:
		return u.castInt8SliceToInt64Slice(i.([]int8)), nil
	case []int16:
		return u.castInt16SliceToInt64Slice(i.([]int16)), nil
	case []int32:
		return u.castInt32SliceToInt64Slice(i.([]int32)), nil
	case []int64:
		return i.([]int64), nil
	default:
		return nil, createError("[]int")
	}
}

func (u *OperatorUtilities) CastIntegerSliceToFrequencyDistributionMap(i interface{}) (map[int64]int64, error) {
	switch i.(type) {
	case []int:
		return u.castIntSliceToFrequencyDistributionMap(i.([]int)), nil
	case []int8:
		return u.castInt8SliceToFrequencyDistributionMap(i.([]int8)), nil
	case []int16:
		return u.castInt16SliceToFrequencyDistributionMap(i.([]int16)), nil
	case []int32:
		return u.castInt32SliceToFrequencyDistributionMap(i.([]int32)), nil
	case []int64:
		return u.castInt64SliceToFrequencyDistributionMap(i.([]int64)), nil
	default:
		return nil, createError("[]int")
	}
}

func (u *OperatorUtilities) CastIntegerSliceToSet(i interface{}) (mapset.Set, error) {
	switch i.(type) {
	case []int:
		set := mapset.NewSet()
		for _, el := range i.([]int) {
			set.Add(int64(el))
		}
		return set, nil
	case []int8:
		set := mapset.NewSet()
		for _, el := range i.([]int8) {
			set.Add(int64(el))
		}
		return set, nil
	case []int16:
		set := mapset.NewSet()
		for _, el := range i.([]int16) {
			set.Add(int64(el))
		}
		return set, nil
	case []int32:
		set := mapset.NewSet()
		for _, el := range i.([]int32) {
			set.Add(int64(el))
		}
		return set, nil
	case []int64:
		set := mapset.NewSet()
		for _, el := range i.([]int64) {
			set.Add(el)
		}
		return set, nil
	default:
		return nil, createError("[]int")
	}
}

func (u *OperatorUtilities) CastToStringValue(i interface{}) (string, error) {
	switch i.(type) {
	case string:
		return i.(string), nil
	default:
		return "", createError("string")
	}
}

func (u *OperatorUtilities) CastToStringSlice(i interface{}) ([]string, error) {
	switch i.(type) {
	case []string:
		return i.([]string), nil
	default:
		return nil, createError("[]string")
	}
}

func (u *OperatorUtilities) CastStringSliceToFrequencyDistributionMap(i interface{}) (map[string]int64, error) {
	switch i.(type) {
	case []string:
		return u.castStringSliceToFrequencyDistributionMap(i.([]string)), nil
	default:
		return nil, createError("[]string")
	}
}

func (u *OperatorUtilities) CastStringSliceToSet(i interface{}) (mapset.Set, error) {
	switch i.(type) {
	case []string:
		set := mapset.NewSet()
		for _, el := range i.([]string) {
			set.Add(el)
		}
		return set, nil
	default:
		return nil, createError("[]string")
	}
}

func (u *OperatorUtilities) CastToTimeValue(i interface{}) (*time.Time, error) {
	switch i.(type) {
	case time.Time:
		val := i.(time.Time)
		return &val, nil
	case *time.Time:
		return i.(*time.Time), nil
	default:
		return nil, createError("times")
	}
}

func (u *OperatorUtilities) CastToTimeSlice(i interface{}) ([]*time.Time, error) {
	switch i.(type) {
	case []time.Time:
		val := i.([]time.Time)
		var castedTimeSlice []*time.Time
		for _, t := range val {
			t0 := t
			castedTimeSlice = append(castedTimeSlice, &t0)
		}
		return castedTimeSlice, nil
	case []*time.Time:
		return i.([]*time.Time), nil
	default:
		return nil, createError("[]times")
	}
}

func (u *OperatorUtilities) castFloat32SliceToFloat64Slice(slice []float32) []float64 {
	var castedSlice []float64
	for _, val := range slice {
		castedSlice = append(castedSlice, float64(val))
	}
	return castedSlice
}

func (u *OperatorUtilities) castIntSliceToInt64Slice(slice []int) []int64 {
	var castedSlice []int64
	for _, val := range slice {
		castedSlice = append(castedSlice, int64(val))
	}
	return castedSlice
}

func (u *OperatorUtilities) castInt8SliceToInt64Slice(slice []int8) []int64 {
	var castedSlice []int64
	for _, val := range slice {
		castedSlice = append(castedSlice, int64(val))
	}
	return castedSlice
}

func (u *OperatorUtilities) castInt16SliceToInt64Slice(slice []int16) []int64 {
	var castedSlice []int64
	for _, val := range slice {
		castedSlice = append(castedSlice, int64(val))
	}
	return castedSlice
}

func (u *OperatorUtilities) castInt32SliceToInt64Slice(slice []int32) []int64 {
	var castedSlice []int64
	for _, val := range slice {
		castedSlice = append(castedSlice, int64(val))
	}
	return castedSlice
}

func (u *OperatorUtilities) castBooleanSliceToFrequencyDistributionMap(slice []bool) map[bool]int64 {
	fdMap := make(map[bool]int64)
	for i := 0; i < len(slice); i++ {
		fdMap[slice[i]] = fdMap[slice[i]] + 1
	}
	return fdMap
}

func (u *OperatorUtilities) castFloat32SliceToFrequencyDistributionMap(slice []float32) map[float64]int64 {
	fdMap := make(map[float64]int64)
	for i := 0; i < len(slice); i++ {
		fdMap[float64(slice[i])] = fdMap[float64(slice[i])] + 1
	}
	return fdMap
}

func (u *OperatorUtilities) castFloat64SliceToFrequencyDistributionMap(slice []float64) map[float64]int64 {
	fdMap := make(map[float64]int64)
	for i := 0; i < len(slice); i++ {
		fdMap[slice[i]] = fdMap[slice[i]] + 1
	}
	return fdMap
}

func (u *OperatorUtilities) castIntSliceToFrequencyDistributionMap(slice []int) map[int64]int64 {
	fdMap := make(map[int64]int64)
	for i := 0; i < len(slice); i++ {
		fdMap[int64(slice[i])] = fdMap[int64(slice[i])] + 1
	}
	return fdMap
}

func (u *OperatorUtilities) castInt8SliceToFrequencyDistributionMap(slice []int8) map[int64]int64 {
	fdMap := make(map[int64]int64)
	for i := 0; i < len(slice); i++ {
		fdMap[int64(slice[i])] = fdMap[int64(slice[i])] + 1
	}
	return fdMap
}

func (u *OperatorUtilities) castInt16SliceToFrequencyDistributionMap(slice []int16) map[int64]int64 {
	fdMap := make(map[int64]int64)
	for i := 0; i < len(slice); i++ {
		fdMap[int64(slice[i])] = fdMap[int64(slice[i])] + 1
	}
	return fdMap
}

func (u *OperatorUtilities) castInt32SliceToFrequencyDistributionMap(slice []int32) map[int64]int64 {
	fdMap := make(map[int64]int64)
	for i := 0; i < len(slice); i++ {
		fdMap[int64(slice[i])] = fdMap[int64(slice[i])] + 1
	}
	return fdMap
}

func (u *OperatorUtilities) castInt64SliceToFrequencyDistributionMap(slice []int64) map[int64]int64 {
	fdMap := make(map[int64]int64)
	for i := 0; i < len(slice); i++ {
		fdMap[slice[i]] = fdMap[slice[i]] + 1
	}
	return fdMap
}

func (u *OperatorUtilities) castStringSliceToFrequencyDistributionMap(slice []string) map[string]int64 {
	fdMap := make(map[string]int64)
	for i := 0; i < len(slice); i++ {
		fdMap[slice[i]] = fdMap[slice[i]] + 1
	}
	return fdMap
}

func createError(interfaceType string) error {
	return fmt.Errorf(constants.ErrorInvalidDataType, interfaceType)
}

func InitializeOperatorUtilities() *OperatorUtilities {
	return &OperatorUtilities{}
}
