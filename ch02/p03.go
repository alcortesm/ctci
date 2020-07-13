package ch02

// we assume the linked list contains the node and it is not the first,
// nor the last element.
func DeleteMiddle(l LinkedList, n *Node) {
	if n == nil {
		return
	}

	n.Data = n.Next.Data
	n.Next = n.Next.Next
}
