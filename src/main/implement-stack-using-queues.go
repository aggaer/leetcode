package main

import "stack"

func main() {
	myStack := stack.Constructor()
	myStack.Push(1)
	myStack.Push(2)
	println(myStack.Top())
}
