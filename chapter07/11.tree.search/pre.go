package main

func preSearch(root *Node, targetNum int) bool {
	// todo 也可以使用for循环实现

	if root == nil {
		return false
	}

	// 前序遍历: 先搜root，再搜左边，再搜右边 (每个都要搜！不去判断大小再搜)
	totalCompare++
	if root.data == targetNum {
		return true
	}
	if preSearch(root.left, targetNum) {

		return true
	}
	if preSearch(root.right, targetNum) {
		return true
	}
	return false
}
