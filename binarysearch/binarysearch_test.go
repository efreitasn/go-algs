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
		found bool
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			5,
			4,
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			9,
			0,
			false,
		},
		{
			[]int{4, 5, 6, 7, 8, 9, 10, 30, 300, 400},
			300,
			8,
			true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v - %v", test.s, test.value), func(t *testing.T) {
			i, found := Exec(test.s, test.value)

			if i != test.i || found != test.found {
				t.Errorf("got %v and %v, want %v and %v", i, found, test.i, test.found)
			}
		})
	}
}
