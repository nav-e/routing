package osm_test

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/nav-e/routing/osm"
)

// Implement a GraphStore for testing
type countingStore struct {
	osm.GraphStore
	nodes, ways int
}

func (c *countingStore) Add(node *osm.Node) {
	c.nodes++
}
func (c *countingStore) Get(id int64) *osm.Node {
	return nil
}
func (c *countingStore) Connect(way *osm.Way) {
	c.ways++
}
func (c *countingStore) Previous(node *osm.Node) []*osm.Node {
	return nil
}
func (c *countingStore) Next(node *osm.Node) []*osm.Node {
	return nil
}

// Count the nodes in the monaco.pbf file
const monacoFile = "../resources/monaco-latest.osm.pbf"
const monacoNodes = 18647
const monacoWays = 2620 * 2

func TestWriteTo(t *testing.T) {
	file, err := os.Open(monacoFile)
	if err != nil {
		log.Fatal(err)
	}
	s := osm.PBFSource{Reader: bufio.NewReader(file)}
	c := countingStore{}
	s.WriteTo(&c)

	if c.nodes != monacoNodes {
		t.Errorf("Expected %d nodes, got %d", monacoNodes, c.nodes)
	}

	if c.ways != monacoWays {
		t.Errorf("Expected %d ways, got %d", monacoWays, c.ways)
	}
}
