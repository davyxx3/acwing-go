package sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
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
		{"test1", args{[]int{3, 5, 6, 1, 2}, 0, 4}},
		{"test2", args{[]int{7, 2, 5, 1, 3, 0, 4}, 0, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.a, tt.args.l, tt.args.r)
			fmt.Print(tt.args.a)
		})
	}
}
