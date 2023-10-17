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
	printNode(middleNode(node))
	// fmt.Printf("%v -> 3 2 1 \n", printNode(reverseList(node)))

}

func printNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	mid := head

	for n := 1; head != nil; n++ {
		if n%2 == 0 {
			mid = mid.Next
		}
		head = head.Next
	}

	return mid
}
