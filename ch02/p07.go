package ch02

func Intersect(a, b LinkedList) *Node {
	if a.First == nil || b.First == nil {
		return nil
	}

	aLen, aLast := LenghtAndLastNode(a)
	bLen, bLast := LenghtAndLastNode(b)

	// if the last element is not the same in both lists, they do not
	// intersect.
	if aLast != bLast {
		return nil
	}

	// now we know they intersect, we need to learn where.

	// get the first nodes of each list
	na, nb := a.First, b.First

	// skip nodes in the longest list until both have the same number
	// of nodes.
	if aExcess := aLen - bLen; aExcess > 0 {
		na = getNode(a, aExcess)
	} else {
		nb = getNode(b, -aExcess)
	}

	// find the first shared node
	for ca, cb := na, nb; ca != nil && cb != nil; ca, cb = ca.Next, cb.Next {
		if ca == cb {
			return ca
		}
	}

	// unreachable
	return nil
}

func LenghtAndLastNode(l LinkedList) (int, *Node) {
	count := 0
	var n *Node

	for n = l.First; n.Next != nil; n = n.Next {
		count++
	}

	return count, n
}

// return the nth node
func getNode(l LinkedList, n int) *Node {
	result := l.First

	for i := 0; i < n; i++ {
		if result.Next == nil {
			return nil
		}
		result = result.Next
	}

	return result
}
