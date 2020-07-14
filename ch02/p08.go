package ch02

func FirstNodeInLoop(l LinkedList) *Node {
	if l.First == nil || l.First.Next == nil {
		return nil
	}

	// slow and fast runners
	s, f := l.First, l.First

	// find first collision point
	for {
		// move slow
		if s == nil {
			return nil
		}
		s = s.Next

		// move fast
		if f == nil {
			return nil
		}
		f = f.Next

		if f == nil {
			return nil
		}
		f = f.Next

		// check for collision
		if s == f {
			break
		}

	}

	// move slow to first, and increment both runners with the same
	// speed until the collision again, that should be the first node in
	// the loop
	for s = l.First; s != f; s, f = s.Next, f.Next {
	}

	return s
}
