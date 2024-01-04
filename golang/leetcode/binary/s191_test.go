package leetcode

import (
	"reflect"
	"testing"
)

func Test_s191(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	Func := hammingWeight
	paris := []Pair{
		{Func(0b00000000000000000000000000001011), 3, ""},
		{Func(0b00000000000000000000000010000000), 1, ""},
		{Func(0b11111111111111111111111111111101), 31, ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

func hammingWeight(num uint32) int {
	var match uint32 = 1
	ans := 0
	for i := 0; i < 32; i++ {

		if match&num > 0 {
			ans++
		}
		match = match << 1
	}
	return ans
}
