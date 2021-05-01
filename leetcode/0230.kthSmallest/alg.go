package _230_kthSmallest

var (
	step   = 0
	result = 0
)

func kthSmallest(root *TreeNode, k int) int {
	dfs(root, k)
	return result
}

func dfs(root *TreeNode, k int) {
	if root == nil {
		return
	}
	dfs(root.Left, k)
	step++
	if k == step {
		result = root.Val
		return
	}
	dfs(root.Right, k)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
