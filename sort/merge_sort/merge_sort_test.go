package sort

import (
	"fmt"
	"testing"
)

func Test_merge_sort(t *testing.T) {
	type args struct {
		a []int
		l int
		r int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{5, 3, 7, 1, 2}, 0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge_sort(tt.args.a, tt.args.l, tt.args.r)
			fmt.Print(tt.args.a)
		})
	}
}
