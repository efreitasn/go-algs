// Package bfs provides functions to perform a traverse using the breadth-first search algorithm.
package bfs

import (
	"github.com/efreitasn/go-datas/graph"
	"github.com/efreitasn/go-datas/queue"
)

// Graph performs a search in a graph using breadth-first-search.
func Graph[T comparable](g *graph.Graph[T], startVertex T, cb func(v T) bool) {
	if !g.HasVertex(startVertex) {
		return
	}

	q := queue.New[T]()
	visited := map[T]bool{}

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

		adjVertices.Traverse(true, func(adjV T) bool {
			if !visited[adjV] {
				visited[adjV] = true

				q.Enqueue(adjV)
			}

			return true
		})
	}
}
