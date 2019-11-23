package mod

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		a   int
		p   int
		out int
	}{
		{-1, 3, 2},
		{28, 5, 3},
		{50, 10, 0},
		{5, -10, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%vmod%v", test.a, test.p), func(t *testing.T) {
			res := Calc(test.a, test.p)

			if res != test.out {
				t.Errorf("got %v, want %v", res, test.out)
			}
		})
	}
}
