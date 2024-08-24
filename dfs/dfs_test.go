package dfs

import (
	"reflect"
	"testing"

	"github.com/efreitasn/go-datas/binarysearchtree"
)

func TestBinarySearchTreeNLR(t *testing.T) {
	bst := binarysearchtree.New[int]()

	bst.Insert(100)
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(95)
	bst.Insert(120)
	bst.Insert(110)
	bst.Insert(130)

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

	BinarySearchTreeNLR(bst, func(v int) bool {
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

	BinarySearchTreeNLR(bst, func(v int) bool {
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
	bst := binarysearchtree.New[int]()

	bst.Insert(100)
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(95)
	bst.Insert(120)
	bst.Insert(110)
	bst.Insert(130)

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

	BinarySearchTreeRNL(bst, func(v int) bool {
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

	BinarySearchTreeRNL(bst, func(v int) bool {
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
	bst := binarysearchtree.New[int]()

	bst.Insert(100)
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(95)
	bst.Insert(120)
	bst.Insert(110)
	bst.Insert(130)

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

	BinarySearchTreeLNR(bst, func(v int) bool {
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

	BinarySearchTreeLNR(bst, func(v int) bool {
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
	bst := binarysearchtree.New[int]()

	bst.Insert(100)
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(95)
	bst.Insert(120)
	bst.Insert(110)
	bst.Insert(130)

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

	BinarySearchTreeLRN(bst, func(v int) bool {
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

	BinarySearchTreeLRN(bst, func(v int) bool {
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
