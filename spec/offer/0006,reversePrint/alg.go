package _006_reversePrint

func reversePrint(head *ListNode) []int {
	result := make([]int, 0)
	var f func(head *ListNode) []int
	f = func(head *ListNode) []int {
		if head == nil {
			return result
		}
		f(head.Next)
		result = append(result, head.Val)
		return result
	}
	f(head)
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
