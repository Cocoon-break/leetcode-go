package linked_list

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	dummy := &ListNode{}
	head := dummy
	//定义两个游标
	movel1 := l1
	movel2 := l2
	inc := 0
	for movel1 != nil || movel2 != nil {
		x := 0
		if movel1 != nil {
			x = movel1.Val
			movel1 = movel1.Next
		}
		y := 0
		if movel2 != nil {
			y = movel2.Val
			movel2 = movel2.Next
		}
		// 两个节点值+进位值
		value := x + y + inc
		//取个位值
		val := value % 10
		//个位值+进位值
		tmp := &ListNode{Val: val}
		inc = int(value/10) % 10
		head.Next = tmp
		head = head.Next
	}
	//最后进位还要增加一个节点
	if inc > 0 {
		head.Next = &ListNode{Val: inc}
	}
	//返回反转之后的链表
	return dummy.Next
}
