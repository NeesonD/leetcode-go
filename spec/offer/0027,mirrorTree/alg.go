package _027_mirrorTree

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Right)
	right := mirrorTree(root.Left)
	root.Left = left
	root.Right = right
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
