package ch02

import "fmt"

type LinkedList struct {
	First *Node
}

type Node struct {
	Data int
	Next *Node
}

func NewLinkedList(a []int) LinkedList {
	result := LinkedList{}

	if len(a) == 0 {
		return result
	}

	result.First = &Node{
		Data: a[0],
	}

	lastData := result.First

	for _, e := range a[1:] {
		lastData.Next = &Node{
			Data: e,
		}
		lastData = lastData.Next
	}

	return result
}

func (l LinkedList) Clone() LinkedList {
	if l.First == nil {
		return LinkedList{}
	}

	result := LinkedList{
		First: &Node{
			Data: l.First.Data,
		},
	}

	lastDuplicated := result.First

	for n := l.First.Next; n != nil; n = n.Next {
		lastDuplicated.Next = &Node{
			Data: n.Data,
		}
		lastDuplicated = lastDuplicated.Next
	}

	return result
}

func (l LinkedList) String() string {
	if l.First == nil {
		return ""
	}

	a := []int{}

	for n := l.First; n != nil; n = n.Next {
		a = append(a, n.Data)
	}

	return fmt.Sprintf("%d", a)

}

func (l LinkedList) Equals(ll LinkedList) bool {
	return l.String() == ll.String()
}
