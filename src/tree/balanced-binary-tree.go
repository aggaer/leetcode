package tree

import "github.com/aggaer/leetcode/src/model"

func isBalanced(root *model.TreeNode) bool {
	if root == nil {
		return false
	}
	if abs(height(root.Left)-height(root.Right)) < 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}
}

func height(root *model.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}
