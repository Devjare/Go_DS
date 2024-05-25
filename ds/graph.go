package ds

import "fmt"

type Vertex[T comparable] struct {
	id int
	val T
}

type DirectedGraph[T comparable] struct {
	edges map[int][]int
	vertices []Vertex[T]
}

func (g DirectedGraph[T]) Size() int {
	if g.vertices != nil {
		return len(g.vertices)
	}
	return 0
}

func (g DirectedGraph[T]) GetValue(idx int) T {
	var value T
	if g.Size() >= 1 {
		value = g.vertices[idx].val
	}	
	return value
}

// Print the graph with values
func (g DirectedGraph[T]) PrintGraph() {
	fmt.Printf("Graph: \n")
	
	for v, edges := range g.edges {
		fmt.Printf("(%v, %v): ", v, g.vertices[v].val)
		for _, e := range edges {
			fmt.Printf("(%v, %v) ", e, g.vertices[e].val)
		}
		fmt.Print("\n")
	}
}

// Insert a new vertex into the graph
func (g *DirectedGraph[T]) AddVertex(value T) int {
	if g.Size() == 0 {
		g.vertices = []Vertex[T]{}
		g.edges = map[int][]int{}
	}
	newNode := Vertex[T]{id: g.Size() - 1, val: value}
	g.vertices = append(g.vertices, newNode)
	g.edges[g.Size() - 1] = []int{}
	return newNode.id
}

// Insert a new edge into the graph
func (g *DirectedGraph[T]) AddEdge(x, y int) {
	g.edges[x] = append(g.edges[x], y)	
}
