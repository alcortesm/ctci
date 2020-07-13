package ch02

func Partition(l *LinkedList, x int) {
	// split the nodes in two lists, left and right,
	// alse remember the last element of left for later
	left := LinkedList{}
	right := LinkedList{}
	var lastLeft *Node
	for n := l.First; n != nil; n = n.Next {
		if n.Data < x {
			left.First = &Node{
				Data: n.Data,
				Next: left.First,
			}
			if lastLeft == nil {
				lastLeft = left.First
			}
		} else {
			right.First = &Node{
				Data: n.Data,
				Next: right.First,
			}
		}
	}

	// replace l with left + right
	l.First = left.First
	if lastLeft != nil {
		lastLeft.Next = right.First
	} else {
		l.First = right.First
	}
}
