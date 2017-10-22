package routing

import (
	"fmt"
	"math"

	"github.com/nav-e/routing/osm"
)

// Dijkstra implements the Dijkstra shortest path algorithm
type Dijkstra struct {
	Graph  osm.GraphStore
	Metric Metric
}

// pop removes the next node (measured by cost) and returns it
func pop(frontier *NodeSet, cost map[*osm.Node]float64) *osm.Node {
	var res *osm.Node
	min := math.MaxFloat64
	for n := range frontier.Nodes {
		if cost[n] < min {
			min = cost[n]
			res = n
		}
	}
	frontier.Delete(res)
	return res
}

// backtrack takes a map of parents, a start and a goal and returns the path
// from start to goal
func backtrack(start, goal *osm.Node, parent map[*osm.Node]*osm.Node) []*osm.Node {
	res := make([]*osm.Node, 0)
	res = append(res, goal)
	curr := goal
	for curr != start {
		res = append([]*osm.Node{parent[curr]}, res...) // Prepend to res
		curr = parent[curr]
	}
	return res
}

// ShortestPath using the Dijkstra algorithm
func (a *Dijkstra) ShortestPath(start, goal *osm.Node) ([]*osm.Node, error) {
	frontier := NewNodeSet()
	dist := make(map[*osm.Node]float64)
	prev := make(map[*osm.Node]*osm.Node)

	frontier.Add(start)
	dist[start] = 0.0

	for len(frontier.Nodes) > 0 {
		u := pop(frontier, dist)
		if u == goal {
			return backtrack(start, goal, prev), nil
		}
		for _, v := range a.Graph.Next(u) {
			alt := dist[u] + a.Metric.Cost(u, v)
			vDist, ok := dist[v]
			if !ok {
				frontier.Add(v)
				dist[v] = alt
				prev[v] = u
			} else if vDist > alt {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	return nil, fmt.Errorf("No path could be found after %d nodes", len(dist))
}
