package binarysearch

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		s     []int
		value int
		i     int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			5,
			4,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			9,
			-1,
		},
		{
			[]int{1, 2, 3},
			10,
			-1,
		},
		{
			[]int{4, 5, 6, 7, 8, 9, 10, 30, 300, 400},
			300,
			8,
		},
		{
			[]int{4, 5, 6, 7, 8, 9, 10, 30, 300, 400},
			10,
			6,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v - %v", test.s, test.value), func(t *testing.T) {
			i := Exec(test.s, test.value)

			if i != test.i {
				t.Errorf("got %v, want %v", i, test.i)
			}
		})
	}
}
