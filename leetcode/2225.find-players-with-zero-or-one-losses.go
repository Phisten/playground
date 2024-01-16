/*
 * @lc app=leetcode id=2225 lang=golang
 *
 * [2225] Find Players With Zero or One Losses
 */

// @lc code=start
func findWinners(matches [][]int) [][]int {
	ans := [][]int{{}, {}}

	dict := map[int]int{}

	for i := 0; i < len(matches); i++ {
		dict[matches[i][1]]++
		if _, ok := dict[matches[i][0]]; !ok {
			dict[matches[i][0]] = 200000
		}
	}

	for idx, v := range dict {
		if v > 0 {
			switch v % 200000 {
			case 0:
				ans[0] = append(ans[0], idx)
			case 1:
				ans[1] = append(ans[1], idx)
			}
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])

	return ans
}

// @lc code=end
