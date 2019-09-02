package list

import "leetcode/src/model"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *model.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
