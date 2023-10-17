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
				Val:  1,
				Next: nil,
			},
		},
	}
	fmt.Printf("%v\n", isPalindrome(node))

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
func isPalindrome(head *ListNode) bool {
	hist := [100001]int{}

	len := 0

	for tmpHead := head; tmpHead != nil; len++ {
		hist[len] = tmpHead.Val
		tmpHead = tmpHead.Next
	}

	for i, tmpHead := len-1, head; i >= 0; i-- {
		if tmpHead.Val != hist[i] {
			return false
		}
		tmpHead = tmpHead.Next
	}

	return true
}
