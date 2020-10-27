package linked_list

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

进阶：

你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
*/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//快慢指针
	slow, fast := head, head
	for fast != nil && fast.Next == nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow.Next
	//断链
	slow.Next = nil
	//递归处理
	right := sortList(head)
	left := sortList(tmp)
	//合并
	result := mergeTwoLists(right, left)
	return result
}
