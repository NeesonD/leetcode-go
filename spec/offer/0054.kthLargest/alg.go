package _054_kthLargest

var count, result int

func kthLargest(root *TreeNode, k int) int {
	count = 0
	kthLargestH(root, k)
	return result
}

func kthLargestH(root *TreeNode, k int) {
	if root == nil {
		return
	}
	kthLargestH(root.Right, k)
	count++
	if k == count {
		result = root.Val
		return
	}
	kthLargestH(root.Left, k)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
