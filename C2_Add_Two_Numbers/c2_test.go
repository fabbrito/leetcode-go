package c2

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct{ l1, l2 []int }
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Same number of digits",
			args: args{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "Same number of digits and carry",
			args: args{
				l1: []int{2, 4, 5},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 0, 1},
		},
		{
			name: "Different number of digits",
			args: args{
				l1: []int{1},
				l2: []int{9, 9, 8},
			},
			want: []int{0, 0, 9},
		},
		{
			name: "Different number of digits and carry",
			args: args{
				l1: []int{9, 9, 9, 9, 9, 9, 9},
				l2: []int{9, 9, 9, 9},
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := arrayToList(tt.args.l1)
			l2 := arrayToList(tt.args.l2)
			if got := listToArray(addTwoNumbers(l1, l2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ListNode(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "Empty array",
			nums: []int{},
		},
		{
			name: "One element",
			nums: []int{1},
		},
		{
			name: "Multiple elements",
			nums: []int{3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := arrayToList(tt.nums)
			if got := listToArray(head); !reflect.DeepEqual(got, tt.nums) {
				t.Errorf("Error to convert from array to ListNode or vice versa")
			}
		})
	}
}
