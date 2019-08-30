package tree

import (
	"model"
	"strconv"
)

func BinaryTreePaths(root *model.TreeNode) []string {
	resp := traversalPaths(root, "", []string{})
	return resp
}

func traversalPaths(root *model.TreeNode, str string, result []string) []string {
	if root == nil {
		return result
	}
	str += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		result = append(result, str)
		return result
	}
	str += "->"
	result = traversalPaths(root.Left, str, result)
	result = traversalPaths(root.Right, str, result)
	return result
}
