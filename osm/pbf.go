package osm

import (
	"io"
	"log"
	"runtime"

	"github.com/qedus/osmpbf"
)

// PBFSource provides an Source that reads from a pbf resource, like a pbf file
type PBFSource struct {
	Reader io.Reader
}

func reverseWay(nodeIDs []int64) []int64 {
	reversed := make([]int64, len(nodeIDs))
	for index := 0; index < len(nodeIDs); index++ {
		reversed[len(nodeIDs)-1-index] = nodeIDs[index]
	}
	return reversed
}

// WriteTo is necessary to implement an Source
func (p PBFSource) WriteTo(g GraphStore) {
	d := osmpbf.NewDecoder(p.Reader)
	d.SetBufferSize(osmpbf.MaxBlobSize)

	err := d.Start(runtime.GOMAXPROCS(-1))
	if err != nil {
		log.Fatal(err)
	}

	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			case *osmpbf.Node:
				g.Add(&Node{ID: v.ID, Lat: v.Lat, Lon: v.Lon, Tags: v.Tags})
			case *osmpbf.Way:
				g.Connect(&Way{ID: v.ID, NodeIDs: v.NodeIDs, Tags: v.Tags})
				// Some ways can be used both ways
				// TODO check for one way streets etc
				g.Connect(&Way{ID: v.ID, NodeIDs: reverseWay(v.NodeIDs), Tags: v.Tags})
			default:
				// Don't do anything
			}
		}
	}
}
