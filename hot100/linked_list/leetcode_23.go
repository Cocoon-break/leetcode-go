package linked_list

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

*/
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, start int, end int) *ListNode {
	//开始和结束一样，说明只有一个链表，直接返回
	if start == end {
		return lists[start]
	}
	// 否则如果开始下标大于结束下标，直接返回空指针nil
	if start > end {
		return nil
	}
	// 如果不是上面两种情况，就分而治之
	// 先找到当前子数组的中间下标
	mid := start + (end-start)/2
	left := merge(lists, start, mid)
	right := merge(lists, mid+1, end)

	// 得到的结果是两条合并后的有序链表
	// 最后再把这两条链表也合并即可
	return mergeTwoLists(left, right)
}
