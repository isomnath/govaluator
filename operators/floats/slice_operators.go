package floats

import (
	"reflect"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type FloatSlice struct {
	utils       operatorUtilities
	operatorMap map[string]descriptors.Executor
}

func (fso *FloatSlice) GetOperators() []string {
	var operators []string
	for k := range fso.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (fso *FloatSlice) GetExecutor(operatorName string) descriptors.Executor {
	return fso.operatorMap[operatorName]
}

func (fso *FloatSlice) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// unOrderedEquals To be read as slice val1 and val2 are equal ignoring order
func (fso *FloatSlice) unOrderedEquals(val1 interface{}, val2 interface{}) bool {
	val1FdMap, err := fso.utils.CastFloatSliceToFrequencyDistributionMap(val1)
	val2FdMap, err := fso.utils.CastFloatSliceToFrequencyDistributionMap(val2)
	if err != nil {
		return false
	}
	return reflect.DeepEqual(val1FdMap, val2FdMap)
}

// unOrderedNotEquals To be read as slice val1 and val2 are not equal ignoring order
func (fso *FloatSlice) unOrderedNotEquals(val1 interface{}, val2 interface{}) bool {
	return !fso.unOrderedEquals(val1, val2)
}

// orderedEquals To be read as slice val1 and val2 are equal considering order
func (fso *FloatSlice) orderedEquals(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := fso.utils.CastToFloatSlice(val1)
	bSlice2, err := fso.utils.CastToFloatSlice(val2)
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
func (fso *FloatSlice) orderedNotEquals(val1 interface{}, val2 interface{}) bool {
	return !fso.orderedEquals(val1, val2)
}

// anyLessThan To be read as any-of val1 less-than val2
func (fso *FloatSlice) anyLessThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := fso.utils.CastToFloatSlice(val1)
	bVal2, err2 := fso.utils.CastToFloatValue(val2)
	if err1 != nil || err2 != nil || len(bSlice1) < 1 {
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
func (fso *FloatSlice) noneLessThan(val1 interface{}, val2 interface{}) bool {
	return !fso.anyLessThan(val1, val2)
}

// anyLessThanOrEqualTo To be read as any-of val1 less-than-or-equal-to val2
func (fso *FloatSlice) anyLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := fso.utils.CastToFloatSlice(val1)
	bVal2, err2 := fso.utils.CastToFloatValue(val2)
	if err1 != nil || err2 != nil || len(bSlice1) < 1 {
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
func (fso *FloatSlice) noneLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	return !fso.anyLessThanOrEqualTo(val1, val2)
}

// allLessThan To be read as all-of val1 less-than val2
func (fso *FloatSlice) allLessThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := fso.utils.CastToFloatSlice(val1)
	bVal2, err2 := fso.utils.CastToFloatValue(val2)
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
func (fso *FloatSlice) allLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := fso.utils.CastToFloatSlice(val1)
	bVal2, err2 := fso.utils.CastToFloatValue(val2)
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
func (fso *FloatSlice) anyGreaterThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := fso.utils.CastToFloatSlice(val1)
	bVal2, err2 := fso.utils.CastToFloatValue(val2)
	if err1 != nil || err2 != nil || len(bSlice1) < 1 {
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
func (fso *FloatSlice) noneGreaterThan(val1 interface{}, val2 interface{}) bool {
	return !fso.anyGreaterThan(val1, val2)
}

// anyGreaterThanOrEqualTo To be read as any-of val1 greater-than-or-equal-to val2
func (fso *FloatSlice) anyGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := fso.utils.CastToFloatSlice(val1)
	bVal2, err2 := fso.utils.CastToFloatValue(val2)
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
func (fso *FloatSlice) noneGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	return !fso.anyGreaterThanOrEqualTo(val1, val2)
}

// allGreaterThan To be read as all-of val1 less-than val2
func (fso *FloatSlice) allGreaterThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := fso.utils.CastToFloatSlice(val1)
	bVal2, err2 := fso.utils.CastToFloatValue(val2)
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
func (fso *FloatSlice) allGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := fso.utils.CastToFloatSlice(val1)
	bVal2, err2 := fso.utils.CastToFloatValue(val2)
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
func (fso *FloatSlice) reversedBetween(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := fso.utils.CastToFloatSlice(val1)
	bVal2, err2 := fso.utils.CastToFloatValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	if len(bSlice1) < 2 || len(bSlice1) > 2 {
		return false
	}
	return bVal2 > bSlice1[0] && bVal2 < bSlice1[1]
}

// reversedNotBetween To be read as val2 is not within the range of values contained in val1  (boundaries not considered)
func (fso *FloatSlice) reversedNotBetween(val1 interface{}, val2 interface{}) bool {
	return !fso.reversedBetween(val1, val2)
}

// supersetOf To be read as val1 is a superset-of val2
func (fso *FloatSlice) supersetOf(val1 interface{}, val2 interface{}) bool {
	set1, err := fso.utils.CastFloatSliceToSet(val1)
	set2, err := fso.utils.CastFloatSliceToSet(val2)
	if err != nil {
		return false
	}
	return set1.IsSuperset(set2)
}

// notSupersetOf To be read as val1 is not a superset-of val2
func (fso *FloatSlice) notSupersetOf(val1 interface{}, val2 interface{}) bool {
	return !fso.supersetOf(val1, val2)
}

// subsetOf To be read as val1 is a subset-of val2
func (fso *FloatSlice) subsetOf(val1 interface{}, val2 interface{}) bool {
	set1, err := fso.utils.CastFloatSliceToSet(val1)
	set2, err := fso.utils.CastFloatSliceToSet(val2)
	if err != nil {
		return false
	}
	return set1.IsSubset(set2)
}

// notSubsetOf To be read as val1 is not a subset-of val2
func (fso *FloatSlice) notSubsetOf(val1 interface{}, val2 interface{}) bool {
	return !fso.subsetOf(val1, val2)
}

// intersection To be read as val1 intersects val2
func (fso *FloatSlice) intersection(val1 interface{}, val2 interface{}) bool {
	set1, err := fso.utils.CastFloatSliceToSet(val1)
	set2, err := fso.utils.CastFloatSliceToSet(val2)
	if err != nil {
		return false
	}
	intersectionSet := set1.Intersect(set2)
	return len(intersectionSet.ToSlice()) > 0
}

// NotIntersection To be read as val1 does not intersect val2
func (fso *FloatSlice) notIntersection(val1 interface{}, val2 interface{}) bool {
	return !fso.intersection(val1, val2)
}

// reversedAnyOf To be read as val2 is any-of val1
func (fso *FloatSlice) reversedAnyOf(val1 interface{}, val2 interface{}) bool {
	bVal2, err1 := fso.utils.CastToFloatValue(val2)
	bSlice1, err2 := fso.utils.CastToFloatSlice(val1)
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
func (fso *FloatSlice) reversedNoneOf(val1 interface{}, val2 interface{}) bool {
	return !fso.reversedAnyOf(val1, val2)
}

// sizeEquals To be read as size of val1 is equal-to val2
func (fso *FloatSlice) sizeEquals(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := fso.utils.CastToFloatSlice(val1)
	bVal2, err := fso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) == bVal2
}

// sizeNotEquals To be read as size of val1 is not-equal-to val2
func (fso *FloatSlice) sizeNotEquals(val1 interface{}, val2 interface{}) bool {
	return !fso.sizeEquals(val1, val2)
}

// sizeLessThan To be read as size of val1 is less-than val2
func (fso *FloatSlice) sizeLessThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := fso.utils.CastToFloatSlice(val1)
	bVal2, err := fso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) < bVal2
}

// sizeLessThanOrEqualTo To be read as size of val1 is less-than-or-equal-to val2
func (fso *FloatSlice) sizeLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := fso.utils.CastToFloatSlice(val1)
	bVal2, err := fso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) <= bVal2
}

// sizeGreaterThan To be read as size of val1 is greater-than val2
func (fso *FloatSlice) sizeGreaterThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := fso.utils.CastToFloatSlice(val1)
	bVal2, err := fso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) > bVal2
}

// sizeGreaterThanOrEqualTo To be read as size of val1 is greater-than-or-equal-to val2
func (fso *FloatSlice) sizeGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := fso.utils.CastToFloatSlice(val1)
	bVal2, err := fso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) >= bVal2
}

func InitializeSliceOperators(operatorUtilities operatorUtilities) *FloatSlice {
	op := FloatSlice{
		utils: operatorUtilities,
	}

	f := make(map[string]descriptors.Executor, 0)
	f[constants.UnorderedEquals] = op.unOrderedEquals
	f[constants.UnorderedNotEquals] = op.unOrderedNotEquals
	f[constants.OrderedEquals] = op.orderedEquals
	f[constants.OrderedNotEquals] = op.orderedNotEquals

	f[constants.AnyLessThan] = op.anyLessThan
	f[constants.AnyLessThanOrEqualTo] = op.anyLessThanOrEqualTo
	f[constants.NoneLessThan] = op.noneLessThan
	f[constants.NoneLessThanOrEqualTo] = op.noneLessThanOrEqualTo
	f[constants.AllLessThan] = op.allLessThan
	f[constants.AllLessThanOrEqualTo] = op.allLessThanOrEqualTo
	f[constants.AnyGreaterThan] = op.anyGreaterThan
	f[constants.AnyGreaterThanOrEqualTo] = op.anyGreaterThanOrEqualTo
	f[constants.NoneGreaterThan] = op.noneGreaterThan
	f[constants.NoneGreaterThanOrEqualTo] = op.noneGreaterThanOrEqualTo
	f[constants.AllGreaterThan] = op.allGreaterThan
	f[constants.AllGreaterThanOrEqualTo] = op.allGreaterThanOrEqualTo
	f[constants.ReversedBetween] = op.reversedBetween
	f[constants.ReversedNotBetween] = op.reversedNotBetween

	f[constants.SupersetOf] = op.supersetOf
	f[constants.NotSupersetOf] = op.notSupersetOf
	f[constants.SubsetOf] = op.subsetOf
	f[constants.NotSubsetOf] = op.notSubsetOf
	f[constants.Intersection] = op.intersection
	f[constants.NotIntersection] = op.notIntersection
	f[constants.ReversedAnyOf] = op.reversedAnyOf
	f[constants.ReversedNoneOf] = op.reversedNoneOf

	f[constants.SizeEquals] = op.sizeEquals
	f[constants.SizeNotEquals] = op.sizeNotEquals
	f[constants.SizeLessThan] = op.sizeLessThan
	f[constants.SizeLessThanOrEqualTo] = op.sizeLessThanOrEqualTo
	f[constants.SizeGreaterThan] = op.sizeGreaterThan
	f[constants.SizeGreaterThanOrEqualTo] = op.sizeGreaterThanOrEqualTo

	op.operatorMap = f
	return &op
}
