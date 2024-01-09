package leetcode

func minWindow(s string, t string) string {
	target, window := map[byte]int{}, map[byte]int{}

	lenS := len(s)
	ansL, ansR := 0, lenS

	for _, v := range t {
		target[byte(v)] += 1
	}
	valid, targetValid := 0, 0
	for range target {
		targetValid++
	}

	for left, right := 0, 0; right < lenS; right++ {
		rightChar := s[right]

		if target[rightChar] > 0 {
			window[rightChar]++

			if window[rightChar] == target[rightChar] {
				valid++
			}
		}

		for left <= right && valid == targetValid {
			if (ansR - ansL) > (right - left) {
				ansL = left
				ansR = right
			}

			leftChar := s[left]
			left++

			if target[leftChar] > 0 {
				window[leftChar]--

				if window[leftChar] < target[leftChar] {
					valid--
				}
			}
		}
	}

	if ansR < lenS {
		return s[ansL : ansR+1]
	}
	return ""
}
