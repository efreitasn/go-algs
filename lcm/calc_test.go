package lcm

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		a   int
		b   int
		out int
	}{
		{30, 20, 60},
		{14, 7, 14},
		{64, 48, 192},
		{113, 37, 4181},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v, %v", test.a, test.b), func(t *testing.T) {
			res := Calc(test.a, test.b)

			if res != test.out {
				t.Errorf("got %v, want %v", res, test.out)
			}
		})
	}
}
