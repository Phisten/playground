package leetcode

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test_s297(t *testing.T) {
	type Pair struct {
		got      interface{}
		expected interface{}
		note     string
	}

	codec := Constructor()
	root := codec.deserialize("1,2,3,null,null,4,5")

	paris := []Pair{
		{codec.serialize(root), "1,2,3,null,null,4,5", ""},
	}
	for i, v := range paris {
		if !reflect.DeepEqual(v.expected, v.got) {
			f := "i:%v Expected %v, but got %v, note:%v"
			t.Errorf(f, i, v.expected, v.got, v.note)
		}
	}
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder

	queue := []*TreeNode{root}
	for i := 0; i < len(queue); i++ {
		cur := queue[i]

		if cur != nil {
			sb.WriteString(strconv.Itoa(cur.Val))
			sb.WriteString(",")
			queue = append(queue, cur.Left, cur.Right)
		} else {
			sb.WriteString("null,")
		}
	}

	ans := sb.String()

	return ans[:len(ans)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" || data == "null" {
		return nil
	}

	strNodes := strings.Split(data, ",")
	num, _ := strconv.Atoi(strNodes[0])
	nodes := []*TreeNode{{Val: num}}

	length := len(strNodes) - 1
	for i := 0; i < length; i++ {
		j := i / 2
		isLeft := i%2 == 0

		num, err := strconv.Atoi(strNodes[i+1])
		if err == nil {
			cur := &TreeNode{Val: num}
			if isLeft {
				nodes[j].Left = cur
			} else {
				nodes[j].Right = cur
			}
			nodes = append(nodes, cur)
		}
	}
	return nodes[0]
}

/**
* Your Codec object will be instantiated and called as such:
* ser := Constructor();
* deser := Constructor();
* data := ser.serialize(root);
* ans := deser.deserialize(data);
 */
