package main

func postSearch(root *Node, targetNum int) bool {
	// todo 也可以使用for循环实现

	if root == nil {
		return false
	}

	// 后序遍历: 先搜左边，再搜右边，再搜root (每个都要搜！不去判断大小再搜)
	totalCompare++
	if root.data > targetNum {
		if postSearch(root.left, targetNum) {
			return true
		}
	}
	if root.data < targetNum {
		if postSearch(root.right, targetNum) {
			return true
		}
	}

	if root.data == targetNum {
		return true
	}
	return false
}
