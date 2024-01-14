package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_s1657(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	Func := closeStrings
	paris := []Pair{
		{Func("abc", "bca"), true, ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	dict1, dict2 := map[byte]int{}, map[byte]int{}

	for i := 0; i < len(word1); i++ {
		dict1[word1[i]]++
		dict2[word2[i]]++
	}
	value1, value2 := []int{}, []int{}
	for i := 0; i < len(word1); i++ {
		value1 = append(value1, dict1[word1[i]])
		value2 = append(value2, dict2[word2[i]])
	}
	sort.Ints(value1)
	sort.Ints(value2)

	for i := 0; i < len(word1); i++ {
		if value1[i] != value2[i] {
			return false
		}
	}

	for key := range dict1 {
		_, ok := dict2[key]
		if !ok {
			return false
		}
	}

	return len(dict1) == len(dict2)
}
