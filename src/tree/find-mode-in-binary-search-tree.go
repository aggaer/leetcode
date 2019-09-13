package tree

import "github.com/aggaer/leetcode/src/model"

var resp []int
var pre *model.TreeNode
var cnt int

//noinspection SpellCheckingInspection
var maxcnt int

func FindMode(root *model.TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		resp = append(resp, root.Val)
		return resp
	}
	traversal(root)
	return resp
}

//若节点值等于pre节点值，cnt+1
//若不等于，则将比较cnt和max：若max<cnt，则清空resp，插入pre，并将max置为cnt，若max=cnt，则resp追加pre，若max>cnt则不处理
//pre指向当前节点值
func traversal(root *model.TreeNode) {
	if root == nil {
		return
	}
	traversal(root.Left)
	if pre != nil {
		if root.Val == pre.Val {
			cnt++
		} else {
			cnt = 1
		}
	}
	if cnt == maxcnt {
		resp = append(resp, root.Val)
	}
	if cnt > maxcnt {
		resp = resp[:0]
		resp = append(resp, root.Val)
		maxcnt = cnt
	}
	pre = root
	traversal(root.Right)
}
