package ch02

func IsPalindrome(l LinkedList) bool {
	r := newReverse(l)
	return l.Equals(r)
}

func newReverse(l LinkedList) LinkedList {
	result := LinkedList{}

	for n := l.First; n != nil; n = n.Next {
		result.First = &Node{
			Data: n.Data,
			Next: result.First,
		}
	}

	return result
}
