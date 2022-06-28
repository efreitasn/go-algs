// Package dfs provides functions to perform a traverse using the depth-first search algorithm.
package dfs

import (
	"github.com/efreitasn/go-datas/binarysearchtree"
	"golang.org/x/exp/constraints"
)

// BinarySearchTreeNLR performs a traverse using the depth-first search algorithm with a pre-order strategy in a binary search tree of an ordered type T.
func BinarySearchTreeNLR[T constraints.Ordered](bts *binarysearchtree.BinarySearchTree[T], cb func(v T) bool) {
	if bts.Size() == 0 {
		return
	}

	binarySearchTreeNLRRecursive(bts.Root(), cb)
}

func binarySearchTreeNLRRecursive[T constraints.Ordered](n *binarysearchtree.Node[T], cb func(v T) bool) bool {
	if n == nil {
		return true
	}

	valR := cb(n.Value())

	if !valR {
		return false
	}

	if n.Left() != nil {
		leftR := binarySearchTreeNLRRecursive(n.Left(), cb)

		if !leftR {
			return false
		}
	}

	if n.Right() != nil {
		rightR := binarySearchTreeNLRRecursive(n.Right(), cb)

		if !rightR {
			return false
		}
	}

	return true
}

// BinarySearchTreeLNR performs a traverse using the depth-first search algorithm with an in-order strategy in a binary search tree of an ordered type T.
func BinarySearchTreeLNR[T constraints.Ordered](bts *binarysearchtree.BinarySearchTree[T], cb func(v T) bool) {
	if bts.Size() == 0 {
		return
	}

	binarySearchTreeLNRRecursive(bts.Root(), cb)
}

func binarySearchTreeLNRRecursive[T constraints.Ordered](n *binarysearchtree.Node[T], cb func(v T) bool) bool {
	if n == nil {
		return true
	}

	if n.Left() != nil {
		leftR := binarySearchTreeLNRRecursive(n.Left(), cb)

		if !leftR {
			return false
		}
	}

	valR := cb(n.Value())

	if !valR {
		return false
	}

	if n.Right() != nil {
		rightR := binarySearchTreeLNRRecursive(n.Right(), cb)

		if !rightR {
			return false
		}
	}

	return true
}

// BinarySearchTreeRNL performs a traverse using the depth-first search algorithm with an out-order strategy in a binary search tree of an ordered type T.
func BinarySearchTreeRNL[T constraints.Ordered](bts *binarysearchtree.BinarySearchTree[T], cb func(v T) bool) {
	if bts.Size() == 0 {
		return
	}

	binarySearchTreeRNLRecursive(bts.Root(), cb)
}

func binarySearchTreeRNLRecursive[T constraints.Ordered](n *binarysearchtree.Node[T], cb func(v T) bool) bool {
	if n == nil {
		return true
	}

	if n.Right() != nil {
		rightR := binarySearchTreeRNLRecursive(n.Right(), cb)

		if !rightR {
			return false
		}
	}

	valR := cb(n.Value())

	if !valR {
		return false
	}

	if n.Left() != nil {
		leftR := binarySearchTreeRNLRecursive(n.Left(), cb)

		if !leftR {
			return false
		}
	}

	return true
}

// BinarySearchTreeLRN performs a traverse using the depth-first search algorithm with a post-order strategy in a binary search tree of an ordered type T.
func BinarySearchTreeLRN[T constraints.Ordered](bts *binarysearchtree.BinarySearchTree[T], cb func(v T) bool) {
	if bts.Size() == 0 {
		return
	}

	binarySearchTreeLRNRecursive(bts.Root(), cb)
}

func binarySearchTreeLRNRecursive[T constraints.Ordered](n *binarysearchtree.Node[T], cb func(v T) bool) bool {
	if n == nil {
		return true
	}

	if n.Left() != nil {
		leftR := binarySearchTreeLRNRecursive(n.Left(), cb)

		if !leftR {
			return false
		}
	}

	if n.Right() != nil {
		rightR := binarySearchTreeLRNRecursive(n.Right(), cb)

		if !rightR {
			return false
		}
	}

	valR := cb(n.Value())

	if !valR {
		return false
	}

	return true
}
