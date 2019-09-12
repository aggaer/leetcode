package list

import "github.com/aggaer/leetcode/src/model"

//快慢指针，找到中间点，反转后半部分再和前半部分对比
func IsPalindrome(head *model.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	pre, nex := &model.ListNode{}, &model.ListNode{}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		nex = slow.Next
		slow.Next = pre
		pre = slow
		slow = nex
	}
	if fast != nil {
		nex = slow.Next
	}
	for nex != nil {
		if pre.Val != nex.Val {
			return false
		}
		pre = pre.Next
		nex = nex.Next
	}
	return true
}
