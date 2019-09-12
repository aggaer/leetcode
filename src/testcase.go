package main

import (
	"leetcode/src/model"
	"leetcode/src/tree"
)

//"---"
//["asdfghjkl","qwertyuiop","zxcvbnm"]
//3
func main() {
	//println(str.LicenseKeyFormatting("---", 3))
	//println(arr.FindMaxConsecutiveOnes([]int{1,0,1,1}))
	//println(others.ConstructRectangle(12450)[1])
	root := &model.TreeNode{Val: 1}
	root.Right = &model.TreeNode{Val: 2}
	root.Right.Left = &model.TreeNode{Val: 2}
	for _, v := range tree.FindMode(root) {
		println(v)
	}
}
