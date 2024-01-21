/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	} else if l == 2 {
		return max(nums[0], nums[1])
	} else if l == 3 {
		return max(nums[0]+nums[2], nums[1])
	}

	var dp []int = make([]int, 3)
	dp[2] = nums[0]
	dp[1] = nums[1]
	dp[0] = max(nums[2]+nums[0], nums[1])

	for i := 3; i < l; i++ {
		cur := nums[i] + max(dp[1], dp[2])
		dp = append([]int{cur}, dp[:len(dp)-1]...)
	}

	return max(dp[0], dp[1])
}

// @lc code=end
