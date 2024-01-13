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
	printNode(reverseList(node))
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
func reverseList(head *ListNode) *ListNode {
	pre := head
	cur := head
	i := 0
	for cur != nil {

		// set
		next := cur.Next
		if i == 0 {
			cur.Next = nil
		} else {
			cur.Next = pre
		}

		// next
		pre = cur
		cur = next

		// last time
		if cur == nil {
			return pre
		} else {
			i++
		}
	}

	return nil
}
