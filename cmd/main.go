package main

import (
	"ds/ds"
	"fmt"
)

func main() {
	list := &ds.LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Print()
	idx := list.Find(3)
	fmt.Printf("Found %v at index: %v", 1, idx)
}
