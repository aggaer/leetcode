package tree

import (
	"leetcode/src/model"
)

var cur, result int

func KthSmallest(root *model.TreeNode, k int) int {
	bfs(root, k)
	return result
}

func bfs(root *model.TreeNode, k int) {
	if root == nil {
		return
	}
	bfs(root.Left, k)
	cur++
	if cur == k {
		result = root.Val
		return
	}
	bfs(root.Right, k)
}

