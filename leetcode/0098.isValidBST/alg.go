package _098_isValidBST

func isValidBST(root *TreeNode) bool {

	var bfs func(root, min, max *TreeNode) bool
	bfs = func(root, min, max *TreeNode) bool {
		if root == nil {
			return true
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		if min != nil && root.Val <= min.Val {
			return false
		}
		return bfs(root.Left, min, root) && bfs(root.Right, root, max)
	}
	return bfs(root, nil, nil)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
