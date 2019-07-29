// Package dfs provides functions to perform a traverse using the depth-first search algorithm.
package dfs

import (
	"github.com/efreitasn/go-datas/graph"
	"github.com/efreitasn/go-datas/stack"
	"github.com/efreitasn/go-datas/tree"
)

// GraphPreOrder performs a traverse using the depth-first search algorithm with a pre-order strategy in a graph of ints.
func GraphPreOrder(g *graph.Graph, startVertex int, cb func(v int) bool) {
	if !g.HasVertex(startVertex) {
		return
	}

	s := stack.New(startVertex)
	visited := map[int]bool{
		startVertex: true,
	}

	for s.Size() != 0 {
		v, _ := s.Pop()

		r := cb(v)

		if !r {
			break
		}

		adjVertices, _ := g.AdjacentVertices(v)

		adjVertices.Traverse(false, func(adjV int) {
			if !visited[adjV] {
				s.Push(adjV)

				visited[adjV] = true
			}
		})
	}
}

// TreePreOrder performs a traverse using the depth-first search algorithm with a pre-order strategy in a tree of ints.
func TreePreOrder(tr *tree.Tree, startNode int, cb func(v int) bool) {
	if !tr.HasNode(startNode) {
		return
	}

	s := stack.New(startNode)

	for s.Size() != 0 {
		v, _ := s.Pop()

		r := cb(v)

		if !r {
			break
		}

		children, _ := tr.NodeChildren(v)

		children.Traverse(false, func(child int) {
			s.Push(child)
		})
	}
}
