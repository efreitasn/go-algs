// Package bfs provides functions to perform a traverse using the breadth-first search algorithm.
package bfs

import (
	"github.com/efreitasn/go-datas/graph"
	"github.com/efreitasn/go-datas/queue"
)

// Graph performs a search in a graph using breadth-first-search.
func Graph(g *graph.Graph, startVertex int, cb func(v int) bool) {
	if !g.HasVertex(startVertex) {
		return
	}

	q := queue.New()
	visited := map[int]bool{}

	q.Enqueue(startVertex)
	visited[startVertex] = true

	for q.Size() != 0 {
		v, _ := q.Peek()

		q.Dequeue()

		r := cb(v)

		if !r {
			break
		}

		adjVertices, _ := g.AdjacentVertices(v)

		adjVertices.Traverse(true, func(adjV int) {
			if !visited[adjV] {
				visited[adjV] = true

				q.Enqueue(adjV)
			}
		})
	}
}
