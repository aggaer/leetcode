package main

import (
	"list"
	"model"
)

func main() {
	head := &model.ListNode{Val: 1}
	head.Next = &model.ListNode{Val: 2}
	head.Next.Next = &model.ListNode{Val: 3}
	head.Next.Next.Next = &model.ListNode{Val: 2}
	head.Next.Next.Next.Next = &model.ListNode{Val: 2}
	print(list.IsPalindrome(head))
}
