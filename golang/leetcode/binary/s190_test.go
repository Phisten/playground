package leetcode

import (
	"reflect"
	"testing"
)

func Test_s190(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}
	Func := hammingWeight
	paris := []Pair{
		{Func(0b00000010100101000001111010011100), 0b00111001011110000010100101000000, ""},
		{Func(0b11111111111111111111111111111101), 0b10111111111111111111111111111111, ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

// follow up 問到重複call該如何優化
// 最暴力方法應該是解過的進hashmap？

// 頭尾互換
func reverseBits(num uint32) uint32 {
	var lPeek uint32 = 0x80000000
	var rPeek uint32 = 0x00000001
	var ans uint32 = 0

	for i := 0; i < 16; i++ {
		left := lPeek & num
		right := rPeek & num

		if left > 0 {
			ans += rPeek
		}
		if right > 0 {
			ans += lPeek
		}

		lPeek >>= 1
		rPeek <<= 1
	}
	return ans
}

// 位元運算 log位元數
// https://leetcode.com/problems/reverse-bits/solutions/3784942/go-no-loop/
func reverseBits_(num uint32) uint32 {
	// These masks are straightforward
	num = (num & 0xffff0000 >> 16) | (num & 0x0000ffff << 16)
	num = (num & 0xff00ff00 >> 8) | (num & 0x00ff00ff << 8)
	num = (num & 0xf0f0f0f0 >> 4) | (num & 0x0f0f0f0f << 4)

	// 0xc = 0b1100, 0x3 = 0b0011
	num = (num & 0xcccccccc >> 2) | (num & 0x33333333 << 2)

	// 0xa = 0b10101010, 0x5 = 0b0101
	num = (num & 0xaaaaaaaa >> 1) | (num & 0x55555555 << 1)

	return num
}
