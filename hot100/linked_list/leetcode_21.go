package linked_list

/*

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义哨子节点
	head := &ListNode{}
	node := head
	for l1 != nil && l2 != nil {
		// node 的下一个节点指向小的节点
		// 被指向的节点，向前
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		// node 节点向前
		node = node.Next
	}
	// 指向剩下长的节点
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head.Next
}
