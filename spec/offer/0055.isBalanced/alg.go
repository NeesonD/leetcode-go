package _055_isBalanced

import "math"

func isBalanced(root *TreeNode) bool {
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		if left == -1 {
			return -1
		}
		right := helper(root.Right)
		if right == -1 {
			return -1
		}
		if math.Abs(float64(left-right)) < 2 {
			return int(math.Max(float64(left), float64(right)) + 1)
		} else {
			return -1
		}
	}
	i := helper(root)
	if i == -1 {
		return false
	}
	return true

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
