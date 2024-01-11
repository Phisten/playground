package leetcode

func missingNumber(nums []int) int {
	l := len(nums)
	ht := make(map[int]bool, l)

	for _, v := range nums {
		ht[v] = true
	}

	for i := 0; i <= l; i++ {
		if ht[i] == false {
			return i
		}
	}
	return 0
}
