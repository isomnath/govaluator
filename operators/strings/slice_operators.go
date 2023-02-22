package strings

import (
	"reflect"
	"regexp"

	"github.com/isomnath/govaluator/constants"
	"github.com/isomnath/govaluator/descriptors"
)

type StringSlice struct {
	utils       operatorUtilities
	operatorMap map[string]descriptors.Executor
}

func (sso *StringSlice) GetOperators() []string {
	var operators []string
	for k := range sso.operatorMap {
		operators = append(operators, k)
	}
	return operators
}

func (sso *StringSlice) GetExecutor(operatorName string) descriptors.Executor {
	return sso.operatorMap[operatorName]
}

func (sso *StringSlice) ExecuteOperator(fn descriptors.Executor, val1 interface{}, val2 interface{}) bool {
	return fn(val1, val2)
}

// unOrderedEquals To be read as slice val1 and val2 are equal ignoring order
func (sso *StringSlice) unOrderedEquals(val1 interface{}, val2 interface{}) bool {
	val1FdMap, err := sso.utils.CastStringSliceToFrequencyDistributionMap(val1)
	val2FdMap, err := sso.utils.CastStringSliceToFrequencyDistributionMap(val2)
	if err != nil {
		return false
	}
	return reflect.DeepEqual(val1FdMap, val2FdMap)
}

// unOrderedNotEquals To be read as slice val1 and val2 are not equal ignoring order
func (sso *StringSlice) unOrderedNotEquals(val1 interface{}, val2 interface{}) bool {
	return !sso.unOrderedEquals(val1, val2)
}

// orderedEquals To be read as slice val1 and val2 are equal considering order
func (sso *StringSlice) orderedEquals(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := sso.utils.CastToStringSlice(val1)
	bSlice2, err := sso.utils.CastToStringSlice(val2)
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
func (sso *StringSlice) orderedNotEquals(val1 interface{}, val2 interface{}) bool {
	return !sso.orderedEquals(val1, val2)
}

// supersetOf To be read as val1 is a superset-of val2
func (sso *StringSlice) supersetOf(val1 interface{}, val2 interface{}) bool {
	set1, err := sso.utils.CastStringSliceToSet(val1)
	set2, err := sso.utils.CastStringSliceToSet(val2)
	if err != nil {
		return false
	}
	return set1.IsSuperset(set2)
}

// notSupersetOf To be read as val1 is not a superset-of val2
func (sso *StringSlice) notSupersetOf(val1 interface{}, val2 interface{}) bool {
	return !sso.supersetOf(val1, val2)
}

// subsetOf To be read as val1 is a subset-of val2
func (sso *StringSlice) subsetOf(val1 interface{}, val2 interface{}) bool {
	set1, err := sso.utils.CastStringSliceToSet(val1)
	set2, err := sso.utils.CastStringSliceToSet(val2)
	if err != nil {
		return false
	}
	return set1.IsSubset(set2)
}

// notSubsetOf To be read as val1 is not a subset-of val2
func (sso *StringSlice) notSubsetOf(val1 interface{}, val2 interface{}) bool {
	return !sso.subsetOf(val1, val2)
}

// intersection To be read as val1 intersects val2
func (sso *StringSlice) intersection(val1 interface{}, val2 interface{}) bool {
	set1, err := sso.utils.CastStringSliceToSet(val1)
	set2, err := sso.utils.CastStringSliceToSet(val2)
	if err != nil {
		return false
	}
	intersectionSet := set1.Intersect(set2)
	return len(intersectionSet.ToSlice()) > 0
}

// NotIntersection To be read as val1 does not intersect val2
func (sso *StringSlice) notIntersection(val1 interface{}, val2 interface{}) bool {
	return !sso.intersection(val1, val2)
}

// reversedAnyOf To be read as val2 is any-of val1
func (sso *StringSlice) reversedAnyOf(val1 interface{}, val2 interface{}) bool {
	bVal2, err1 := sso.utils.CastToStringValue(val2)
	bSlice1, err2 := sso.utils.CastToStringSlice(val1)
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
func (sso *StringSlice) reversedNoneOf(val1 interface{}, val2 interface{}) bool {
	return !sso.reversedAnyOf(val1, val2)
}

// sizeEquals To be read as size of val1 is equal-to val2
func (sso *StringSlice) sizeEquals(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := sso.utils.CastToStringSlice(val1)
	bVal2, err := sso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) == bVal2
}

// sizeNotEquals To be read as size of val1 is not-equal-to val2
func (sso *StringSlice) sizeNotEquals(val1 interface{}, val2 interface{}) bool {
	return !sso.sizeEquals(val1, val2)
}

// sizeLessThan To be read as size of val1 is less-than val2
func (sso *StringSlice) sizeLessThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := sso.utils.CastToStringSlice(val1)
	bVal2, err := sso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) < bVal2
}

// sizeLessThanOrEqualTo To be read as size of val1 is less-than-or-equal-to val2
func (sso *StringSlice) sizeLessThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := sso.utils.CastToStringSlice(val1)
	bVal2, err := sso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) <= bVal2
}

// sizeGreaterThan To be read as size of val1 is greater-than val2
func (sso *StringSlice) sizeGreaterThan(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := sso.utils.CastToStringSlice(val1)
	bVal2, err := sso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) > bVal2
}

// sizeGreaterThanOrEqualTo To be read as size of val1 is greater-than-or-equal-to val2
func (sso *StringSlice) sizeGreaterThanOrEqualTo(val1 interface{}, val2 interface{}) bool {
	bSlice1, err := sso.utils.CastToStringSlice(val1)
	bVal2, err := sso.utils.CastToIntegerValue(val2)
	if err != nil {
		return false
	}
	return int64(len(bSlice1)) >= bVal2
}

// anyRegexMatch To be read as any of the elements in val1 matches regex val2
func (sso *StringSlice) anyRegexMatch(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := sso.utils.CastToStringSlice(val1)
	bVal2, err2 := sso.utils.CastToStringValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal1 := range bSlice1 {
		matched, err := regexp.MatchString(bVal2, bVal1)
		if err != nil {
			flag = false
		}
		if matched {
			return true
		}
	}
	return flag
}

// noneRegexMatch To be read as none of the elements in val1 matches regex val2
func (sso *StringSlice) noneRegexMatch(val1 interface{}, val2 interface{}) bool {
	return !sso.anyRegexMatch(val1, val2)
}

// allRegexMatch To be read as all the elements in val1 matches regex val2
func (sso *StringSlice) allRegexMatch(val1 interface{}, val2 interface{}) bool {
	bSlice1, err1 := sso.utils.CastToStringSlice(val1)
	bVal2, err2 := sso.utils.CastToStringValue(val2)
	if err1 != nil || err2 != nil {
		return false
	}
	flag := false
	for _, bVal1 := range bSlice1 {
		matched, err := regexp.MatchString(bVal2, bVal1)
		if err != nil || !matched {
			return false
		}
		if matched {
			flag = true
		}
	}
	return flag
}

func InitializeSliceOperators(operatorUtilities operatorUtilities) *StringSlice {
	op := StringSlice{
		utils: operatorUtilities,
	}

	f := make(map[string]descriptors.Executor, 0)
	f[constants.UnorderedEquals] = op.unOrderedEquals
	f[constants.UnorderedNotEquals] = op.unOrderedNotEquals
	f[constants.OrderedEquals] = op.orderedEquals
	f[constants.OrderedNotEquals] = op.orderedNotEquals

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

	f[constants.AnyRegexMatch] = op.anyRegexMatch
	f[constants.NoneRegexMatch] = op.noneRegexMatch
	f[constants.AllRegexMatch] = op.allRegexMatch

	op.operatorMap = f
	return &op
}
