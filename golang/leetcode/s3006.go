package leetcode

func beautifulIndices(s string, a string, b string, k int) []int {
	length := len(s)

	aPos := []int{}
	bPos := []int{}

	for i := 0; i < length; i++ {
		if i <= length-len(a) {
			for j := 0; j < len(a); j++ {
				if s[i+j] != a[j] {
					break
				} else if j == len(a)-1 {
					aPos = append(aPos, i)
				}
			}
		}

		if i <= length-len(b) {
			for j := 0; j < len(b); j++ {
				if s[i+j] != b[j] {
					break
				} else if j == len(b)-1 {
					bPos = append(bPos, i)
				}
			}
		}
	}

	ans := []int{}
	for i := 0; i < len(aPos); i++ {
		for j := 0; j < len(bPos); j++ {
			if abs(aPos[i]-bPos[j]) <= k {
				ans = append(ans, aPos[i])
				break
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
