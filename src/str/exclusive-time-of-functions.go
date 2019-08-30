package str

import (
	"strconv"
	"strings"
)

func ExclusiveTime(n int, logs []string) []int {
	resp := make([]int, n)
	countStack := Constructor()
	timeStack := Constructor()
	for _, v := range logs {
		left := strings.Index(v, ":")
		idx, _ := strconv.Atoi(v[0:left])
		v = v[left+1:]
		right := strings.Index(v, ":")
		flag := v[0:right]
		cost, _ := strconv.Atoi(v[right+1:])
		if flag == "start" {
			countStack.Push(0)
			timeStack.Push(cost)
		} else {
			timeDiff := cost - timeStack.Pop() + 1
			resp[idx] += timeDiff - countStack.Pop()
			if ! countStack.Empty() {
				countStack.Push(countStack.Pop() + timeDiff)
			}
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
