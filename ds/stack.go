package ds

type TailedLinkedList[T comparable] struct {
	list LinkedList[T]
	tail *LinkedList[T]
}

func (t *TailedLinkedList[T]) Append(value T) {
	t.list.Append(value)
    t.tail = t.list.head
    for t.tail.head != nil {
        t.tail = t.tail.head
    }
}

func (t *TailedLinkedList[T]) String() string {
	return t.list.String()
}

type EmptyErrorList struct { }

func (e *EmptyErrorList) Error() string {
	return "Cannot remove element, list is empty"
}

func (t *TailedLinkedList[T]) RemoveTail() (T, error) {
	
	var val T
	if t.list.head == nil {
		return val, &EmptyErrorList{}
	}
	
	val = t.tail.val
	t.list.Remove(t.tail.val)

	return val, nil
}

type Stack[T comparable] struct {
	list TailedLinkedList[T]
}

func (s *Stack[T]) Push(value T) { 
	s.list.Append(value)
}

func (s *Stack[T]) Peek() T {
	return s.list.tail.val
}

func (s *Stack[T]) Pop() (T, error) { 
	return s.list.RemoveTail()
}

func (s *Stack[T]) String() string {
	return s.list.String()
}
