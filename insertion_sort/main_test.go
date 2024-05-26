package main

import (
	"reflect"
	"testing"
)

func TestInsertionSortBinarySearch(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case 0",
			args: args{
				nums: []int{2, 3, 1},
			},
			want: []int{1, 2, 3},
		},
		/* {
			name: "Test case 1",
			args: args{
				nums: []int{5, 2, 4, 6, 1, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		}, */
		/* 	{
			name: "Test case 2",
			args: args{
				nums: []int{31, 41, 59, 26, 41, 58},
			},
			want: []int{26, 31, 41, 41, 58, 59},
		}, */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSortBinarySearch(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSortBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
