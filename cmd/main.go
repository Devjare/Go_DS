package main

import "ds/ds"

func main() {
	list := &ds.LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Print()
}
