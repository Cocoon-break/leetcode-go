package linked_list

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 定义哨子节点
	dummy := &ListNode{}
	dummy.Next = head
	fast, slow := dummy, dummy
	//定义快慢游标，快游标先走n步
	for i := 0; i < n; i++ {
		if fast != nil {
			fast = fast.Next
		} else {
			// 超过链表长度的n
			return head
		}
	}
	//快慢游标一直向前一步，直到快游标结束
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//将慢游标指向下一节点的下一个节点
	slow.Next = slow.Next.Next
	return dummy.Next
}
