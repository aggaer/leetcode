package tree

import "model"

func minDepth(root *model.TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := minDepth(root.Left)
	rDepth := minDepth(root.Right)
	if lDepth == 0 {
		return rDepth + 1
	} else if rDepth == 0 {
		return lDepth + 1
	} else {
		return 1 + min(lDepth, rDepth)
	}
}

func min(l int, r int) int {
	if l > r {
		return r
	} else {
		return l
	}
}

//[3,9,20,null,null,15,7]
func main() {
	root := model.TreeNode{Val: 3}
	root.Left = &model.TreeNode{Val: 9}
	root.Right = &model.TreeNode{Val: 20}
	root.Right.Left = &model.TreeNode{Val: 15}
	root.Right.Right = &model.TreeNode{Val: 7}
	print(minDepth(&root))
}
