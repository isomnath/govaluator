package booleans

import (
	"reflect"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type BooleanSlice struct {
	utils       operatorUtilities
	operatorMap map[string]descriptors.Executor
}

func (bso *BooleanSlice) GetOperators() []string {
	var operators []string
	for k := range bso.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (bso *BooleanSlice) GetExecutor(operatorName string) descriptors.Executor {
	return bso.operatorMap[operatorName]
}

func (bso *BooleanSlice) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// unOrderedEquals To be read as slice val1 and val2 are equal ignoring order
func (bso *BooleanSlice) unOrderedEquals(val1 interface{}, val2 interface{}) bool {
	val1FdMap, err := bso.utils.CastBooleanSliceToFrequencyDistributionMap(val1)
	val2FdMap, err := bso.utils.CastBooleanSliceToFrequencyDistributionMap(val2)
	if err != nil {
		return false
	}
	return reflect.DeepEqual(val1FdMap, val2FdMap)
}

// unOrderedNotEquals To be read as slice val1 and val2 are not equal ignoring order
func (bso *BooleanSlice) unOrderedNotEquals(val1 interface{}, val2 interface{}) bool {
	return !bso.unOrderedEquals(val1, val2)
}

// orderedEquals To be read as slice val1 and val2 are equal considering order
func (bso *BooleanSlice) orderedEquals(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := bso.utils.CastToBooleanSlice(val1)
	bSlice2, err := bso.utils.CastToBooleanSlice(val2)
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
func (bso *BooleanSlice) orderedNotEquals(val1 interface{}, val2 interface{}) bool {
	return !bso.orderedEquals(val1, val2)
}

// supersetOf To be read as val1 is a superset-of val2
func (bso *BooleanSlice) supersetOf(val1 interface{}, val2 interface{}) bool {
	set1, err := bso.utils.CastBooleanSliceToSet(val1)
	set2, err := bso.utils.CastBooleanSliceToSet(val2)
	if err != nil {
		return false
	}
	return set1.IsSuperset(set2)
}

// notSupersetOf To be read as val1 is not a superset-of val2
func (bso *BooleanSlice) notSupersetOf(val1 interface{}, val2 interface{}) bool {
	return !bso.supersetOf(val1, val2)
}

// subsetOf To be read as val1 is a subset-of val2
func (bso *BooleanSlice) subsetOf(val1 interface{}, val2 interface{}) bool {
	set1, err := bso.utils.CastBooleanSliceToSet(val1)
	set2, err := bso.utils.CastBooleanSliceToSet(val2)
	if err != nil {
		return false
	}
	return set1.IsSubset(set2)
}

// notSubsetOf To be read as val1 is not a subset-of val2
func (bso *BooleanSlice) notSubsetOf(val1 interface{}, val2 interface{}) bool {
	return !bso.subsetOf(val1, val2)
}

// intersection To be read as val1 intersects val2
func (bso *BooleanSlice) intersection(val1 interface{}, val2 interface{}) bool {
	set1, err := bso.utils.CastBooleanSliceToSet(val1)
	set2, err := bso.utils.CastBooleanSliceToSet(val2)
	if err != nil {
		return false
	}
	intersectionSet := set1.Intersect(set2)
	return len(intersectionSet.ToSlice()) > 0
}

// notIntersection To be read as val1 does not intersect val2
func (bso *BooleanSlice) notIntersection(val1 interface{}, val2 interface{}) bool {
	return !bso.intersection(val1, val2)
}

// reversedAnyOf To be read as val2 is any-of val1
func (bso *BooleanSlice) reversedAnyOf(val1 interface{}, val2 interface{}) bool {
	bVal2, err1 := bso.utils.CastToBooleanValue(val2)
	bSlice1, err2 := bso.utils.CastToBooleanSlice(val1)
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
func (bso *BooleanSlice) reversedNoneOf(val1 interface{}, val2 interface{}) bool {
	return !bso.reversedAnyOf(val1, val2)
}

// sizeEquals To be read as size of val1 is equal-to val2
func (bso *BooleanSlice) sizeEquals(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := bso.utils.CastToBooleanSlice(val1)
	bVal2, err := bso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) == bVal2
}

// sizeNotEquals To be read as size of val1 is not-equal-to val2
func (bso *BooleanSlice) sizeNotEquals(val1 interface{}, val2 interface{}) bool {
	return !bso.sizeEquals(val1, val2)
}

// sizeLessThan To be read as size of val1 is less-than val2
func (bso *BooleanSlice) sizeLessThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := bso.utils.CastToBooleanSlice(val1)
	bVal2, err := bso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) < bVal2
}

// sizeLessThanOrEqualTo To be read as size of val1 is less-than-or-equal-to val2
func (bso *BooleanSlice) sizeLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := bso.utils.CastToBooleanSlice(val1)
	bVal2, err := bso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) <= bVal2

}

// sizeGreaterThan To be read as size of val1 is greater-than val2
func (bso *BooleanSlice) sizeGreaterThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := bso.utils.CastToBooleanSlice(val1)
	bVal2, err := bso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) > bVal2
}

// sizeGreaterThanOrEqualTo To be read as size of val1 is greater-than-or-equal-to val2
func (bso *BooleanSlice) sizeGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := bso.utils.CastToBooleanSlice(val1)
	bVal2, err := bso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) >= bVal2
}

func InitializeSliceOperators(operatorUtilities operatorUtilities) *BooleanSlice {
	op := BooleanSlice{
		utils: operatorUtilities,
	}

	operatorMap := make(map[string]descriptors.Executor, 0)
	operatorMap[constants.UnorderedEquals] = op.unOrderedEquals
	operatorMap[constants.UnorderedNotEquals] = op.unOrderedNotEquals
	operatorMap[constants.OrderedEquals] = op.orderedEquals
	operatorMap[constants.OrderedNotEquals] = op.orderedNotEquals

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
