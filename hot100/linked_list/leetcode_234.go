package linked_list

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

*/

func isPalindrome(head *ListNode) bool {
	// 只有一个节点时返回true
	if head == nil || head.Next == nil {
		return true
	}
	//利用哨子节点，判断只有两个节点的情况
	dummy := &ListNode{}
	dummy.Next = head
	//利用快慢指针，找出链表的前半部分和后半部分
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 重置快指针
	fast = head
	// 反转为nil
	if slow == nil {
		return false
	}
	// 反转链表的后半部分
	slow = reverseList(slow.Next)

	//比较节点中的值
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}
