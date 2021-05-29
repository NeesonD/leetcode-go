package _030_minstack

type MinStack struct {
	StackA []int
	StackB []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := new(MinStack)
	m.StackA = make([]int, 0)
	m.StackB = make([]int, 0)
	return *m
}

func (this *MinStack) Push(x int) {
	this.StackA = append(this.StackA, x)
	if len(this.StackB) == 0 || this.StackB[len(this.StackB)-1] >= x {
		this.StackB = append(this.StackB, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.StackA) == 0 {
		return
	}
	if len(this.StackB) != 0 && this.StackB[len(this.StackB)-1] == this.StackA[len(this.StackA)-1] {
		this.StackB = this.StackB[:len(this.StackB)-1]
	}
	this.StackA = this.StackA[:len(this.StackA)-1]
}

func (this *MinStack) Top() int {
	if len(this.StackA) == 0 {
		return -1
	}
	return this.StackA[len(this.StackA)-1]
}

func (this *MinStack) Min() int {
	if len(this.StackB) == 0 {
		return -1
	}
	return this.StackB[len(this.StackB)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
