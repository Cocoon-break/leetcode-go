package linked_list

import "testing"

// go test -v . -test.run TestMergeKLists
func TestMergeKLists(t *testing.T) {
	l1 := GenListNode([]int{1, 4, 5})
	l2 := GenListNode([]int{1, 3, 4})
	l3 := GenListNode([]int{2, 6})
	lists := []*ListNode{l1, l2, l3}
	result := mergeKLists(lists)
	LogListNode(result)
}
