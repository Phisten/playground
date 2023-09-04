package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	fmt.Printf("%v -> f \n", hasCycle(node))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	curHead := head

	for i := 0; i < 10000; i++ {
		if curHead == nil || curHead.Next == nil {
			return false
		}
		curHead = curHead.Next
	}
	return true
}
