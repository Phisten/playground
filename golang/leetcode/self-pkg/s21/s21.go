package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(arr []int) *ListNode {
	var head, tail *ListNode
	for _, val := range arr {
		node := &ListNode{Val: val}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	return head
}

// 打印链表
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	list1Arr := []int{1}
	list2Arr := []int{}
	// list1Arr := []int{1, 2, 4}
	// list2Arr := []int{1, 3, 4}

	list1 := createLinkedList(list1Arr)
	list2 := createLinkedList(list2Arr)

	printLinkedList(list1)
	printLinkedList(list2)
	printLinkedList(mergeTwoLists(list1, list2))
	// fmt.Printf("%v ->  \n", mergeTwoLists(list1, list2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail *ListNode

	// init node
	head = &ListNode{}
	if list1 == nil && list2 == nil {
		return nil
	} else if list2 == nil || (list1 != nil && list1.Val < list2.Val) {
		head.Val = list1.Val
		list1 = list1.Next
	} else {
		head.Val = list2.Val
		list2 = list2.Next
	}
	tail = head

	for {
		if list1 == nil && list2 == nil {
			return head
		} else if list2 == nil || (list1 != nil && list1.Val < list2.Val) {
			node := &ListNode{Val: list1.Val, Next: nil}
			tail.Next = node
			list1 = list1.Next
		} else {
			node := &ListNode{Val: list2.Val, Next: nil}
			tail.Next = node
			list2 = list2.Next
		}
		tail = tail.Next
	}
}
