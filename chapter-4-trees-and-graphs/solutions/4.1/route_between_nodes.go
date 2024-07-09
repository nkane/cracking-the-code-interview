package route_between_nodes

import (
	"graph"

	"github.com/golang-collections/collections/queue"
)

/*
	4.1: Route Between Nodes: Given a directed graph and two nodes (S and E), design an algorithm to
	find out whether there is a route from S to E.
*/

/*
This problem can be solved by just simple graph traversal, such as depth-first search or breadth-first
search. We start with one of the two nodes and, during traversal, check if the other node is found. We
should mark any node found in the course of the algorithm as "already visited" to avoid cycles and
repitition of the nodes.
*/

const (
	Unvisited = iota
	Visited
	Visiting
)

func Search(g *graph.Graph[int], start *graph.Node[int], end *graph.Node[int]) bool {
	if start == end {
		return true
	}
	// operates as a queue
	q := queue.New()
	for i := 0; i < len(g.Nodes); i++ {
		g.Nodes[i].State = Unvisited
	}
	start.State = Visiting
	q.Enqueue(start)
	var u *graph.Node[int]
	for q.Len() > 0 {
		u = q.Dequeue().(*graph.Node[int])
		if u != nil {
			for _, v := range u.Children {
				if v.State == Unvisited {
					if v == end {
						return true
					} else {
						v.State = Visiting
						q.Enqueue(v)
					}
				}
			}
		}
		u.State = Visited
	}
	return false
}
