package list

import "github.com/aggaer/leetcode/src/model"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RotateRight(head *model.ListNode, k int) *model.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	cnt := 1
	//首尾连成环
	cur := head
	for cur.Next != nil {
		cnt++
		cur = cur.Next
	}
	cur.Next = head
	//在cnt - k%cnt - 1处断开连接
	pre := head
	newHead := cnt - k%cnt
	for i := 0; i < newHead-1; i++ {
		pre = pre.Next
	}
	resp := pre.Next
	pre.Next = nil
	return resp
}
