package _701_insertIntoBST

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)

	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)

	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
