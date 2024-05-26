package main

import (
	"ds/ds"
	"fmt"
)

func main() {
	graph := ds.Graph[string]{}
	graph.AddVertex("Mante")
	graph.AddVertex("Victoria")
	graph.AddVertex("Tula")
	graph.AddVertex("Reynosa")
	graph.AddVertex("Matamoros")
	graph.AddVertex("Rio Bravo")

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(1, 5)
	
	graph.AddEdge(2, 1)
	
	graph.AddEdge(3, 4)
	
	graph.AddEdge(4, 5)
	
	graph.AddEdge(5, 3)

	graph.PrintGraph()
	
	
	n1, n2 := 4, 3
	isNeighbor := graph.IsNeighbor(n1, n2)	
	fmt.Printf("%v IsNeighbor of %v = %v\n", graph.GetValue(n1), graph.GetValue(n2), isNeighbor)
	
	n1, n2 = 0, 4
	isNeighbor = graph.IsNeighbor(n1, n2)	
	fmt.Printf("%v IsNeighbor of %v = %v\n", graph.GetValue(n1), graph.GetValue(n2), isNeighbor)
	
	graph.RemoveEdge(1,0)
	graph.RemoveEdge(2,1)
	graph.RemoveEdge(4,5)
	graph.PrintGraph()

	graph.RemoveVertex("Mante")	
	graph.PrintGraph()
}
