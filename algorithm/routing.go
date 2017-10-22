package algorithm

import (
	"math"

	"github.com/nav-e/routing/osm"
)

// Metric determines the Cost of going from one Node to the next
type Metric interface {
	Cost(from, to *osm.Node) float64
}

// Pathfinder structs are general routing algorithms
type Pathfinder interface {
	ShortestPath(start, goal *osm.Node) []*osm.Node
}

// Meter computest the Cost between Nodes in meters
type Meter struct {
	Metric
}

const r = 6378100 // Earth radius in m

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Cost between Nodes in meters
func (m *Meter) Cost(from, to *osm.Node) float64 {
	var lat1, lon1, lat2, lon2 = from.Lat, from.Lon, to.Lat, to.Lon
	la1 := lat1 * math.Pi / 180
	lo1 := lon1 * math.Pi / 180
	la2 := lat2 * math.Pi / 180
	lo2 := lon2 * math.Pi / 180

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)
	return 2 * r * math.Asin(math.Sqrt(h))
}
