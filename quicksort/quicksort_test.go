package quicksort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{20, 50, 30, 10, 33},
			[]int{10, 20, 30, 33, 50},
		},
		{
			[]int{1, 2, 3, 50, 60},
			[]int{1, 2, 3, 50, 60},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.in), func(t *testing.T) {
			Exec(test.in)

			if !reflect.DeepEqual(test.in, test.out) {
				t.Errorf("got %v, want %v", test.in, test.out)
			}
		})
	}
}

func BenchmarkExec(b *testing.B) {
	size := int(1e5)
	s := make([]int, size)

	for i := size - 1; i >= 0; i-- {
		s[i] = i
	}

	for i := 0; i < b.N; i++ {
		Exec(s)
	}
}
