package leetcode

type NodeStack []*TreeNode

func (s *NodeStack) pop() *TreeNode {
	popNode := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return popNode
}
func (s *NodeStack) push(node *TreeNode) {
	*s = append(*s, node)
}
func (s NodeStack) isEmpty() bool {
	return len(s) == 0
}
func (s NodeStack) peek() *TreeNode {
	return s[len(s)-1]
}
