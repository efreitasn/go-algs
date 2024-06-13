package logic

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHalfAdder(t *testing.T) {
	tests := []struct {
		a     uint8
		b     uint8
		carry uint8
		sum   uint8
	}{
		{0, 0, 0, 0},
		{0, 1, 0, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("a=%b_b=%b", test.a, test.b), func(t *testing.T) {
			carry, sum := HalfAdder(test.a, test.b)

			if diff := cmp.Diff(test.sum, sum); diff != "" {
				t.Errorf("invalid sum:\n%v", diff)
			}

			if diff := cmp.Diff(test.carry, carry); diff != "" {
				t.Fatalf("invalid carry:\n%v", diff)
			}
		})
	}
}

func TestFullAdder(t *testing.T) {
	tests := []struct {
		a        uint8
		b        uint8
		carryIn  uint8
		carryOut uint8
		sum      uint8
	}{
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
		{1, 1, 1, 1, 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("a=%b_b=%b", test.a, test.b), func(t *testing.T) {
			carryOut, sum := FullAdder(test.a, test.b, test.carryIn)

			if diff := cmp.Diff(test.sum, sum); diff != "" {
				t.Errorf("invalid sum:\n%v", diff)
			}

			if diff := cmp.Diff(test.carryOut, carryOut); diff != "" {
				t.Fatalf("invalid carry:\n%v", diff)
			}
		})
	}
}

func TestFullSubtractor(t *testing.T) {
	tests := []struct {
		a          uint8
		b          uint8
		borrowIn   uint8
		borrowOut  uint8
		difference uint8
	}{
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 0, 0, 0},
		{0, 1, 0, 1, 1},

		{0, 0, 1, 1, 1},
		{1, 0, 1, 0, 0},
		{1, 1, 1, 1, 1},
		{0, 1, 1, 1, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("a=%b_b=%b", test.a, test.b), func(t *testing.T) {
			borrowOut, difference := FullSubtractor(test.a, test.b, test.borrowIn)

			if diff := cmp.Diff(test.difference, difference); diff != "" {
				t.Errorf("invalid difference:\n%v", diff)
			}

			if diff := cmp.Diff(test.borrowOut, borrowOut); diff != "" {
				t.Fatalf("invalid borrow:\n%v", diff)
			}
		})
	}
}
