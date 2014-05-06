package grapho

import (
	"errors"
	"github.com/ichinaski/grapho/container"
)

const (
	type_bfs = iota
	type_dfs
	type_dijkstra
	type_astar
)

// searchstate is a graph position for which its ancestors have been evaluated.
// Contains the nodeId, its parent int, and the total cost of traversing
// the graph to reach this position
type searchstate struct {
	node, parent uint64
	cost         int // OpenSet takes int as priority type. TODO: add uint64 support
}

// OpenSet defines the functions that any container used to keep track of
// non-expanded nodes must implement
type OpenSet interface {
	Push(item interface{}, priority int)
	Pop() interface{}
	Len() int
}

// Wrap Queue to match OpenSet interface signature
type OpenQueue struct{ *container.Queue }

func (q *OpenQueue) Push(item interface{}, priority int) { q.Queue.Push(item) }

// Wrap Stack to match OpenSet interface signature
type OpenStack struct{ *container.Stack }

func (s *OpenStack) Push(item interface{}, priority int) { s.Stack.Push(item) }

// Heuristic calculates the estimated cost between two nodes
type Heuristic func(node, goal uint64) int

func NullHeuristic(node, goal uint64) int { return 0 }

func BreathFirstSearch(graph *Graph, start, goal uint64) ([]uint64, error) {
	return search(graph, start, goal, type_bfs, nil)
}

func DepthFirstSearch(graph *Graph, start, goal uint64) ([]uint64, error) {
	return search(graph, start, goal, type_dfs, nil)
}

func Dijkstra(graph *Graph, start, goal uint64) ([]uint64, error) {
	return search(graph, start, goal, type_dijkstra, nil)
}

func Astar(graph *Graph, start, goal uint64, heuristic Heuristic) ([]uint64, error) {
	return search(graph, start, goal, type_astar, heuristic)
}

func search(graph *Graph, start, goal uint64, search_type uint, heuristic Heuristic) ([]uint64, error) {
	if heuristic == nil {
		heuristic = NullHeuristic
	}

	// Initialize the open set, according to the type of search passed in
	var openSet OpenSet
	switch search_type {
	case type_bfs:
		openSet = &OpenQueue{container.NewQueue()} // FIFO approach to expand nodes
	case type_dfs:
		openSet = &OpenStack{container.NewStack()} // LIFO approach to expand nodes
	case type_dijkstra, type_astar:
		openSet = &container.PQueue{} // Priority queue if Dijkstra or A*
	}

	closedSet := make(map[uint64]uint64) // Visited nodes, with a reference to their direct ancestor

	state := &searchstate{start, 0, 0}
	openSet.Push(state, 0)

	for openSet.Len() > 0 {
		item := openSet.Pop()
		state = item.(*searchstate)

		// Only consider non expanded nodes (not present in closedSet)
		if _, ok := closedSet[state.node]; !ok {
			// Store this node in the closed list, with a reference to its parent
			closedSet[state.node] = state.parent

			if state.node == goal {
				return calculatePath(start, goal, closedSet), nil
			}

			// Add the nodes not present in the closedSet into the openSet
			succ, ok := graph.Neighbors(state.node)
			if !ok {
				continue // Malformed Graph. Skip this node
			}

			// for depth-first search, we have to alter the order in which we add the successors to the stack,
			// to ensure items are expanded as expected (in the order they have been passed in)
			if search_type == type_dfs {
				for i, j := 0, len(succ)-1; i < j; i, j = i+1, j-1 {
					succ[i], succ[j] = succ[j], succ[i]
				}
			}

			for _, node := range succ {
				if _, ok := closedSet[node]; !ok {
					if edge, ok := graph.Edge(state.node, node); ok {
						nextState := &searchstate{node, state.node, state.cost + (int)(edge.Weight)} // TODO: Avoid the int cast!
						openSet.Push(nextState, nextState.cost+heuristic(node, goal))
					}
				}
			}
		}
	}

	return nil, errors.New("Path not found")
}

func calculatePath(start, goal uint64, closedSet map[uint64]uint64) []uint64 {
	path := make([]uint64, 0, len(closedSet))
	// fetch all the nodes in a descendant way, from goal to start
	node := goal
	for node != 0 {
		path = append(path, node)
		node = closedSet[node]
	}

	// Reverse the slice
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
