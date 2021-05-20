package _025_mergeList

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	tmpNode := &ListNode{Val: -1}
//	dump := tmpNode
//	for l1 != nil && l2 != nil {
//		if l1.Val > l2.Val {
//			tmpNode.Next = l2
//			l2 = l2.Next
//			tmpNode = tmpNode.Next
//		} else {
//			tmpNode.Next = l1
//			l1 = l1.Next
//			tmpNode = tmpNode.Next
//		}
//	}
//	if l1 != nil {
//		tmpNode.Next = l1
//	}
//	if l2 != nil {
//		tmpNode.Next = l2
//	}
//	return dump.Next
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
