package main

import (
	"log"
	"os"

	"github.com/nav-e/routing/server"
)

func main() {
	// file, err := os.Open("./resources/monaco-latest.osm.pbf")
	file, err := os.Open("./SanFrancisco.osm.pbf")
	if err != nil {
		log.Fatal("Monaco pbf could not be found")
	}
	server.Start(file)
}
