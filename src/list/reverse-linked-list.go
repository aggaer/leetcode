package list

import "leetcode/src/model"

func reverseList(head *model.ListNode) *model.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre,next *model.ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
