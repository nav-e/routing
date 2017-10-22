package osm

// Node represents one OSM node, that is defined by it's unique ID and location
type Node struct {
	ID   int64             `json:"osm_id"`
	Lat  float64           `json:"lat"`
	Lon  float64           `json:"lon"`
	Tags map[string]string `json:"tags"`
}

// Way represents an OSM way, defined by an unique ID and a list of Node ids
type Way struct {
	ID      int64
	Tags    map[string]string
	NodeIDs []int64
}

// Source has to call `WriteNode` and `WriteWay` on a GraphStore.
// data can come from .pbf files, .xml files, the web etc
type Source interface {
	WriteTo(g GraphStore)
}

// GraphStore is used to persist OSM data in a way that the graph structure is
// kept intact. Besides getting nodes by ID and searching for specific tags
// like street names, the GraphStore must provide information about the
// neighbors of certain nodes
type GraphStore interface {
	Add(n *Node)
	Get(id int64) *Node

	Connect(w *Way)
	Next(n *Node) []*Node
}

// SearchResult contains a display name and an osm reference id
type SearchResult struct {
	Name string `json:"display_name"`
	ID   int64  `json:"osm_id"`
}

// SearchStore has the capability to search for tag/value pairs
type SearchStore interface {
	GraphStore
	Find(tag string, value string) []SearchResult
}
