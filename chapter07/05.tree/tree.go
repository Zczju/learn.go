package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
	root  *Node
}

func main() {
	root := buildTree()
	root = insertTreeNode(root, &Node{data: 43})
	root = insertTreeNode(root, &Node{data: 28})
	fmt.Println("done")

	fmt.Println("delete node")
	deleteNode(root, 43)
	deleteNode(root, 51)
	deleteNode(root, 0)
	fmt.Println("complete delete")
}

func buildTree() *Node {
	n1 := &Node{data: 51}
	n2 := &Node{data: 35}
	n3 := &Node{data: 65}

	n1.left = n2
	n1.right = n3

	n2.root = n1
	n3.root = n1

	return n1
}

func insertTreeNode(root *Node, newNode *Node) *Node {
	if root == nil {
		return newNode
	}

	if newNode.data == root.data {
		// 新数据已经在了，直接return掉
		return root
	}

	// 使用循环的方式插入树节点
	//tmpNode := root
	//for {
	//	if newNode.data > tmpNode.data {
	//		if tmpNode.right == nil {
	//			tmpNode.right = newNode
	//			return root
	//		}
	//		tmpNode = tmpNode.right
	//	} else {
	//		if newNode.data < tmpNode.data {
	//			if tmpNode.left == nil {
	//				tmpNode.left = newNode
	//				return root
	//			}
	//			tmpNode = tmpNode.left
	//		}
	//	}
	//}

	// 使用递归的方式插入树节点
	if newNode.data < root.data {
		if root.left == nil {
			// 找到位置，插入数据
			root.left = newNode
			newNode.root = root
			return root
		} else {
			insertTreeNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			// 找到位置，插入数据
			root.right = newNode
			newNode.root = root
			return root
		} else {
			insertTreeNode(root.right, newNode)
		}
	}
	return root
}

// 用于理解删除叶子节点的过程
func deleteLeafNode(deleteNode *Node, data int) *Node {
	DNLeft := deleteNode // 两个指头，一个指向要删除的节点，一个指向该节点的root
	if DNLeft.data == data && DNLeft.left == nil && DNLeft.right == nil {
		DNRight := DNLeft
		DNLeft = DNLeft.root
		if DNLeft.left == DNRight {
			// 删除左边叶节点
			DNRight.root = nil
			DNLeft.left = nil
			return DNLeft
		}
		if DNLeft.right == DNRight {
			// 删除右边叶节点
			DNRight.root = nil
			DNLeft.right = nil
			return DNLeft
		}
	}
	return deleteNode
}

func findNextGenFromLeft(rootLeft *Node) *Node {
	if rootLeft == nil {
		return nil
	}
	// 从root左边第一个节点开始找
	tmpNode := rootLeft
	for {
		if tmpNode.right != nil {
			tmpNode = tmpNode.right
		}
		return tmpNode
	}
}

func findNextGenFromRight(rootRight *Node) *Node {
	if rootRight == nil {
		return nil
	}
	// 从root右边第一个节点开始找
	tmpNode := rootRight
	for {
		if tmpNode.left != nil {
			tmpNode = tmpNode.left
		}
		return tmpNode
	}
}

func deleteNode(root *Node, v int) *Node {
	if root.left == nil && root.right == nil && v != root.data {
		return root
	}
	if v < root.data {
		deleteNode(root.left, v)
	} else if v > root.data {
		deleteNode(root.right, v)
	} else if v == root.data {
		// 此时的root即为要删除的节点
		leftNextGen := findNextGenFromLeft(root.left)
		rightNextGen := findNextGenFromRight(root.right)
		if leftNextGen == nil && rightNextGen == nil {
			// 要删除的root是叶子节点, 即最底部的节点
			down := root
			top := root.root
			if top.left == down {
				// 表示是左子树
				down.root = nil
				top.left = nil
				return nil
			} else {
				// 表示是右子树
				down.root = nil
				top.right = nil
				return nil
			}
		} else if leftNextGen != nil {
			root.data = leftNextGen.data
			deleteNode(leftNextGen, leftNextGen.data)
		} else if rightNextGen != nil {
			root.data = rightNextGen.data
			deleteNode(rightNextGen, rightNextGen.data)
		}
	}

	return root
}
