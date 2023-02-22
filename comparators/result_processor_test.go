package comparators

import (
	"reflect"
	"testing"

	"github.com/isomnath/govaluator/models"
)

func Test_renderResult(t *testing.T) {
	type args struct {
		criterion  models.Criterion
		resultFlag bool
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "return true when truthy value is not set and result flag is true",
			args: args{
				criterion:  models.Criterion{},
				resultFlag: true,
			},
			want: true,
		},
		{
			name: "return truthy when truthy value is set and result flag is true",
			args: args{
				criterion:  models.Criterion{TruthyValue: "PASS"},
				resultFlag: true,
			},
			want: "PASS",
		},

		{
			name: "return false when falsey value is not set and result flag is false",
			args: args{
				criterion:  models.Criterion{},
				resultFlag: false,
			},
			want: false,
		},
		{
			name: "return falsey value when falsey value is set and result flag is false",
			args: args{
				criterion:  models.Criterion{FalseyValue: "FAIL"},
				resultFlag: false,
			},
			want: "FAIL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderResult(tt.args.criterion, tt.args.resultFlag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
