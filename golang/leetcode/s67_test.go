package leetcode

import (
	"reflect"
	"testing"
)

func Test_s67(t *testing.T) {
	type Pair struct {
		got      string
		expected string
		note     string
	}

	Func := addBinary
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

// 兩字串從尾部往回跑到最頭, 相加並拼接字串
func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	lc := max(len(a), len(b)) + 1

	sum := make([]byte, lc)

	for i := 1; i <= lc; i++ {
		aNum := byte(0)
		if la-i >= 0 {
			aNum = a[la-i] - 48
		}
		bNum := byte(0)
		if lb-i >= 0 {
			bNum = b[lb-i] - 48
		}
		sum[lc-i] = sum[lc-i] + aNum + bNum
		if sum[lc-i] > 1 {
			sum[lc-i] -= 2
			sum[lc-i-1] += 1
		}
		sum[lc-i] += 48
	}

	if sum[0] == 48 {
		return string(sum[1:])
	}
	return string(sum)
}

// 11
// 01
//010

//  11
// 111
//1000

// 11
// 11
//110
