package _146_lru

type LRUCache struct {
	capacity   int
	dict       map[int]*Node
	doubleList *DoubleList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:   capacity,
		dict:       make(map[int]*Node),
		doubleList: NewDoubleList(capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.dict[key]
	if !ok {
		return -1
	}
	this.markRecently(key)
	return node.v
}

func (this *LRUCache) Put(key, value int) {
	_, ok := this.dict[key]
	if ok {
		this.deleteKey(key)
	}
	// 如果满了，清理掉尾部元素
	if len(this.dict) == this.capacity {
		this.removeLeastRecently()
	}
	this.addRecently(key, value)
}

func (this *LRUCache) markRecently(key int) {
	node := this.dict[key]
	this.doubleList.remove(node)
	this.doubleList.addLast(node)
}

func (this *LRUCache) addRecently(key, value int) {
	node := &Node{
		Prev:  nil,
		Next:  nil,
		k:     key,
		v:     value,
		dummy: false,
	}
	this.dict[key] = node
	this.doubleList.addLast(node)
}

func (this *LRUCache) deleteKey(key int) {
	node, ok := this.dict[key]
	if !ok {
		return
	}
	delete(this.dict, key)
	this.doubleList.remove(node)
}

func (this *LRUCache) removeLeastRecently() {
	if this.doubleList.head == this.doubleList.tail {
		return
	}
	node := this.doubleList.head.Next
	this.doubleList.removeFirst()
	delete(this.dict, node.k)
}

type Node struct {
	Prev  *Node
	Next  *Node
	k     int
	v     int
	dummy bool
}

type DoubleList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleList(size int) *DoubleList {
	head := &Node{k: -1, v: -1, dummy: true}
	tail := &Node{k: -1, v: -1, dummy: true}
	tail.Prev = head

	return &DoubleList{
		head: head,
		tail: tail,
		size: size,
	}
}

func (n *DoubleList) addLast(node *Node) {
	node.Prev = n.tail.Prev
	node.Next = n.tail
	n.tail.Prev.Next = node
	n.tail.Prev = node
	n.size++
}

func (n *DoubleList) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	n.size--
}

func (n *DoubleList) removeFirst() *Node {
	if n.head == n.tail {
		return nil
	}
	first := n.head.Next
	n.remove(first)
	n.size--
	return first
}

func (n *DoubleList) Size() int {
	return n.size
}
