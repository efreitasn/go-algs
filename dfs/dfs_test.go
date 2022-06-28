package dfs

import (
	"reflect"
	"testing"

	"github.com/efreitasn/go-datas/binarysearchtree"
)

func TestBinarySearchTreeNLR(t *testing.T) {
	bts := binarysearchtree.New[int]()

	bts.Insert(100)
	bts.Insert(90)
	bts.Insert(80)
	bts.Insert(95)
	bts.Insert(120)
	bts.Insert(110)
	bts.Insert(130)

	vals := []int{}
	expectedVals := []int{
		100,
		90,
		80,
		95,
		120,
		110,
		130,
	}

	BinarySearchTreeNLR(bts, func(v int) bool {
		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

	vals = []int{}
	expectedVals = []int{
		100,
		90,
		80,
		95,
	}

	BinarySearchTreeNLR(bts, func(v int) bool {
		if v == 120 {
			return false
		}

		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

}

func TestBinarySearchTreeRNL(t *testing.T) {
	bts := binarysearchtree.New[int]()

	bts.Insert(100)
	bts.Insert(90)
	bts.Insert(80)
	bts.Insert(95)
	bts.Insert(120)
	bts.Insert(110)
	bts.Insert(130)

	vals := []int{}
	expectedVals := []int{
		130,
		120,
		110,
		100,
		95,
		90,
		80,
	}

	BinarySearchTreeRNL(bts, func(v int) bool {
		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

	vals = []int{}
	expectedVals = []int{
		130,
		120,
		110,
	}

	BinarySearchTreeRNL(bts, func(v int) bool {
		if v == 100 {
			return false
		}

		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

}

func TestBinarySearchTreeLNR(t *testing.T) {
	bts := binarysearchtree.New[int]()

	bts.Insert(100)
	bts.Insert(90)
	bts.Insert(80)
	bts.Insert(95)
	bts.Insert(120)
	bts.Insert(110)
	bts.Insert(130)

	vals := []int{}
	expectedVals := []int{
		80,
		90,
		95,
		100,
		110,
		120,
		130,
	}

	BinarySearchTreeLNR(bts, func(v int) bool {
		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

	vals = []int{}
	expectedVals = []int{
		80,
		90,
		95,
	}

	BinarySearchTreeLNR(bts, func(v int) bool {
		if v == 100 {
			return false
		}

		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

}

func TestBinarySearchTreeLRN(t *testing.T) {
	bts := binarysearchtree.New[int]()

	bts.Insert(100)
	bts.Insert(90)
	bts.Insert(80)
	bts.Insert(95)
	bts.Insert(120)
	bts.Insert(110)
	bts.Insert(130)

	vals := []int{}
	expectedVals := []int{
		80,
		95,
		90,
		110,
		130,
		120,
		100,
	}

	BinarySearchTreeLRN(bts, func(v int) bool {
		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

	vals = []int{}
	expectedVals = []int{
		80,
		95,
		90,
	}

	BinarySearchTreeLRN(bts, func(v int) bool {
		if v == 110 {
			return false
		}

		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}
}
