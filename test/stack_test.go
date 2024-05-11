package test

import (
	"ds/ds"
	"testing"
)


func TestPush(t *testing.T) {
	stack := &ds.Stack[int]{}
	stack.Push(2)
	stack.Push(3)	
	stack.Push(4)	

	stack_string := stack.String()
	expected := "2 -> 3 -> 4 -> nil"
    if expected != stack_string {
        t.Fatalf(`%v is not equals to %v`, expected, stack_string)
    }
}

func TestPeek(t *testing.T) {
	stack := &ds.Stack[int]{}
	stack.Push(2)
	stack.Push(3)	
	stack.Push(4)	

	peek := stack.Peek()
	expected := 4
    if expected != peek {
        t.Fatalf(`%v is not equals to %v`, expected, peek)
    }
}

func TestPop(t *testing.T) {
	stack := &ds.Stack[int]{}
	stack.Push(2)
	stack.Push(3)	
	stack.Push(4)	

	peek, _ := stack.Pop()
	expected := 4
    if expected != peek {
        t.Fatalf(`%v is not equals to %v`, expected, peek)
    }
}

