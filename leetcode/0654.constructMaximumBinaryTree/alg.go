package _654_constructMaximumBinaryTree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := -1
	maxIndex := -1
	for i, num := range nums {
		if num > max {
			max = num
			maxIndex = i
		}
	}
	t := new(TreeNode)
	t.Val = max
	if maxIndex != len(nums)-1 {
		t.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	}
	t.Left = constructMaximumBinaryTree(nums[:maxIndex])
	return t
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
