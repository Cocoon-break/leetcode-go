package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//根据数组生成链表节点
func GenListNode(values []int) *ListNode {
	//定义哨子节点
	head := &ListNode{}
	node := head
	for _, v := range values {
		tmp := &ListNode{Val: v, Next: nil}
		node.Next = tmp
		node = node.Next
	}
	return head.Next
}

//打印链接节点
func LogListNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
