package _105_buildTree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := preorder[0]
	s := -1
	for i, i2 := range inorder {
		if root == i2 {
			s = i
			break
		}
	}
	r := new(TreeNode)
	r.Val = root
	r.Left = buildTree(preorder[1:s+1], inorder[:s])
	r.Right = buildTree(preorder[s+1:], inorder[s+1:])
	return r
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
