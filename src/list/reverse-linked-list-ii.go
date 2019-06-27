package list

import "model"

func reverseBetween(head *model.ListNode, m int, n int) *model.ListNode {
	dummy := model.ListNode{Val: -1, Next: head}
	pre := &dummy
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := m; i < n; i++ {
		nex := cur.Next
		cur.Next = nex.Next
		nex.Next = cur
		pre.Next = nex
	}
	return dummy.Next
}
