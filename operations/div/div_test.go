package operations

import (
	"reflect"
	"testing"
)

func Test_div(t *testing.T) {
	type args struct {
		a []int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{5, 9}, 8}, []int{1, 1}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := div(tt.args.a, tt.args.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("div() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("div() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
