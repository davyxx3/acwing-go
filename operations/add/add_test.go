package operations

import (
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
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
		{"test1", args{[]int{2, 1}, []int{3, 2}}, []int{5, 3}},
		{"test2", args{[]int{7, 6, 8, 5}, []int{4, 2, 8, 7, 1}}, []int{1, 9, 6, 3, 2}},
		{"test3", args{[]int{7, 9, 2, 8}, []int{3, 1, 2, 6}}, []int{0, 1, 5, 4, 1}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
