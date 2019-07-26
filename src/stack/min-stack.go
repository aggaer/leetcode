package stack


type MinStack struct {
	container []int
	min       []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.min) == 0 || x <= this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	}
	this.container = append(this.container, x)
}

func (this *MinStack) Pop() {
	if this.min[len(this.min)-1] == this.container [len(this.container)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.container = this.container [:len(this.container)-1]
}

func (this *MinStack) Top() int {
	return this.container[len(this.container)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */