package dfs

import (
	"reflect"
	"testing"

	"github.com/efreitasn/go-datas/graph"
)

func TestGraphPreOrder(t *testing.T) {
	g := graph.New(false)
	vals := []int{}

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)
	g.AddVertex(9)
	g.AddVertex(10)
	g.AddVertex(11)
	g.AddVertex(12)

	g.AddEdge(1, 2)
	g.AddEdge(1, 7)
	g.AddEdge(1, 8)

	g.AddEdge(2, 3)
	g.AddEdge(2, 6)

	g.AddEdge(3, 4)
	g.AddEdge(3, 5)

	g.AddEdge(8, 9)
	g.AddEdge(8, 12)

	g.AddEdge(9, 10)
	g.AddEdge(9, 11)

	GraphPreOrder(g, 1, func(v int) bool {
		vals = append(vals, v)

		return true
	})

	expectedVals := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
	}

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

	vals = []int{}

	GraphPreOrder(g, 1, func(v int) bool {
		if v == 6 {
			return false
		}

		vals = append(vals, v)

		return true
	})

	expectedVals = []int{
		1,
		2,
		3,
		4,
		5,
	}

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}
}
