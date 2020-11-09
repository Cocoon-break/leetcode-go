package linked_list

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
