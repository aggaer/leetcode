package list

import "model"

func removeElements(head *model.ListNode, val int) *model.ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}

func removeElements2(head *model.ListNode, val int) *model.ListNode {
	dummy := model.ListNode{}
	dummy.Next = head
	pre := &dummy
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return dummy.Next
}
