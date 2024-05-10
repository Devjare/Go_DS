package test

import (
	"ds/ds"
	"testing"
)


func TestAppend(t *testing.T) {
	list := &ds.LinkedList[int]{}
	list.Append(2)
	list.Append(3)	
	list.Append(4)	
	list.Append(5)

	string_list := list.String()
	want := "2 -> 3 -> 4 -> 5 -> nil"
    if want != string_list {
        t.Fatalf(`%v is not equals to %v`, want, string_list)
    }
}
