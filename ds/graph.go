package ds

import "fmt"

type Vertex[T comparable] struct {
	id int
	val T
}

type Graph[T comparable] struct {
	edges map[int][]int
	vertices []Vertex[T]
}

func (g Graph[T]) Size() int {
	if g.vertices != nil {
		return len(g.vertices)
	}
	return 0
}

func (g Graph[T]) GetValue(idx int) T {
	var value T
	if g.Size() >= 1 {
		value = g.vertices[idx].val
	}	
	return value
}

// Print the graph with values
func (g Graph[T]) PrintGraph() {
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
func (g *Graph[T]) AddVertex(value T) int {
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
func (g *Graph[T]) AddEdge(x, y int) {
	g.edges[x] = append(g.edges[x], y)
	g.edges[y] = append(g.edges[y], x)
}

// Checks if x is neighbor of y
func (g *Graph[T]) IsNeighbor(x, y int) int {
	for idx, v := range g.edges[y] {
		if v == x {
			return idx
		}
	}
	return -1
}

// Remove edge between x and y
func (g *Graph[T]) RemoveEdge(x, y int) {
	neighborIdx := g.IsNeighbor(x, y)
	if neighborIdx != -1 {
		g.edges[y] = append(g.edges[y][:neighborIdx], g.edges[y][neighborIdx+1:]...)
	}
}

// Remove vertex from the graph
func (g *Graph[T]) RemoveVertex(value T) int {
	if g.Size() == 0 {
		return -1
	}
	
	var vertexId int
		
	for v := range g.vertices {
		if g.vertices[v].val == value {
			vertexId = g.vertices[v].id
			g.vertices = append(g.vertices[:vertexId], g.vertices[vertexId+1:]...)
		}
	}
	
	// edges map[int][]int
	delete(g.edges, vertexId)
	for vertex, edges := range g.edges {
		neighborIdx := g.IsNeighbor(vertexId, vertex)
		if neighborIdx != -1 {
			g.edges[vertex] = append(edges[:vertex], edges[vertex+1:]...)
		}
	}

	return vertexId
}
