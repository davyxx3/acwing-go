package operations

import (
	"reflect"
	"testing"
)

func Test_mul(t *testing.T) {
	type args struct {
		a []int
		b int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{6, 5, 2, 3}, 69}, []int{4, 6, 6, 4, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mul(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
