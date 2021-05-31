package _033_verifyPostorder

func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}
	p := left
	for postorder[p] < postorder[right] {
		p++
	}
	index := p
	for postorder[p] > postorder[right] {
		p++
	}
	return p == right && recur(postorder, left, index-1) && recur(postorder, index, right-1)
}
