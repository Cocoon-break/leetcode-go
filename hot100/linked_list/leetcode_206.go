package linked_list

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

func reverseList(head *ListNode) *ListNode {
	// 定义新的节点，list  nil
	var list *ListNode
	for head != nil {
		//备份下一个节点
		tmp := head.Next
		// 将当前节点的Next 指向上一个节点
		head.Next = list
		// 置换，将当前节点为新的节点
		list = head
		// head 继续向前
		head = tmp
	}
	return list
}
