package _052_getIntersectionNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	node1 := headA
	node2 := headB
	for node1 != nil || node2 != nil {
		if node1 == node2 {
			return node1
		}
		if node1 == nil {
			node1 = headB
		} else {
			node1 = node1.Next
		}
		if node2 == nil {
			node2 = headA
		} else {
			node2 = node2.Next
		}
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
