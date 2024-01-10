package pattern

func pattern(str string, patternString string) {
	pattern, window := map[byte]int{}, map[byte]int{}
	length := len(str)

	// init pattern
	for _, v := range patternString {
		pattern[byte(v)]++
	}

	for right, left := 0, 0; right < length; right++ {
		rightChar := str[right]

		// something update
		window[rightChar]++

		for left <= right {
			leftChar := str[left]
			left++

			// something update
			window[leftChar]--
		}
	}
}
