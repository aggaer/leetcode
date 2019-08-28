package main

import (
	"model"
	"tree"
)

func main() {
	root := &model.TreeNode{Val: 1}
	root.Left = &model.TreeNode{Val: 2}
	root.Right = &model.TreeNode{Val: 3}
	root.Left.Left = &model.TreeNode{Val: 5}
	paths := tree.BinaryTreePaths(root)
	for _, v := range paths {
		println(v)
	}
}
