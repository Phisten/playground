/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0

	var dfs func(node *ListNode, cur int)
	dfs = func(node *ListNode, cur int) {
		if node == nil {
			length = cur
			return
		}

		dfs(node.Next, cur+1)

		if length-n-1 == cur {
			if node.Next != nil {
				node.Next = node.Next.Next
			} else {
				node.Next = nil
			}
		}
	}
	dfs(head, 0)

	if length == 1 {
		return nil
	} else if n == length {
		return head.Next
	}

	return head
}

// @lc code=end
