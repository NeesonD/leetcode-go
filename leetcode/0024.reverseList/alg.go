package _024_reverseList

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

type ListNode struct {
	val  int
	Next *ListNode
}

var tag *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		tag = head.Next
		return head
	}
	node := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = tag
	return node
}

func reverseMN(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseMN(head.Next, m-1, n-1)
	return head
}
