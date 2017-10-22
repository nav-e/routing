package algorithm

import (
	"github.com/nav-e/routing/osm"
)

// NodeSet is a set containing Nodes
type NodeSet struct {
	Nodes map[*osm.Node]bool
}

// NewNodeSet initializes the set
func NewNodeSet() *NodeSet {
	s := NodeSet{}
	s.Nodes = make(map[*osm.Node]bool)

	return &s
}

// Contains returns true on Nodes in the set
func (s *NodeSet) Contains(n *osm.Node) bool {
	_, ok := s.Nodes[n]
	return ok
}

// Add adds a Node to the set, without duplicates
func (s *NodeSet) Add(n *osm.Node) {
	s.Nodes[n] = true
}

// Delete removes a Node from the set
func (s *NodeSet) Delete(n *osm.Node) {
	delete(s.Nodes, n)
}
