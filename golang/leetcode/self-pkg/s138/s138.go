package main

import (
	"fmt"
)

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {
	arr := [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	head := createLinkedList(arr)

	fmt.Printf("%v -> %v \n", linkedListToSlice(copyRandomList(head)), arr)

}

func createLinkedList(arr [][]int) *Node {
	if len(arr) == 0 {
		return nil
	}

	nodeMap := make(map[int]*Node)
	head := &Node{Val: arr[0][0]}
	nodeMap[0] = head
	current := head

	for i := 1; i < len(arr); i++ {
		newNode := &Node{Val: arr[i][0]}
		current.Next = newNode
		current = newNode
		nodeMap[i] = newNode
	}

	for i := 0; i < len(arr); i++ {
		if arr[i][1] != -1 {
			nodeMap[i].Random = nodeMap[arr[i][1]]
		}
	}

	return head
}
func linkedListToSlice(head *Node) [][]int {
	result := [][]int{}
	indexMap := map[*Node]int{}
	current := head
	index := 0

	for current != nil {
		indexMap[current] = index
		index++
		current = current.Next
	}

	current = head
	for current != nil {
		randomIndex := -1
		if current.Random != nil {
			randomIndex = indexMap[current.Random]
		}
		result = append(result, []int{current.Val, randomIndex})
		current = current.Next
	}

	return result
}

func copyRandomList(head *Node) *Node {
	arr := linkedListToSlice(head)
	newHead := createLinkedList(arr)
	return newHead
}
