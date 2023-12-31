package leetcode

import (
	"strings"
	"testing"
)

func Test_ReplaceLeetcodeBrackets(t *testing.T) {
	t.Errorf(replaceBrackets("\n [[0,2],[-3,3],[-2,5]]"))
}

func replaceBrackets(str string) string {
	replaced := strings.ReplaceAll(str, "[", "{")
	replaced = strings.ReplaceAll(replaced, "]", "}")
	return replaced
}
