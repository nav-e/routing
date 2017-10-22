package routing_test

import (
	"log"
	"os"
	"testing"

	"github.com/nav-e/routing/osm"
	"github.com/nav-e/routing/routing"
)

// Parameters for test route
const pathLen = 65
const startID = 25200449
const goalID = 1770577832

func TestDijkstra(t *testing.T) {
	// Setup
	file, err := os.Open("../resources/monaco-latest.osm.pbf")
	if err != nil {
		log.Fatal("Monaco pbf could not be found")
	}
	pbf := osm.PBFSource{Reader: file}
	cache := osm.NewCache()
	pbf.WriteTo(cache)

	// Routing test
	dijkstra := routing.Dijkstra{Graph: cache, Metric: &routing.Meter{}}
	start := cache.Get(startID)
	goal := cache.Get(goalID)
	path, err := dijkstra.ShortestPath(start, goal)
	if err != nil {
		t.Error(err)
	}
	if len(path) != pathLen {
		t.Errorf("Expected example path length to be %d, got %d instead",
			pathLen, len(path))
	}
}
