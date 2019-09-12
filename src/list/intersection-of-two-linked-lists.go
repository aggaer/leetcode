package list

import "github.com/aggaer/leetcode/src/model"

func getIntersectionNode(headA, headB *model.ListNode) *model.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	tail := headB
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = headB
	slow, fast := headA, headA
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = headA
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			tail.Next = nil
			return fast
		}
	}
	tail.Next = nil
	return nil
}
