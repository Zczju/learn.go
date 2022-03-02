package main

func midSearch(root *Node, targetNum int) bool {
	// todo 也可以使用for循环实现

	if root == nil {
		return false
	}

	// 中序遍历: 先搜左边，再搜root，再搜右边 (每个都要搜！不去判断大小再搜)
	totalCompare++
	if midSearch(root.left, targetNum) {

		return true
	}
	if root.data == targetNum {
		return true
	}
	if midSearch(root.right, targetNum) {
		return true
	}
	return false
}
