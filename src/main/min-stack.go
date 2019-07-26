package main

import "stack"

func main() {
	minStack := stack.Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	println(minStack.GetMin())
	minStack.Pop()
	println(minStack.Top())
	println(minStack.GetMin())
}
