func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	temp := x

	reverse := 0

	for temp > 0 {
		reverse *= 10
		reverse += temp % 10

		temp /= 10
	}

	return x == reverse
}
