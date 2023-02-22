package exitonstrategies

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/isomnath/govaluator/models"
)

type LastFalse struct {
	suite.Suite
	strategy ExitOnStrategy
}

func (suite *LastFalse) SetupTest() {
	suite.strategy = initializeLastFalse()
}

func (suite *LastFalse) TestGetResultValue() {
	type args struct {
		results []models.TransientResult
	}
	tests := []struct {
		name        string
		args        args
		expectation interface{}
	}{
		{
			name: "Value Return",
			args: args{
				results: []models.TransientResult{
					{
						Flag:   true,
						Result: "A",
					},
					{
						Flag:   false,
						Result: "B",
					},
					{
						Flag:   true,
						Result: "C",
					},
					{
						Flag:   false,
						Result: "D",
					},
				},
			},
			expectation: "D",
		},
		{
			name: "Nil Return",
			args: args{
				results: []models.TransientResult{
					{
						Flag:   true,
						Result: "A",
					},
					{
						Flag:   true,
						Result: "B",
					},
					{
						Flag:   true,
						Result: "C",
					},
					{
						Flag:   true,
						Result: "D",
					},
				},
			},
			expectation: nil,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if got := suite.strategy.GetResultValue(tt.args.results); !reflect.DeepEqual(got, tt.expectation) {
				suite.T().Errorf("filter() = %v, expectation %v", got, tt.expectation)
			}
		})
	}
}

func TestLastFalse(t *testing.T) {
	suite.Run(t, new(LastFalse))
}
