package _028_isSymmetric

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 != nil {
		return false
	}
	if node2 == nil && node1 != nil {
		return false
	}
	if node1 == nil && node2 == nil {
		return true
	}
	if node1.Val != node2.Val {
		return false
	}
	return helper(node1.Left, node2.Right) && helper(node1.Right, node2.Left)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
