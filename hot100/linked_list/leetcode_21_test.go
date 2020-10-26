package linked_list

import "testing"

// go test -v . -test.run TestMergeList
func TestMergeList(t *testing.T) {
	l1 := GenListNode([]int{1, 2, 4})
	l2 := GenListNode([]int{1, 3, 4})
	out := mergeTwoLists(l1, l2)
	LogListNode(out)
}
