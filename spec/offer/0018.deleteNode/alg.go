package _018_deleteNode

func deleteNode(head *ListNode, val int) *ListNode {
	pre := head
	curr := head.Next
	if head.Val == val {
		return head.Next
	}
	for curr != nil {
		if curr.Val == val {
			pre.Next = curr.Next
			break
		}
		pre = curr
		curr = curr.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
