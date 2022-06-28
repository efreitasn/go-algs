package bfs

import (
	"reflect"
	"testing"

	"github.com/efreitasn/go-datas/graph"
)

func TestGraph(t *testing.T) {
	// Directed graph
	g := graph.New[int](true)
	vals := []int{}

	g.AddVertex(10)
	g.AddVertex(20)
	g.AddVertex(50)
	g.AddVertex(90)

	g.AddEdge(10, 20)
	g.AddEdge(10, 50)
	g.AddEdge(20, 90)

	expectedVals := []int{
		10,
		20,
		50,
		90,
	}

	Graph(g, 10, func(v int) bool {
		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

	g = graph.New[int](false)
	vals = []int{}

	g.AddVertex(10)
	g.AddVertex(20)
	g.AddVertex(50)
	g.AddVertex(90)

	g.AddEdge(10, 20)
	g.AddEdge(10, 50)
	g.AddEdge(20, 90)

	Graph(g, 10, func(v int) bool {
		vals = append(vals, v)

		return true
	})

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

}
