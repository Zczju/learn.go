package main

import "fmt"

type LinkNode struct {
	data int // interface{}
	next *LinkNode
}

func main() {
	// 创建单链表
	n1 := &LinkNode{
		data: 1,
		next: nil,
	}
	n2 := &LinkNode{
		data: 2,
		next: nil,
	}
	n3 := &LinkNode{
		data: 3,
		next: nil,
	}
	n4 := &LinkNode{
		data: 4,
		next: nil,
	}
	n6 := &LinkNode{
		data: 6,
		next: nil,
	}

	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6

	rangeLink(n1)

	fmt.Println("输入 新的节点")
	n5 := &LinkNode{
		data: 5,
		next: nil,
	}
	root := insertNode(n1, n5)
	root = insertNode(n1, &LinkNode{
		data: 7,
		next: nil,
	})
	root = insertNode(n1, &LinkNode{
		data: 5,
		next: nil,
	})
	root = insertNode(n1, &LinkNode{
		data: 3,
		next: nil,
	})
	root = insertNode(n1, &LinkNode{
		data: 0,
		next: nil,
	})
	rangeLink(root)

	fmt.Println("删除 新的节点")
	root = deleteNode(root, 5)
	root = deleteNode(root, 0)
	root = deleteNode(root, 3)
	rangeLink(root)

}

func deleteNode(root *LinkNode, data int) *LinkNode {
	tmpNode := root
	for {
		right := tmpNode.next
		if root != nil && root.data == data { // 判断要删除的是否是第一个
			if root.next == nil { // 是第一个的情况下，判断是否只有一个
				return nil
			}
			// 并非只有一个，则删除root位置的链表
			root = root.next
			tmpNode.next = nil
			break
		}
		if right == nil { // 走到头还没找到，没有要删除的数据
			break
		}
		if right.data == data {
			// 找到删除位置，开始删除
			tmpNode.next = right.next
			right.next = nil
			return root
		}
		tmpNode = tmpNode.next
	}
	return root
}

// 单链表插入
func insertNode(root *LinkNode, newNode *LinkNode) (newRoot *LinkNode) {
	tmpNode := root
	for {
		if tmpNode != nil {
			if newNode.data > tmpNode.data {
				if tmpNode.next == nil {
					// 已经到结尾，直接追加
					tmpNode.next = newNode
					break
				} else {
					if newNode.data <= tmpNode.next.data {
						// 找到合适位置，准备插入数据
						newNode.next = tmpNode.next
						tmpNode.next = newNode
						break // 插入数据成功，break掉循环
					}
				}
			} else {
				newNode.next = tmpNode
				return newNode
			}
		} else {
			break
		}
		tmpNode = tmpNode.next
	}
	return root
}

// 单链表遍历
func rangeLink(root *LinkNode) {
	tmpNode := root
	for {
		if tmpNode != nil {
			fmt.Println(tmpNode.data)
		} else {
			break
		}
		tmpNode = tmpNode.next
	}
}
