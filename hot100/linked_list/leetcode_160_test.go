package linked_list

import "testing"

// go test -v . -test.run TestGetIntersectionNode
func TestGetIntersectionNode(t *testing.T) {
	l1 := GenListNode([]int{4, 1})
	l2 := GenListNode([]int{})
	//公共部分
	// l3 := GenListNode([]int{8, 4, 5})
	// l1.Next = l3
	// l2.Next = l3
	out := getIntersectionNode(l1, l2)
	t.Log(out)
}
