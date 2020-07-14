package ch02

import (
	"math"
)

func Sum(a, b LinkedList) LinkedList {
	nA, nB := readInt(a), readInt(b)
	asArray := arrayFromInt(nA + nB)
	reverse(asArray)
	return NewLinkedList(asArray)
}

func readInt(l LinkedList) int {
	result := 0

	var power float64 = 0
	for n := l.First; n != nil; n = n.Next {
		result += n.Data * int(math.Pow(10.0, power))
		power++
	}

	return result
}

func arrayFromInt(n int) []int {
	if n == 0 {
		return []int{0}
	}

	maxPower := int(math.Log10(float64(n)))
	result := make([]int, 0, maxPower+1)

	for i := maxPower; i >= 0; i-- {
		base := int(math.Pow(10.0, float64(i)))
		factor := n / base
		result = append(result, factor)
		n -= factor * base
	}

	return result
}

func reverse(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func SumForward(a, b LinkedList) LinkedList {
	nA, nB := readIntForward(a), readIntForward(b)
	asArray := arrayFromInt(nA + nB)
	return NewLinkedList(asArray)
}

func readIntForward(l LinkedList) int {
	// count elements in l to know the number of digits
	count := 0
	for n := l.First; n != nil; n = n.Next {
		count++
	}

	if count == 0 {
		return 0
	}

	var result int
	power := count - 1

	for n := l.First; n != nil; n = n.Next {
		result += n.Data * int(math.Pow(10.0, float64(power)))
		power--
	}

	return result
}
