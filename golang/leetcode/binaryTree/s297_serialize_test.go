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
	var ans strings.Builder

	queue := []*TreeNode{root}
	for i := 0; i < len(queue); i++ {
		cur := queue[i]

		if cur != nil {
			ans.WriteString(strconv.Itoa(cur.Val))
			queue = append(queue, cur.Left, cur.Right)
		} else {
			ans.WriteString("null")
		}
		if i < len(queue) {
			ans.WriteString(",")
		}
	}

	return ans.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var root *TreeNode
	strNodes := strings.Split(data, ",")
	nodes := []*TreeNode{}

	if len(strNodes) > 0 {
		num, nilRoot := strconv.Atoi(strNodes[0])
		if nilRoot == nil {
			root = &TreeNode{
				Val:   num,
				Left:  nil,
				Right: nil,
			}
			nodes = append(nodes, root)
		}
	}

	for i := 1; i < len(strNodes); i++ {
		j := (i - 1) / 2
		isLeft := i%2 == 1

		numL, nilL := strconv.Atoi(strNodes[i])
		if nilL == nil {
			node := &TreeNode{
				Val:   numL,
				Left:  nil,
				Right: nil,
			}
			if isLeft {
				nodes[j].Left = node
			} else {
				nodes[j].Right = node
			}
			nodes = append(nodes, node)
		}
	}
	return root
}

/**
* Your Codec object will be instantiated and called as such:
* ser := Constructor();
* deser := Constructor();
* data := ser.serialize(root);
* ans := deser.deserialize(data);
 */
