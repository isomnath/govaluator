package utilities

import (
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set"

	"github.com/stretchr/testify/suite"
)

type OperatorUtilitiesTestSuite struct {
	suite.Suite
	operatorUtilities *OperatorUtilities
}

func (suite *OperatorUtilitiesTestSuite) SetupSuite() {
	suite.operatorUtilities = InitializeOperatorUtilities()
}

func (suite *OperatorUtilitiesTestSuite) TestBooleanCastToValueShouldReturnError() {
	var inter interface{}
	inter = "test"

	val, err := suite.operatorUtilities.CastToBooleanValue(inter)
	suite.Equal(false, val)
	suite.Equal("interface is not of type: bool", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestBooleanCastToValueShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = true

	val, err := suite.operatorUtilities.CastToBooleanValue(inter)
	suite.Equal(true, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestBooleanCastToSliceShouldReturnError() {
	var inter interface{}
	inter = []string{"test"}

	val, err := suite.operatorUtilities.CastToBooleanSlice(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []bool", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestBooleanCastToSliceShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []bool{true, false, true, false, false}

	val, err := suite.operatorUtilities.CastToBooleanSlice(inter)
	suite.Equal([]bool{true, false, true, false, false}, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastBooleanSliceToFrequencyDistributionMapShouldReturnError() {
	var inter interface{}
	inter = []string{"test"}

	val, err := suite.operatorUtilities.CastBooleanSliceToFrequencyDistributionMap(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []bool", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestCastBooleanSliceToFrequencyDistributionMapShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []bool{true, true, false, false, true, false, false, false}

	expectedFdMap := map[bool]int64{
		true:  3,
		false: 5,
	}
	val, err := suite.operatorUtilities.CastBooleanSliceToFrequencyDistributionMap(inter)
	suite.Equal(expectedFdMap, val)
	suite.NoError(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastBooleanSliceToSetShouldReturnError() {
	var inter interface{}
	inter = []string{"true", "false"}

	val, err := suite.operatorUtilities.CastBooleanSliceToSet(inter)
	suite.Empty(val)
	suite.Equal("interface is not of type: []bool", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestCastBooleanSliceToSetShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []bool{true, true, false, false, true, false, false, false}

	expectedSet := mapset.NewSet()
	expectedSet.Add(true)
	expectedSet.Add(false)
	val, err := suite.operatorUtilities.CastBooleanSliceToSet(inter)
	suite.Equal(expectedSet, val)
	suite.NoError(err)
}

func (suite *OperatorUtilitiesTestSuite) TestStringCastToValueShouldReturnError() {
	var inter interface{}
	inter = true

	val, err := suite.operatorUtilities.CastToStringValue(inter)
	suite.Equal("", val)
	suite.Equal("interface is not of type: string", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestStringCastToValueShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = "test_one"

	val, err := suite.operatorUtilities.CastToStringValue(inter)
	suite.Equal("test_one", val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestStringCastToSliceShouldReturnError() {
	var inter interface{}
	inter = []int{1, 2, 3}

	val, err := suite.operatorUtilities.CastToStringSlice(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []string", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestStringCastToSliceShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []string{"test_one", "test_two", "test_three"}

	val, err := suite.operatorUtilities.CastToStringSlice(inter)
	suite.Equal([]string{"test_one", "test_two", "test_three"}, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastStringSliceToFrequencyDistributionMapShouldReturnError() {
	var inter interface{}
	inter = []bool{true, false}

	val, err := suite.operatorUtilities.CastStringSliceToFrequencyDistributionMap(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []string", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestCastStringSliceToFrequencyDistributionMapShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []string{"test_one", "test_two", "test_one"}

	expectedFdMap := map[string]int64{
		"test_one": 2,
		"test_two": 1,
	}
	val, err := suite.operatorUtilities.CastStringSliceToFrequencyDistributionMap(inter)
	suite.Equal(expectedFdMap, val)
	suite.NoError(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastStringSliceToSetShouldReturnError() {
	var inter interface{}
	inter = []int{1, 2}

	val, err := suite.operatorUtilities.CastStringSliceToSet(inter)
	suite.Empty(val)
	suite.Equal("interface is not of type: []string", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestCastStringSliceToSetShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []string{"test_one", "test_two", "test_one"}

	expectedSet := mapset.NewSet()
	expectedSet.Add("test_one")
	expectedSet.Add("test_two")
	val, err := suite.operatorUtilities.CastStringSliceToSet(inter)
	suite.Equal(expectedSet, val)
	suite.NoError(err)
}

func (suite *OperatorUtilitiesTestSuite) TestFloat32CastToValueShouldReturnError() {
	var inter interface{}
	inter = true

	val, err := suite.operatorUtilities.CastToFloatValue(inter)
	suite.Equal(0.0, val)
	suite.Equal("interface is not of type: float", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestFloat32CastShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = float32(12)

	val, err := suite.operatorUtilities.CastToFloatValue(inter)
	suite.Equal(12.0, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestFloat64CastToValueShouldReturnError() {
	var inter interface{}
	inter = true

	val, err := suite.operatorUtilities.CastToFloatValue(inter)
	suite.Equal(0.0, val)
	suite.Equal("interface is not of type: float", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestFloat64CastShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = float64(12)

	val, err := suite.operatorUtilities.CastToFloatValue(inter)
	suite.Equal(12.0, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestFloat32CastToSliceShouldReturnError() {
	var inter interface{}
	inter = true

	val, err := suite.operatorUtilities.CastToFloatSlice(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []float", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestFloat32CastToSliceShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []float32{float32(12), float32(16)}

	val, err := suite.operatorUtilities.CastToFloatSlice(inter)
	suite.Equal([]float64{12.0, 16.0}, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestFloat64CastToSliceShouldReturnError() {
	var inter interface{}
	inter = true

	val, err := suite.operatorUtilities.CastToFloatSlice(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []float", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestFloat64CastToSliceShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []float64{float64(12), float64(16)}

	val, err := suite.operatorUtilities.CastToFloatSlice(inter)
	suite.Equal([]float64{12.0, 16.0}, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastFloatSliceToFrequencyDistributionMapShouldReturnError() {
	var inter interface{}
	inter = true

	val, err := suite.operatorUtilities.CastFloatSliceToFrequencyDistributionMap(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []float", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestCastFloat32SliceToFrequencyDistributionMapShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []float32{float32(12), float32(12), float32(134)}

	expectedFdMap := map[float64]int64{
		float64(12):  2,
		float64(134): 1,
	}
	val, err := suite.operatorUtilities.CastFloatSliceToFrequencyDistributionMap(inter)
	suite.Equal(expectedFdMap, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastFloat64SliceToFrequencyDistributionMapShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []float64{float64(12), float64(12), float64(134)}

	expectedFdMap := map[float64]int64{
		float64(12):  2,
		float64(134): 1,
	}
	val, err := suite.operatorUtilities.CastFloatSliceToFrequencyDistributionMap(inter)
	suite.Equal(expectedFdMap, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastFloat32SliceToSetShouldReturnError() {
	var inter interface{}
	inter = []int32{1, 2}

	val, err := suite.operatorUtilities.CastFloatSliceToSet(inter)
	suite.Empty(val)
	suite.Equal("interface is not of type: []float", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestCastFloat32SliceToSetShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []float32{1, 2, 4, 2}

	expectedSet := mapset.NewSet()
	expectedSet.Add(float64(1))
	expectedSet.Add(float64(2))
	expectedSet.Add(float64(4))
	val, err := suite.operatorUtilities.CastFloatSliceToSet(inter)
	suite.Equal(expectedSet, val)
	suite.NoError(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastFloat64SliceToSetShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []float64{1, 2, 4, 2}

	expectedSet := mapset.NewSet()
	expectedSet.Add(float64(1))
	expectedSet.Add(float64(2))
	expectedSet.Add(float64(4))
	val, err := suite.operatorUtilities.CastFloatSliceToSet(inter)
	suite.Equal(expectedSet, val)
	suite.NoError(err)
}

func (suite *OperatorUtilitiesTestSuite) TestIntCastToValueShouldReturnError() {
	var inter interface{}
	inter = true

	val, err := suite.operatorUtilities.CastToIntegerValue(inter)
	suite.Equal(int64(0), val)
	suite.Equal("interface is not of type: int", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestIntCastToValueShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = 10

	val, err := suite.operatorUtilities.CastToIntegerValue(inter)
	suite.Equal(int64(10), val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt8CastToValueShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = int8(10)

	val, err := suite.operatorUtilities.CastToIntegerValue(inter)
	suite.Equal(int64(10), val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt16CastToValueShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = int16(10)

	val, err := suite.operatorUtilities.CastToIntegerValue(inter)
	suite.Equal(int64(10), val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt32CastToValueShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = int32(10)

	val, err := suite.operatorUtilities.CastToIntegerValue(inter)
	suite.Equal(int64(10), val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt64CastToValueShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = int64(10)

	val, err := suite.operatorUtilities.CastToIntegerValue(inter)
	suite.Equal(int64(10), val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestIntCastToSliceShouldReturnError() {
	var inter interface{}
	inter = "testOne"

	val, err := suite.operatorUtilities.CastToIntegerSlice(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []int", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestIntCastToSliceShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int{1, 2, 3}

	val, err := suite.operatorUtilities.CastToIntegerSlice(inter)
	suite.Equal([]int64{1, 2, 3}, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt8CastToSliceShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int8{1, 2, 3}

	val, err := suite.operatorUtilities.CastToIntegerSlice(inter)
	suite.Equal([]int64{1, 2, 3}, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt16CastToSliceShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int16{1, 2, 3}

	val, err := suite.operatorUtilities.CastToIntegerSlice(inter)
	suite.Equal([]int64{1, 2, 3}, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt32CastToSliceShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int32{1, 2, 3}

	val, err := suite.operatorUtilities.CastToIntegerSlice(inter)
	suite.Equal([]int64{1, 2, 3}, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt64CastToSliceShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int64{1, 2, 3}

	val, err := suite.operatorUtilities.CastToIntegerSlice(inter)
	suite.Equal([]int64{1, 2, 3}, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastIntegerSliceToFrequencyDistributionMapShouldReturnError() {
	var inter interface{}
	inter = "testOne"

	val, err := suite.operatorUtilities.CastIntegerSliceToFrequencyDistributionMap(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []int", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestIntCastIntegerSliceToFrequencyDistributionMapShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int{1, 1, 3}

	expectedFdMap := map[int64]int64{
		1: 2,
		3: 1,
	}
	val, err := suite.operatorUtilities.CastIntegerSliceToFrequencyDistributionMap(inter)
	suite.Equal(expectedFdMap, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt8CastIntegerSliceToFrequencyDistributionMapShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int8{1, 1, 3}

	expectedFdMap := map[int64]int64{
		1: 2,
		3: 1,
	}
	val, err := suite.operatorUtilities.CastIntegerSliceToFrequencyDistributionMap(inter)
	suite.Equal(expectedFdMap, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt16CastIntegerSliceToFrequencyDistributionMapShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int16{1, 1, 3}

	expectedFdMap := map[int64]int64{
		1: 2,
		3: 1,
	}
	val, err := suite.operatorUtilities.CastIntegerSliceToFrequencyDistributionMap(inter)
	suite.Equal(expectedFdMap, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt32CastIntegerSliceToFrequencyDistributionMapShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int32{1, 1, 3}

	expectedFdMap := map[int64]int64{
		1: 2,
		3: 1,
	}
	val, err := suite.operatorUtilities.CastIntegerSliceToFrequencyDistributionMap(inter)
	suite.Equal(expectedFdMap, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt64CastIntegerSliceToFrequencyDistributionMapShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int64{1, 1, 3}

	expectedFdMap := map[int64]int64{
		1: 2,
		3: 1,
	}
	val, err := suite.operatorUtilities.CastIntegerSliceToFrequencyDistributionMap(inter)
	suite.Equal(expectedFdMap, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestCastIntegerSliceToSetShouldReturnError() {
	var inter interface{}
	inter = "testOne"

	val, err := suite.operatorUtilities.CastIntegerSliceToSet(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []int", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestIntCastIntegerSliceToSetShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int{1, 1, 3}

	expectedSet := mapset.NewSet()
	expectedSet.Add(int64(1))
	expectedSet.Add(int64(3))
	val, err := suite.operatorUtilities.CastIntegerSliceToSet(inter)
	suite.Equal(expectedSet, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt8CastIntegerSliceToSetShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int8{1, 1, 3}

	expectedSet := mapset.NewSet()
	expectedSet.Add(int64(1))
	expectedSet.Add(int64(3))
	val, err := suite.operatorUtilities.CastIntegerSliceToSet(inter)
	suite.Equal(expectedSet, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt16CastIntegerSliceToSetShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int16{1, 1, 3}

	expectedSet := mapset.NewSet()
	expectedSet.Add(int64(1))
	expectedSet.Add(int64(3))
	val, err := suite.operatorUtilities.CastIntegerSliceToSet(inter)
	suite.Equal(expectedSet, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt32CastIntegerSliceToSetShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int32{1, 1, 3}

	expectedSet := mapset.NewSet()
	expectedSet.Add(int64(1))
	expectedSet.Add(int64(3))
	val, err := suite.operatorUtilities.CastIntegerSliceToSet(inter)
	suite.Equal(expectedSet, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInt64CastIntegerSliceToSetShouldReturnValueSuccessfully() {
	var inter interface{}
	inter = []int64{1, 1, 3}

	expectedSet := mapset.NewSet()
	expectedSet.Add(int64(1))
	expectedSet.Add(int64(3))
	val, err := suite.operatorUtilities.CastIntegerSliceToSet(inter)
	suite.Equal(expectedSet, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestInitializersShouldReturnInterfaces() {
	operatorUtilities := InitializeOperatorUtilities()

	suite.Equal(suite.operatorUtilities, operatorUtilities)
}

func (suite *OperatorUtilitiesTestSuite) TestTimeCastToValueShouldReturnError() {
	var inter interface{}
	inter = "test"

	val, err := suite.operatorUtilities.CastToTimeValue(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: times", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestTimeCastToValueShouldReturnValueSuccessfully() {
	t := time.Now()

	var inter1 interface{}
	inter1 = t

	val, err := suite.operatorUtilities.CastToTimeValue(inter1)
	suite.Equal(&t, val)
	suite.Nil(err)

	var inter2 interface{}
	inter2 = &t
	val, err = suite.operatorUtilities.CastToTimeValue(inter2)
	suite.Equal(&t, val)
	suite.Nil(err)
}

func (suite *OperatorUtilitiesTestSuite) TestTimeCastToSliceShouldReturnError() {
	var inter interface{}
	inter = []string{"test"}

	val, err := suite.operatorUtilities.CastToTimeSlice(inter)
	suite.Nil(val)
	suite.Equal("interface is not of type: []times", err.Error())
}

func (suite *OperatorUtilitiesTestSuite) TestTimeCastToSliceShouldReturnValueSuccessfully() {
	t := time.Now()

	var inter1 interface{}
	inter1 = []time.Time{t}

	val, err := suite.operatorUtilities.CastToTimeSlice(inter1)
	suite.Equal([]*time.Time{&t}, val)
	suite.Nil(err)

	var inter2 interface{}
	inter2 = []*time.Time{&t}
	val, err = suite.operatorUtilities.CastToTimeSlice(inter2)
	suite.Equal([]*time.Time{&t}, val)
	suite.Nil(err)
}

func TestOperatorUtilitiesTestSuite(t *testing.T) {
	suite.Run(t, new(OperatorUtilitiesTestSuite))
}
