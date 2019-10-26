package tree

import (
	"github.com/aggaer/leetcode/src/model"
	"math"
)

var minDiff = math.MaxInt32
var preRoot *model.TreeNode

func getMinimumDifference(root *model.TreeNode) int {
	traversalMid(root)
	return minDiff
}

func traversalMid(root *model.TreeNode) {
	if root == nil {
		return
	}
	traversalMid(root.Left)
	if preRoot != nil {
		cur := abs(root.Val - preRoot.Val)
		if minDiff > cur {
			minDiff = cur
		}
	}
	preRoot = root
	traversalMid(root.Right)
}
