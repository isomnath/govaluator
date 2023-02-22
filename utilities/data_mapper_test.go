package utilities

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DataMapperTestSuite struct {
	suite.Suite
	utils *Utilities
}

func (suite *DataMapperTestSuite) SetupTest() {
	suite.utils = InitializeUtilities()
}

func (suite *DataMapperTestSuite) TestFlattenShouldReturnFlattenedMapStringInterface() {
	mapStringInterfaceValue := map[string]interface{}{
		"parent_1": map[string]interface{}{
			"key_1": 123,
			"key_2": "value",
		},
		"parent_2": map[string]interface{}{
			"key_1": 4536,
			"key_2": []string{"value_1", "value_2"},
			"key_3": true,
		},
		"parent_3": map[string]interface{}{
			"key_1": []interface{}{
				map[string]interface{}{
					"child_key_1": 12,
					"child_key_2": "value_1",
					"child_key_3": "value_2",
				},
				map[string]interface{}{
					"child_key_1": 21,
					"child_key_2": "value_1",
					"child_key_3": "value_2",
				},
			},
		},
	}
	expectedValue := map[string]interface{}{
		"parent_1.key_1":               123,
		"parent_1.key_2":               "value",
		"parent_2.key_1":               4536,
		"parent_2.key_2":               []string{"value_1", "value_2"},
		"parent_2.key_3":               true,
		"parent_3.key_1.0.child_key_1": 12,
		"parent_3.key_1.0.child_key_2": "value_1",
		"parent_3.key_1.0.child_key_3": "value_2",
		"parent_3.key_1.1.child_key_1": 21,
		"parent_3.key_1.1.child_key_2": "value_1",
		"parent_3.key_1.1.child_key_3": "value_2",
	}

	value := suite.utils.Flatten(mapStringInterfaceValue)
	suite.Equal(expectedValue, value)
}

func (suite *DataMapperTestSuite) TestUnFlattenShouldReturnUnFlattenedMapStringInterface() {
	mapStringInterfaceValue := map[string]interface{}{
		"parent_1.key_1": 123,
		"parent_1.key_2": "value",
		"parent_2.key_1": 4536,
		"parent_2.key_2": []string{"value_1", "value_2"},
		"parent_2.key_3": true,
	}
	expectedValue := map[string]interface{}{
		"parent_1": map[string]interface{}{
			"key_1": 123,
			"key_2": "value",
		},
		"parent_2": map[string]interface{}{
			"key_1": 4536,
			"key_2": []string{"value_1", "value_2"},
			"key_3": true,
		},
	}

	value := suite.utils.UnFlatten(mapStringInterfaceValue)
	suite.Equal(expectedValue, value)
}

func TestDataMapperTestSuite(t *testing.T) {
	suite.Run(t, new(DataMapperTestSuite))
}
