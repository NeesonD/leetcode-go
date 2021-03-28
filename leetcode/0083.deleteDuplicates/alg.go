package _083_deleteDuplicates

func deleteDuplicates(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil {
		if fast != slow && fast.Val != slow.Val {
			slow.Next = fast
			slow = fast
		}
		fast = fast.Next
	}
	if slow != nil {
		slow.Next = nil
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
