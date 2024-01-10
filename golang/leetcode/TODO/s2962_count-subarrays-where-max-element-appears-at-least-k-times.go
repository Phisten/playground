package leetcode

/*
剛好 K 個 = 最多 K 個 - 最多 K-1 個
（子集合數）
At least K = (n*(n+1)/2 - lessThanK
*/

func countSubarrays(nums []int, k int) int64 {
	length := len(nums)
	maxNum := nums[0]
	for _, n := range nums {
		if n > maxNum {
			maxNum = n
		}
	}

	window := map[int]int{}
	ans := int64(0)

	for right, left := 0, 0; right < length; right++ {
		cur := nums[right]

		window[cur]++
		if window[maxNum] >= k {
			ans += int64(length - right)
		}

		for left <= right && window[maxNum] >= k {
			cur = nums[left]
			window[cur]--

			if window[maxNum] >= k {
				ans += int64(length - right)
			}

			left++
		}
	}

	return ans
}
