package linked_list

import "testing"

func TestSortList(t *testing.T) {
	list := GenListNode([]int{-1, 5, 3, 4, 0})
	out := sortList(list)
	// out := sortList_O1(list)
	LogListNode(out)
}
