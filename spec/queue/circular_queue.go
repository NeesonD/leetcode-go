package queue

type MyCircularQueue struct {
	items []int
	k     int
	head  int
	tail  int
	count int
}

func Constructor(k int) MyCircularQueue {
	ints := make([]int, k)
	return MyCircularQueue{
		items: ints,
		k:     k,
		head:  0,
		tail:  0,
		count: 0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.items[this.tail] = value
	this.tail = (this.tail + 1) % this.k
	this.count++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.k
	this.count--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.items[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	if this.tail != 0 {
		return this.items[this.tail-1]
	} else {
		return this.items[len(this.items)-1]
	}
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.k
}
