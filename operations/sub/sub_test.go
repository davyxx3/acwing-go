package operations

import (
	"reflect"
	"testing"
)

func Test_sub(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{8, 9}, []int{9, 6}}, []int{9, 2}},
		{"test2", args{[]int{9, 9, 9, 9}, []int{2, 8, 9, 9}}, []int{7, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sub(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
