package osm_test

import (
	"testing"

	"github.com/nav-e/routing/osm"
)

// Test data and mocking
var n0 = &osm.Node{ID: 0, Tags: map[string]string{"tag": "value"}}
var n1 = &osm.Node{ID: 1, Tags: map[string]string{"tag": "value2"}}
var n2 = &osm.Node{ID: 2}
var w0 = &osm.Way{NodeIDs: []int64{0, 1, 2}}
var w1 = &osm.Way{NodeIDs: []int64{0, 2}}

type MockSource struct {
	osm.Source
}

func (m *MockSource) WriteTo(g osm.GraphStore) {
	g.Add(n0)
	g.Add(n1)
	g.Connect(w0)
	g.Connect(w1)
	g.Add(n2) // Add a node AFTER it has been connected
}

// Cache testing
func TestCache(t *testing.T) {
	m := MockSource{}
	c := osm.NewCache()
	m.WriteTo(c)

	if c.Get(0) != n0 || c.Get(1) != n1 || c.Get(2) != n2 {
		t.Errorf("Could not retrieve saved nodes by id")
	}
	if c.Next(n1)[0] != n2 {
		t.Errorf("Could not get next for node with one neighbor")
	}
	if len(c.Next(n0)) != 2 {
		t.Errorf("Could not get next for node with multiple neighbors")
	}
	if len(c.Next(n2)) != 0 {
		t.Errorf("Could not get next for node without neighbors")
	}
}

func TestCacheSearch(t *testing.T) {
	m := MockSource{}
	c := osm.NewCache()
	m.WriteTo(c)

	if res := c.Find("doesnotexist", "xyz"); len(res) > 0 {
		t.Errorf("Found non-existing nodes, got %v", res)
	}
	if res := c.Find("tag", "value"); len(res) != 2 {
		t.Errorf("Found not two expected results(query for tag:value), got %v", res)
	}
	if res := c.Find("tag", "2"); len(res) != 1 {
		t.Errorf("Found not one expected value(query for tag:2), got %v", res)
	}
}
