package linked_list

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
