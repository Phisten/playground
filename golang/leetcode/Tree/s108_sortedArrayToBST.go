package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)

	if l == 0 {
		return nil
	}

	midIdx := l / 2

	left := sortedArrayToBST(nums[:midIdx])
	right := sortedArrayToBST(nums[midIdx+1:])

	top := &TreeNode{
		Val:   nums[midIdx],
		Left:  left,
		Right: right,
	}

	return top

}
