package integers

import (
	"reflect"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type IntegerSlice struct {
	utils       operatorUtilities
	operatorMap map[string]descriptors.Executor
}

func (iso *IntegerSlice) GetOperators() []string {
	var operators []string
	for k := range iso.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (iso *IntegerSlice) GetExecutor(operatorName string) descriptors.Executor {
	return iso.operatorMap[operatorName]
}

func (iso *IntegerSlice) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// unOrderedEquals To be read as slice val1 and val2 are equal ignoring order
func (iso *IntegerSlice) unOrderedEquals(val1 interface{}, val2 interface{}) bool {
	val1FdMap, err := iso.utils.CastIntegerSliceToFrequencyDistributionMap(val1)
	val2FdMap, err := iso.utils.CastIntegerSliceToFrequencyDistributionMap(val2)
	if err != nil {
		return false
	}
	return reflect.DeepEqual(val1FdMap, val2FdMap)
}

// unOrderedNotEquals To be read as slice val1 and val2 are not equal ignoring order
func (iso *IntegerSlice) unOrderedNotEquals(val1 interface{}, val2 interface{}) bool {
	return !iso.unOrderedEquals(val1, val2)
}

// orderedEquals To be read as slice val1 and val2 are equal considering order
func (iso *IntegerSlice) orderedEquals(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := iso.utils.CastToIntegerSlice(val1)
	bSlice2, err := iso.utils.CastToIntegerSlice(val2)
	if err != nil {
		return false
	}
	if len(bSlice1) != len(bSlice2) {
		return false
	}
	for index, bVal1 := range bSlice1 {
		if bVal1 != bSlice2[index] {
			return false
		}
	}
	return true
}

// orderedNotEquals To be read as slice val1 and val2 are not equal considering order
func (iso *IntegerSlice) orderedNotEquals(val1 interface{}, val2 interface{}) bool {
	return !iso.orderedEquals(val1, val2)
}

// anyLessThan To be read as any-of val1 less-than val2
func (iso *IntegerSlice) anyLessThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := iso.utils.CastToIntegerSlice(val1)
	bVal2, err2 := iso.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	for _, bVal1 := range bSlice1 {
		if bVal1 < bVal2 {
			return true
		}
	}
	return false
}

// noneLessThan To be read as none-of val1 less-than val2
func (iso *IntegerSlice) noneLessThan(val1 interface{}, val2 interface{}) bool {
	return !iso.anyLessThan(val1, val2)
}

// anyLessThanOrEqualTo To be read as any-of val1 less-than-or-equal-to val2
func (iso *IntegerSlice) anyLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := iso.utils.CastToIntegerSlice(val1)
	bVal2, err2 := iso.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	for _, bVal1 := range bSlice1 {
		if bVal1 <= bVal2 {
			return true
		}
	}
	return false
}

// noneLessThanOrEqualTo To be read as none-of val1 less-than-or-equal-to val2
func (iso *IntegerSlice) noneLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	return !iso.anyLessThanOrEqualTo(val1, val2)
}

// allLessThan To be read as all-of val1 less-than val2
func (iso *IntegerSlice) allLessThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := iso.utils.CastToIntegerSlice(val1)
	bVal2, err2 := iso.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal1 := range bSlice1 {
		if bVal1 < bVal2 {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

// allLessThanOrEqualTo To be read as all-of val1 less-than-or-equal-to val2
func (iso *IntegerSlice) allLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := iso.utils.CastToIntegerSlice(val1)
	bVal2, err2 := iso.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal1 := range bSlice1 {
		if bVal1 <= bVal2 {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

// anyGreaterThan To be read as any-of val1 greater-than val2
func (iso *IntegerSlice) anyGreaterThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := iso.utils.CastToIntegerSlice(val1)
	bVal2, err2 := iso.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	for _, bVal1 := range bSlice1 {
		if bVal1 > bVal2 {
			return true
		}
	}
	return false
}

// noneGreaterThan To be read as none-of val1 greater-than val2
func (iso *IntegerSlice) noneGreaterThan(val1 interface{}, val2 interface{}) bool {
	return !iso.anyGreaterThan(val1, val2)
}

// anyGreaterThanOrEqualTo To be read as any-of val1 greater-than-or-equal-to val2
func (iso *IntegerSlice) anyGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := iso.utils.CastToIntegerSlice(val1)
	bVal2, err2 := iso.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil || len(bSlice1) < 1 {
		return false
	}
	for _, bVal1 := range bSlice1 {
		if bVal1 >= bVal2 {
			return true
		}
	}
	return false
}

// noneGreaterThanOrEqualTo To be read as none-of val1 greater-than-or-equal-to val2
func (iso *IntegerSlice) noneGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	return !iso.anyGreaterThanOrEqualTo(val1, val2)
}

// allGreaterThan To be read as all-of val1 less-than val2
func (iso *IntegerSlice) allGreaterThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := iso.utils.CastToIntegerSlice(val1)
	bVal2, err2 := iso.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal1 := range bSlice1 {
		if bVal1 > bVal2 {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

// allGreaterThanOrEqualTo To be read as all-of val1 greater-than-or-equal-to val2
func (iso *IntegerSlice) allGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := iso.utils.CastToIntegerSlice(val1)
	bVal2, err2 := iso.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal1 := range bSlice1 {
		if bVal1 >= bVal2 {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

// reversedBetween To be read as val2 is within the range of values contained in val1  (boundaries not considered)
func (iso *IntegerSlice) reversedBetween(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := iso.utils.CastToIntegerSlice(val1)
	bVal2, err2 := iso.utils.CastToIntegerValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	if len(bSlice1) < 2 || len(bSlice1) > 2 {
		return false
	}
	return bVal2 > bSlice1[0] && bVal2 < bSlice1[1]
}

// reversedNotBetween To be read as val2 is not within the range of values contained in val1  (boundaries not considered)
func (iso *IntegerSlice) reversedNotBetween(val1 interface{}, val2 interface{}) bool {
	return !iso.reversedBetween(val1, val2)
}

// supersetOf To be read as val1 is a superset-of val2
func (iso *IntegerSlice) supersetOf(val1 interface{}, val2 interface{}) bool {
	set1, err := iso.utils.CastIntegerSliceToSet(val1)
	set2, err := iso.utils.CastIntegerSliceToSet(val2)
	if err != nil {
		return false
	}
	return set1.IsSuperset(set2)
}

// notSupersetOf To be read as val1 is not a superset-of val2
func (iso *IntegerSlice) notSupersetOf(val1 interface{}, val2 interface{}) bool {
	return !iso.supersetOf(val1, val2)
}

// subsetOf To be read as val1 is a subset-of val2
func (iso *IntegerSlice) subsetOf(val1 interface{}, val2 interface{}) bool {
	set1, err := iso.utils.CastIntegerSliceToSet(val1)
	set2, err := iso.utils.CastIntegerSliceToSet(val2)
	if err != nil {
		return false
	}
	return set1.IsSubset(set2)
}

// notSubsetOf To be read as val1 is not a subset-of val2
func (iso *IntegerSlice) notSubsetOf(val1 interface{}, val2 interface{}) bool {
	return !iso.subsetOf(val1, val2)
}

// intersection To be read as val1 intersects val2
func (iso *IntegerSlice) intersection(val1 interface{}, val2 interface{}) bool {
	set1, err := iso.utils.CastIntegerSliceToSet(val1)
	set2, err := iso.utils.CastIntegerSliceToSet(val2)
	if err != nil {
		return false
	}
	intersectionSet := set1.Intersect(set2)
	return len(intersectionSet.ToSlice()) > 0
}

// NotIntersection To be read as val1 does not intersect val2
func (iso *IntegerSlice) notIntersection(val1 interface{}, val2 interface{}) bool {
	return !iso.intersection(val1, val2)
}

// reversedAnyOf To be read as val2 is any-of val1
func (iso *IntegerSlice) reversedAnyOf(val1 interface{}, val2 interface{}) bool {
	bVal2, err1 := iso.utils.CastToIntegerValue(val2)
	bSlice1, err2 := iso.utils.CastToIntegerSlice(val1)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, v := range bSlice1 {
		if bVal2 == v {
			flag = true
			break
		}
	}
	return flag
}

// reversedNoneOf To be read as val2 is none-of val1
func (iso *IntegerSlice) reversedNoneOf(val1 interface{}, val2 interface{}) bool {
	return !iso.reversedAnyOf(val1, val2)
}

// sizeEquals To be read as size of val1 is equal-to val2
func (iso *IntegerSlice) sizeEquals(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := iso.utils.CastToIntegerSlice(val1)
	bVal2, err := iso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) == bVal2
}

// sizeNotEquals To be read as size of val1 is not-equal-to val2
func (iso *IntegerSlice) sizeNotEquals(val1 interface{}, val2 interface{}) bool {
	return !iso.sizeEquals(val1, val2)
}

// sizeLessThan To be read as size of val1 is less-than val2
func (iso *IntegerSlice) sizeLessThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := iso.utils.CastToIntegerSlice(val1)
	bVal2, err := iso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) < bVal2
}

// sizeLessThanOrEqualTo To be read as size of val1 is less-than-or-equal-to val2
func (iso *IntegerSlice) sizeLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := iso.utils.CastToIntegerSlice(val1)
	bVal2, err := iso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) <= bVal2
}

// sizeGreaterThan To be read as size of val1 is greater-than val2
func (iso *IntegerSlice) sizeGreaterThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := iso.utils.CastToIntegerSlice(val1)
	bVal2, err := iso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) > bVal2
}

// sizeGreaterThanOrEqualTo To be read as size of val1 is greater-than-or-equal-to val2
func (iso *IntegerSlice) sizeGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := iso.utils.CastToIntegerSlice(val1)
	bVal2, err := iso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) >= bVal2
}

func InitializeSliceOperators(operatorUtilities operatorUtilities) *IntegerSlice {
	op := IntegerSlice{
		utils: operatorUtilities,
	}

	operatorMap := make(map[string]descriptors.Executor, 0)
	operatorMap[constants.UnorderedEquals] = op.unOrderedEquals
	operatorMap[constants.UnorderedNotEquals] = op.unOrderedNotEquals
	operatorMap[constants.OrderedEquals] = op.orderedEquals
	operatorMap[constants.OrderedNotEquals] = op.orderedNotEquals

	operatorMap[constants.AnyLessThan] = op.anyLessThan
	operatorMap[constants.AnyLessThanOrEqualTo] = op.anyLessThanOrEqualTo
	operatorMap[constants.NoneLessThan] = op.noneLessThan
	operatorMap[constants.NoneLessThanOrEqualTo] = op.noneLessThanOrEqualTo
	operatorMap[constants.AllLessThan] = op.allLessThan
	operatorMap[constants.AllLessThanOrEqualTo] = op.allLessThanOrEqualTo
	operatorMap[constants.AnyGreaterThan] = op.anyGreaterThan
	operatorMap[constants.AnyGreaterThanOrEqualTo] = op.anyGreaterThanOrEqualTo
	operatorMap[constants.NoneGreaterThan] = op.noneGreaterThan
	operatorMap[constants.NoneGreaterThanOrEqualTo] = op.noneGreaterThanOrEqualTo
	operatorMap[constants.AllGreaterThan] = op.allGreaterThan
	operatorMap[constants.AllGreaterThanOrEqualTo] = op.allGreaterThanOrEqualTo
	operatorMap[constants.ReversedBetween] = op.reversedBetween
	operatorMap[constants.ReversedNotBetween] = op.reversedNotBetween

	operatorMap[constants.SupersetOf] = op.supersetOf
	operatorMap[constants.NotSupersetOf] = op.notSupersetOf
	operatorMap[constants.SubsetOf] = op.subsetOf
	operatorMap[constants.NotSubsetOf] = op.notSubsetOf
	operatorMap[constants.Intersection] = op.intersection
	operatorMap[constants.NotIntersection] = op.notIntersection
	operatorMap[constants.ReversedAnyOf] = op.reversedAnyOf
	operatorMap[constants.ReversedNoneOf] = op.reversedNoneOf

	operatorMap[constants.SizeEquals] = op.sizeEquals
	operatorMap[constants.SizeNotEquals] = op.sizeNotEquals
	operatorMap[constants.SizeLessThan] = op.sizeLessThan
	operatorMap[constants.SizeLessThanOrEqualTo] = op.sizeLessThanOrEqualTo
	operatorMap[constants.SizeGreaterThan] = op.sizeGreaterThan
	operatorMap[constants.SizeGreaterThanOrEqualTo] = op.sizeGreaterThanOrEqualTo

	op.operatorMap = operatorMap
	return &op
}
