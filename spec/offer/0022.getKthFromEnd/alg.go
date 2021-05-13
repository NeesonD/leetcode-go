package _022_getKthFromEnd

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var fast, slow *ListNode
	fast = head
	slow = head
	for i := k; i > 1; i-- {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}
