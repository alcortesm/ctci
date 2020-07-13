package ch02

func RemoveDupsUsingSpace(l LinkedList) {
	if l.First == nil {
		return
	}

	seen := make(map[int]struct{})
	seen[l.First.Data] = struct{}{}

	for prev := l.First; prev.Next != nil; {
		current := prev.Next

		if _, ok := seen[current.Data]; ok {
			prev.Next = current.Next // remove next
			continue
		}

		seen[current.Data] = struct{}{}
		prev = prev.Next
	}
}

func RemoveDupsUsingTime(l LinkedList) {
	if l.First == nil {
		return
	}

	// slow walker
	for s := l.First; s != nil && s.Next != nil; s = s.Next {
		// fast walker
		for f := s; f.Next != nil; {
			if f.Next.Data == s.Data {
				f.Next = f.Next.Next
			} else {
				f = f.Next
			}
		}
	}
}
