package exitonstrategies

import (
	"reflect"
	"testing"

	"github.com/isomnath/govaluator/models"
)

func Test_filter(t *testing.T) {
	type args struct {
		flag    bool
		results []models.TransientResult
	}
	tests := []struct {
		name        string
		args        args
		expectation []models.TransientResult
	}{
		{
			name: "Test - False Flag",
			args: args{
				flag: false,
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
			expectation: []models.TransientResult{
				{
					Flag:   false,
					Result: "B",
				},
				{
					Flag:   false,
					Result: "D",
				},
			},
		},
		{
			name: "Test - True Flag",
			args: args{
				flag: true,
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
			expectation: []models.TransientResult{
				{
					Flag:   true,
					Result: "A",
				},
				{
					Flag:   true,
					Result: "C",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filter(tt.args.flag, tt.args.results); !reflect.DeepEqual(got, tt.expectation) {
				t.Errorf("filter() = %v, expectation %v", got, tt.expectation)
			}
		})
	}
}
