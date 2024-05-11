package ds

import (
	"bytes"
	"fmt"
)

// LinkedList represents a singly-linked list that holds
// values of any type.
type LinkedList[T comparable] struct {
	head *LinkedList[T]
	val  T
}

// Append `value` of type `T` to the current list.
func (list *LinkedList[T]) Append(value T) {
	
	newNode := &LinkedList[T]{nil, value}
	if list.head == nil {
		list.head = newNode
	} else {
		head := list.head
		for head.head != nil {
			// Once the node pointing to nil is found, set it's value	
			head = head.head
		}
		head.head = &LinkedList[T]{nil, value}
	}
}

func (list *LinkedList[T]) Get(idx int) T {
	var match T
	if idx < 0 {
		return match
	}

	current := list.head
	for i := 0; current != nil; i++ {
		if i == idx {
			return current.val
		}
		current = current.head
	}
	return match
}

func (list *LinkedList[T]) Find(value T) int {
	current := list.head
	for i := 0; current != nil; i++ {
		if current.val == value {
			return i
		}
		current = current.head
	}
	return -1
}

func (list *LinkedList[T]) InsertAt(value T, idx int) {
	
	if idx == 0 {
		// create newNode, which effectivly is the new list (pointing to the previous head link.next)
		newNode := &LinkedList[T]{val: value, head: list.head}
		// replace current head, with new head (newNode) wich already has pointer to the whole list
		list.head = newNode 
		return
	}
	
	prev := list
	current := list.head

	for i := 1; current != nil; i++ {
		if i == idx {
			newNode := &LinkedList[T]{val: value, head: current}
			prev.head = newNode
			return
		}
		prev = current
		current = current.head
	}
}

// Removes element matching `value`
// returns the index of the removed value, -1 if the element doesn't exists
func (list *LinkedList[T]) Remove(value T) int {
	prev := list
	current := list.head
	for i := 1; current != nil; i++ {
		if current.val == value {
			prev.head = current.head
			return i
		}
		prev = current
		current = current.head
	}
	return -1
}

func (list *LinkedList[T]) RemoveAt(idx int) T {
	var val T
	prev := list
	current := list.head
	for i := 0; current != nil; i++ {
		if i == idx {
			val = current.val
			prev.head = current.head
			return val
		}
		prev = current
		current = current.head
	}
	return val
}


func (list *LinkedList[T]) Print() {
	current := list.head
	for current != nil {
		val := current.val
		fmt.Printf("%v -> ", val)
		current = current.head
	}
	fmt.Println("nil")
}

func (list *LinkedList[T]) String() string {
	var buf bytes.Buffer
	
	current := list.head
	for current != nil {
		buf.WriteString(fmt.Sprint(current.val))
		buf.WriteString(" -> ")
		current = current.head
	}
	buf.WriteString("nil")
	return buf.String()
}
