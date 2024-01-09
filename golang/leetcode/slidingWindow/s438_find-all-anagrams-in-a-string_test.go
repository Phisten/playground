package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_s438(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	Func := findAnagrams
	paris := []Pair{
		{Func("1000", "1110"), "10110", ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

func findAnagrams(s string, p string) []int {
	window, pattern := map[byte]int{}, map[byte]int{}
	for _, v := range p {
		pattern[byte(v)]++
	}

	ans := []int{}

	length, pLength := len(s), len(p)

	left := 0
	for right := 0; right < length; right++ {
		rightChar := s[right]

		window[rightChar]++

		if right-left+1 == pLength {
			fmt.Printf("%v,%v\n", window, pattern)
			if reflect.DeepEqual(window, pattern) {
				ans = append(ans, left)
			}

			leftChar := s[left]
			window[leftChar]--
			if window[leftChar] == 0 {
				delete(window, leftChar)
			}
			left++
		}
	}

	return ans
}
