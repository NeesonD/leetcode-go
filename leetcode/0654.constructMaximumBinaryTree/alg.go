package _654_constructMaximumBinaryTree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := 0
	start := -1
	for i, num := range nums {
		if num > start {
			start = num
			index = i
		}
	}
	t := new(TreeNode)
	t.Val = start
	left := constructMaximumBinaryTree(nums[:index])
	if index != len(nums) {
		right := constructMaximumBinaryTree(nums[index+1:])
		t.Right = right
	}
	t.Left = left
	return t
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
