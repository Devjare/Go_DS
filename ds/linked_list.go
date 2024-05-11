package ds

import (
	"bytes"
	"fmt"
)

// LinkedList represents a singly-linked list that holds
// values of any type.
type LinkedList[T comparable] struct {
	next *LinkedList[T]
	val  T
}

// Append `value` of type `T` to the current list.
func (list *LinkedList[T]) Append(value T) {
	
	newNode := &LinkedList[T]{nil, value}
	if list.next == nil {
		list.next = newNode
	} else {
		head := list.next
		for head.next != nil {
			// Once the node pointing to nil is found, set it's value	
			head = head.next
		}
		head.next = &LinkedList[T]{nil, value}
	}
}

func (list *LinkedList[T]) Get(idx int) T {
	var match T
	if idx < 0 {
		return match
	}

	current := list.next
	for i := 0; current != nil; i++ {
		if i == idx {
			return current.val
		}
		current = current.next
	}
	return match
}

func (list *LinkedList[T]) Find(value T) int {
	current := list.next
	for i := 0; current != nil; i++ {
		if current.val == value {
			return i
		}
		current = current.next
	}
	return -1
}

func (list *LinkedList[T]) InsertAt(value T, idx int) {
	
	if idx == 0 {
		// create newNode, which effectivly is the new list (pointing to the previous head link.next)
		newNode := &LinkedList[T]{val: value, next: list.next}
		// replace current head, with new head (newNode) wich already has pointer to the whole list
		list.next = newNode 
		return
	}
	
	prev := list
	current := list.next

	for i := 1; current != nil; i++ {
		if i == idx {
			newNode := &LinkedList[T]{val: value, next: current}
			prev.next = newNode
			return
		}
		prev = current
		current = current.next
	}
}

func (list *LinkedList[T]) Remove(value T) {
}

func (list *LinkedList[T]) RemoveAt(idx int) {
}


func (list *LinkedList[T]) Print() {
	current := list.next
	for current != nil {
		val := current.val
		fmt.Printf("%v -> ", val)
		current = current.next
	}
	fmt.Println("nil")
}

func (list *LinkedList[T]) String() string {
	var buf bytes.Buffer
	
	current := list.next
	for current != nil {
		buf.WriteString(fmt.Sprint(current.val))
		buf.WriteString(" -> ")
		current = current.next
	}
	buf.WriteString("nil")
	return buf.String()
}
