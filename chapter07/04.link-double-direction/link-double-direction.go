package main

import (
	"fmt"
)

// type CompareFunc func(left,right interface{}) bool

type DLinkNode struct {
	data       int // interface{}
	prev, next *DLinkNode
}

func main() {
	fmt.Println("插入新节点")
	n2 := &DLinkNode{
		data: 3,
		prev: nil,
		next: nil,
	}
	root := buildDLink()

	root = insertDLNode(root, n2)
	root = insertDLNode(root, &DLinkNode{data: 0})
	root = insertDLNode(root, &DLinkNode{data: 11})
	root = insertDLNode(root, &DLinkNode{data: 7})
	root = insertDLNode(root, &DLinkNode{data: -1})
	rangeDLink(root)

	fmt.Println("删除节点")
	root = deleteDLNode(root, -1)
	root = deleteDLNode(root, 3)
	rangeDLink(root)

}

func buildDLink() *DLinkNode {
	n1 := &DLinkNode{data: 1}
	n3 := &DLinkNode{data: 5}
	n4 := &DLinkNode{data: 10}

	n1.next = n3
	n3.prev = n1

	n3.next = n4
	n4.prev = n3

	return n1
}

func insertDLNode(root *DLinkNode, newNode *DLinkNode) *DLinkNode {
	tmpNode := root

	// 整个链表是空的情况，新增
	if root == nil {
		return newNode
	}

	// 在链表的头，添加节点
	if newNode.data <= root.data {
		newNode.next = root
		root.prev = newNode
		return newNode
	}

	for {
		if tmpNode.next == nil {
			// 已经到头，在尾部追加节点即可
			tmpNode.next = newNode
			newNode.prev = tmpNode
			break
		} else {
			if newNode.data <= tmpNode.next.data {
				// 找到插入位置，开始插入节点
				newNode.prev = tmpNode
				newNode.next = tmpNode.next

				tmpNode.next = newNode
				// tmpNode.next.prev = newNode  // tmp的next已经是newNode了，再往回找会出错
				newNode.next.prev = newNode
				break
			}
		}
		tmpNode = tmpNode.next
	}
	return root
}

func rangeDLink(root *DLinkNode) {
	fmt.Println("从小到大")
	tmpNode := root

	for {
		fmt.Println(tmpNode.data)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}

	fmt.Println("从大到小")

	for {
		fmt.Println(tmpNode.data)
		if tmpNode.prev == nil {
			break
		}
		tmpNode = tmpNode.prev
	}
}

func deleteDLNode(root *DLinkNode, data int) *DLinkNode {
	if root == nil {
		return nil
	}
	if root.data == data {
		// 要删除的数据在第一个节点
		if root.next == nil {
			return nil
		} // 只有一个节点
		//leftHand := root
		//root = root.next
		//
		//leftHand.next= nil
		//root.prev = nil
		newRoot := root.next

		newRoot.prev = nil
		root.next = nil
		return newRoot
	}

	tmpNode := root
	for {
		if tmpNode.next == nil {
			// 走到链表的尾部，仍然没有找到要删除的数据，直接返回原root
			return root
		} else {
			if tmpNode.next.data == data {
				// 找到节点，开始删除，删除完成后返回原root
				rightHand := tmpNode.next
				tmpNode.next = rightHand.next
				rightHand.next.prev = tmpNode

				// 清理掉右手上的link，保证GC正常回收
				rightHand.prev = nil
				rightHand.next = nil
				return root
			}
		}
		tmpNode = tmpNode.next
	}

}
