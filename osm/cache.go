package osm

import (
	"log"
	"strings"
)

// MaxResults is the maximum count for string searches
const MaxResults = 5

// Cache holds OSM data in memory. Can be used for routing in small cities
type Cache struct {
	SearchStore
	Nodes map[int64]*Node
	Edges map[int64][]int64
}

// NewCache initializes a Cache with Node and Edge maps
func NewCache() *Cache {
	c := Cache{}
	c.Nodes = make(map[int64]*Node)
	c.Edges = make(map[int64][]int64)
	return &c
}

// Add a Node to the Cache
func (c *Cache) Add(n *Node) {
	c.Nodes[n.ID] = n
	if _, ok := c.Edges[n.ID]; !ok {
		c.Edges[n.ID] = make([]int64, 0)
	}
}

// Get a Node from the Cache by id
func (c *Cache) Get(id int64) *Node {
	return c.Nodes[id]
}

func (c *Cache) writeEdge(fromID, toID int64) {
	if next, ok := c.Edges[fromID]; ok {
		c.Edges[fromID] = append(next, toID)
	} else {
		c.Edges[fromID] = []int64{toID}
	}

}

// Connect nodes in the Cache from OSM way data
func (c *Cache) Connect(w *Way) {
	for i := 0; i < len(w.NodeIDs)-1; i++ {
		c.writeEdge(w.NodeIDs[i], w.NodeIDs[i+1])
	}
}

// Next neighbors of a specific Node in the Cache
func (c *Cache) Next(n *Node) []*Node {
	nextIDs, ok := c.Edges[n.ID]
	if !ok {
		log.Fatalf("Node-ID was not stored but requested: %d", n.ID)
	}
	results := make([]*Node, len(nextIDs))
	for i, n := range c.Edges[n.ID] {
		results[i] = c.Nodes[n]
	}
	return results
}

// Find nodes by tag values
func (c *Cache) Find(tag string, value string) []SearchResult {
	cnt := 0
	res := make([]SearchResult, 0)
	for _, n := range c.Nodes {
		hasNeighbors := len(c.Edges[n.ID]) > 0
		name, hasName := n.Tags[tag]
		if hasNeighbors && hasName {
			if strings.Contains(name, value) {
				res = append(res, SearchResult{ID: n.ID, Name: name})
				if cnt++; cnt > MaxResults {
					break
				}
			}
		}
	}
	return res
}
