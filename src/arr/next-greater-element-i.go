package arr

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	var resp []int
	mp := make(map[int]int)
	myStack := Constructor()
	for _, v := range nums2 {
		for !myStack.Empty() && myStack.Top() < v {
			mp[myStack.Pop()] = v
		}
		myStack.Push(v)
	}
	for _, v := range nums1 {

		if i, ok := mp[v]; ok {
			resp = append(resp, i)
		} else {
			resp = append(resp, -1)
		}
	}
	return resp
}

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
