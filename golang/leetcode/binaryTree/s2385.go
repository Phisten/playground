package leetcode

/**

1. 由root前序找start, 回傳深度 沒找到就回最大深度
		LR相加就是start往parent逆向感染的長度

2. start 往下檢測深度

3. 1和2取max
不行, 最長路徑可以不經過node也不在start下面


1. 走訪一次, 把所有node的parent存進hashMap供存取, 順便找startNode
2. 從startNode擴張找最長

*/

func amountOfTime(root *TreeNode, start int) int {
	var startNodeChildLength int = 0
	var startNode *TreeNode
	parent := map[*TreeNode]*TreeNode{}

	var findLength func(node *TreeNode, start int) int
	findLength = func(node *TreeNode, start int) int {
		if node == nil {
			return 0
		}

		if node.Left != nil {
			parent[node.Left] = node
		}
		if node.Right != nil {
			parent[node.Right] = node
		}
		length := max(findLength(node.Left, start), findLength(node.Right, start))

		if node.Val == start {
			startNodeChildLength = length
			startNode = node
		}
		return 1 + length
	}
	findLength(root, start)

	var inverseFind func(node *TreeNode, prev *TreeNode) int
	inverseFind = func(node *TreeNode, prev *TreeNode) int {
		if node == nil {
			return 0
		}

		length := 0
		if parent[node] != prev {
			length = inverseFind(parent[node], node)
		}
		if node.Left != prev {
			length = max(length, inverseFind(node.Left, node))
		}
		if node.Right != prev {
			length = max(length, inverseFind(node.Right, node))
		}

		return 1 + length
	}
	inverseFindLength := inverseFind(parent[startNode], startNode)

	return max(startNodeChildLength, inverseFindLength)
}
