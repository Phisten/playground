package leetcode

func maxFrequencyElements(nums []int) int {
	maxF := 0
	dict := map[int]int{}

	for _, v := range nums {
		dict[v]++
		maxF = max(dict[v], maxF)
	}
	ams := 0
	for _, v := range dict {
		if v == maxF {
			ams += maxF
		}
	}

	return ams
}
