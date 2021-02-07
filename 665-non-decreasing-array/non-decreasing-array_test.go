package _65_non_decreasing_array

import "testing"

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"5, 7, 1, 8, 7, 6", args{[]int{5, 7, 1, 8, 7, 6}}, false},
		{"5, 7, 1, 8", args{[]int{5, 7, 1, 8}}, true},
		{"zero", args{[]int{}}, true},
		{"4, 2, 3", args{[]int{4, 2, 3}}, true},
		{"1, 2, 3, 4, 5, 6, 7", args{[]int{1, 2, 3, 4, 5, 6, 7}}, true},
		{"1, 2", args{[]int{1, 2}}, true},
		{"3, 4, 2, 3", args{[]int{3, 4, 2, 3}}, false},
		{"nil", args{nil}, true},
		{"1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 5", args{[]int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
