package tree

import "model"

func isUnivalTree(root *model.TreeNode) bool {
	isLeft := root.Left == nil || root.Val == root.Left.Val && isUnivalTree(root.Left)
	isRight := root.Right == nil || root.Val == root.Right.Val && isUnivalTree(root.Right)
	return isLeft && isRight
}
