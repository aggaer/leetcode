package stack

type MyStack struct {
	container []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.container = append(this.container, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ele := this.container[len(this.container)-1]
	this.container = this.container[0 : len(this.container)-1]
	return ele
}

/** Get the top element. */
func (this *MyStack) Top() int {
	con := this.container
	if len(this.container) == 0 {
		return -1
	}
	return con[len(con)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.container) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
