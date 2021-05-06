package _106_buildTree

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := postorder[len(postorder)-1]

	s := -1
	for i, i2 := range inorder {
		if i2 == root {
			s = i
			break
		}
	}
	t := new(TreeNode)
	t.Val = root
	t.Left = buildTree(inorder[:s], postorder[:s])
	t.Right = buildTree(inorder[s+1:], postorder[s:len(postorder)-1])
	return t
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
