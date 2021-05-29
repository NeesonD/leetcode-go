package _026_isSubStructure

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func recur(node1 *TreeNode, node2 *TreeNode) bool {
	if node2 == nil {
		return true
	}
	if node1 == nil || node1.Val != node2.Val {
		return false
	}
	return recur(node1.Left, node2.Left) && recur(node1.Right, node2.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
