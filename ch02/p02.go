package ch02

import (
	"fmt"
)

func KthToLast(l LinkedList, k int) (int, error) {
	if l.First == nil {
		return 0, fmt.Errorf("empty list")
	}

	jumps := 0
	var behind *Node = nil

	for ahead := l.First; ahead != nil; ahead = ahead.Next {
		if behind != nil {
			behind = behind.Next
		} else {
			if jumps == k {
				behind = l.First
			}
		}

		jumps++
	}

	if behind == nil {
		return 0, fmt.Errorf("list is too short")
	}

	return behind.Data, nil
}
