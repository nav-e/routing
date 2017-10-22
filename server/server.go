package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/bmizerany/pat"
	"github.com/nav-e/routing/algorithm"
	"github.com/nav-e/routing/osm"
)

type serverConfig struct {
	db     osm.SearchStore
	router algorithm.Dijkstra
	port   int
}

var config = &serverConfig{port: 8080}

func routingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	algorithm := r.URL.Query().Get(":algorithm")
	_ = algorithm // algorithm is not used for now
	from, err := strconv.Atoi(r.URL.Query().Get(":from"))
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	to, err := strconv.Atoi(r.URL.Query().Get(":to"))
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	start := config.db.Get(int64(from))
	goal := config.db.Get(int64(to))

	path, err := config.router.ShortestPath(start, goal)
	if err != nil {
		fmt.Fprintf(w, "Error: could not find path")
	} else {
		json.NewEncoder(w).Encode(path)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	name := r.URL.Query().Get(":name")
	results := config.db.Find("name", name)
	json.NewEncoder(w).Encode(results)
}

// Start a routing server
func Start(r io.Reader) {
	log.Println("Starting nav-e server")

	// Create db, parse osm data
	pbf := osm.PBFSource{Reader: r}
	config.db = osm.NewCache()
	pbf.WriteTo(config.db)
	log.Println("Converting osm data to graph")
	config.router = algorithm.Dijkstra{Graph: config.db, Metric: &algorithm.Meter{}}

	m := pat.New()
	m.Get("/:algorithm/from/:from/to/:to", http.HandlerFunc(routingHandler))
	m.Get("/search/:name", http.HandlerFunc(searchHandler))

	http.Handle("/", m)
	log.Printf("Listening on :%d", config.port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
