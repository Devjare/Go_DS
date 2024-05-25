package main

import (
	"ds/ds"
)

func main() {
	graph := ds.DirectedGraph[string]{}
	graph.AddVertex("Mante")
	graph.AddVertex("Victoria")
	graph.AddVertex("Tula")
	graph.AddVertex("Reynosa")
	graph.AddVertex("Matamoros")
	graph.AddVertex("Rio Bravo")

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	
	graph.AddEdge(1, 0)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(1, 5)
	
	graph.AddEdge(2, 1)
	
	graph.AddEdge(3, 3)
	graph.AddEdge(3, 4)
	
	graph.AddEdge(4, 5)
	
	graph.AddEdge(5, 3)
	graph.AddEdge(5, 4)

	graph.PrintGraph()
}
