package _114_flatten

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right
	flatten(left)
	flatten(right)
	tmp := left
	if tmp == nil {
		return
	}

	for tmp.Right != nil {
		tmp = tmp.Right
	}
	tmp.Right = right
	root.Right = left
	root.Left = nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
