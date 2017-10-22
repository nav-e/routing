# `routing` - The nav-e routing server project

Nav-e is an open-source routing project that is focused on energy optimality. 
This is the back end for [our routing frontend](https://github.com/nav-e/nav-e). 

## Routing server for nav-e, written in Go

To build a project that is able to run on smaller devices that can be used for offline navigation, 
the project is written in Go (instead of Java, for example). In order to run the application locally,
install [go](https://golang.org/), clone this repository and run

```zsh
routing> go run main.go
2017/10/22 21:51:08 Starting nav-e server
2017/10/22 21:51:08 Converting osm data to graph
2017/10/22 21:51:08 Listening on :8080
```

## API

As a backend, this projects provides a simple REST api to route

- `/search/:name`
  - `:name` must be string
  - GET: `http://localhost:8080/search/Hec`
  - RESPONSE: `[{"display_name":"Hector-Otto","osm_id":25200449}]`
- `/:algorithm/from/:from/to/:to`
  - `:from`, `:to` must be valid OSM ids, use `/search/:name` to find them
  - `:algorithm` is ignored for now, only dijkstra is implemented and used
  - GET: `http://localhost:8080/dijkstra/from/25200449/to/2185515256`
  - RESPONSE: `[{"osm_id":25200449,"lat":43.7335195,"lon":7.413941400000001,"tags":{"bus":"yes","name":"Hector-Otto","public_transport":"stop_position"}},{"osm_id":25198987,"lat":43.7336467,"lon":7.4141639,"tags":{}},{"osm_id":25198929,"lat":43.733877400000004,"lon":7.4145156000000005,"tags":{}},{"osm_id":25198922,"lat":43.7340929...]
`

## TODO

- [x] Basic routing
- [x] PBF import
- [x] REST API
- []  Use Database for routing
- []  Implement energy optimal shortest path algorithm
- []  Implement one energy consumption algorithm (see other nav-e projects)
- []  Port range polygon (see range-anxiety project)
