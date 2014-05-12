# grapho [![Build Status](https://travis-ci.org/ichinaski/grapho.svg?branch=master)](https://travis-ci.org/ichinaski/grapho) [![GoDoc](http://godoc.org/github.com/ichinaski/grapho?status.png)](http://godoc.org/github.com/ichinaski/grapho)


grapho is a Go implementation of Graph Theory data structures and algorithms. The full documentation can be found [here](http://godoc.org/github.com/ichinaski/grapho)

## Graphs

Graphs, both directed and undirected, are created with the `Graph` struct:

```
graph := grapho.NewGraph(false) // true for directed Graphs, false for undirected
```

`Nodes` (vertices) and `Edges` contain a set of attributes `Attr`, whose key is a `string`, and the value can be anything (type `interface{}`). Attr is in essence a `map[string]interface{}`, and its values are set/get with the regular map syntax:

```
attr := grapho.NewAttr()
attr["name"] = "Bob"
attr["x"] = 1
```

Nodes, uniquely identified by a `uint64` id, can be added explicitly to the Graph:

```
graph.AddNode(1, attr)
```

or implicitly, if adding an Edge references a non-existing Node:

```
graph.AddEdge(1, 2, nil) // Node '2' will be automatically created
```

To check all the available methods to manipulate a Graph, see the [GoDoc](http://godoc.org/github.com/ichinaski/grapho#Graph).

## Algorithms

### Search:
* Dijkstra (uniform cost search)
* A* (best-first search with heuristic)
* Breath-first search
* Depth-first search

For the shortest path implementation, a generic `Search` function is provided, which takes the desired `SearchAlgorithm` to be used as a parameter:

```
path, err := grapho.Search(graph, 1, 8, grapho.Dijkstra, nil)
path, err := grapho.Search(graph, 1, 9, grapho.BreathFirstSearch, nil)
...

```
If successful, a `uint64` slice will be returned, with the node ids that form the shortest path between the specified nodes. Check out the second return value, for any possible error (i.e. no path found).

### Minimum Spanning Tree:
* Prim
* TODO: Kruskal

Given a connected, undirected graph, `MinimumSpanningTree` calculates the minimum cost subgraph that connects all the vertices together:

```
mst, err := grapho.MinimumSpanningTree(graph, grapho.Prim)
```

### TODO:
* Multigraph
* Minimum Cut
* Topological short
* Coloring algorithms
* Strongly Connected Components
* Eulerian path/circuit
* Flow Networks
* ...

## Contributing

I have no priority in mind for future algorithms to implement, so if you want to contribute with any Graph Theory algorithm (whether or not in the list above), don't hesitate in opening an issue or sending a pull request!

## License

This library is distributed under the BSD-style license found in the LICENSE file.
