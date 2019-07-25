package bubblesort

import (
	"reflect"
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{
			"5, 4, 3, 2, 1",
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"20, 50, 30, 10, 33",
			[]int{20, 50, 30, 10, 33},
			[]int{10, 20, 30, 33, 50},
		},
		{
			"1, 2, 3, 50, 60",
			[]int{1, 2, 3, 50, 60},
			[]int{1, 2, 3, 50, 60},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Exec(test.in)

			if !reflect.DeepEqual(test.in, test.out) {
				t.Errorf("got %v, want %v", test.in, test.out)
			}
		})
	}
}
